package filters

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tendermint/tendermint/libs/log"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/stratosnet/stratos-chain/rpc/backend"
	"github.com/stratosnet/stratos-chain/rpc/types"
	evmtypes "github.com/stratosnet/stratos-chain/x/evm/types"
)

// consider a filter inactive if it has not been polled for within deadline
var deadline = 5 * time.Minute

// filter is a helper struct that holds meta information over the filter type
// and associated subscription in the event system.
type filter struct {
	typ      filters.Type
	deadline *time.Timer // filter is inactive when deadline triggers
	hashes   []common.Hash
	crit     filters.FilterCriteria
	logs     []*ethtypes.Log
	s        *Subscription // associated subscription in event system
}

// PublicFilterAPI offers support to create and manage filters. This will allow external clients to retrieve various
// information related to the Ethereum protocol such as blocks, transactions and logs.
type PublicFilterAPI struct {
	logger    log.Logger
	clientCtx client.Context
	backend   backend.BackendI
	events    *EventSystem
	filtersMu sync.Mutex
	filters   map[rpc.ID]*filter
}

// NewPublicAPI returns a new PublicFilterAPI instance.
func NewPublicAPI(logger log.Logger, clientCtx client.Context, eventBus *tmtypes.EventBus, b backend.BackendI) *PublicFilterAPI {
	logger = logger.With("api", "filter")
	api := &PublicFilterAPI{
		logger:    logger,
		clientCtx: clientCtx,
		backend:   b,
		filters:   make(map[rpc.ID]*filter),
		events:    NewEventSystem(clientCtx, logger, eventBus),
	}

	go api.timeoutLoop()

	return api
}

// timeoutLoop runs every 5 minutes and deletes filters that have not been recently used.
// Tt is started when the api is created.
func (api *PublicFilterAPI) timeoutLoop() {
	ticker := time.NewTicker(deadline)
	defer ticker.Stop()

	for {
		<-ticker.C
		api.filtersMu.Lock()
		for id, f := range api.filters {
			select {
			case <-f.deadline.C:
				f.s.Unsubscribe(api.events)
				delete(api.filters, id)
			default:
				continue
			}
		}
		api.filtersMu.Unlock()
	}
}

// NewPendingTransactionFilter creates a filter that fetches pending transaction hashes
// as transactions enter the pending state.
//
// It is part of the filter package because this filter can be used through the
// `eth_getFilterChanges` polling method that is also used for log filters.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newPendingTransactionFilter
func (api *PublicFilterAPI) NewPendingTransactionFilter() rpc.ID {
	if len(api.filters) >= int(api.backend.RPCFilterCap()) {
		return rpc.ID("error creating pending tx filter: max limit reached")
	}

	pendingTxSub, cancelSubs, err := api.events.SubscribePendingTxs()
	if err != nil {
		// wrap error on the ID
		return rpc.ID(fmt.Sprintf("error creating pending tx filter: %s", err.Error()))
	}
	api.filtersMu.Lock()
	api.filters[pendingTxSub.ID()] = &filter{typ: filters.PendingTransactionsSubscription, deadline: time.NewTimer(deadline), hashes: make([]common.Hash, 0), s: pendingTxSub}
	api.filtersMu.Unlock()

	go func(txsCh <-chan coretypes.ResultEvent, errCh <-chan error) {
		defer cancelSubs()

		for {
			select {
			case ev := <-txsCh:
				data, ok := ev.Data.(tmtypes.EventDataTx)
				if !ok {
					err = fmt.Errorf("invalid event data %T, expected %s", ev.Data, tmtypes.EventTx)
					pendingTxSub.err <- err
					return
				}

				txHash, err := types.GetTxHash(api.clientCtx.TxConfig.TxDecoder(), tmtypes.Tx(data.Tx))
				if err != nil {
					pendingTxSub.err <- err
					return
				}

				api.filtersMu.Lock()
				if f, found := api.filters[pendingTxSub.ID()]; found {
					f.hashes = append(f.hashes, txHash)
				}
				api.filtersMu.Unlock()
			case <-errCh:
				api.filtersMu.Lock()
				delete(api.filters, pendingTxSub.ID())
				api.filtersMu.Unlock()
			}
		}
	}(pendingTxSub.eventCh, pendingTxSub.Err())

	return pendingTxSub.ID()
}

// NewPendingTransactions creates a subscription that is triggered each time a transaction
// enters the transaction pool and was signed from one of the transactions this nodes manages.
func (api *PublicFilterAPI) NewPendingTransactions(ctx context.Context) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	rpcSub := notifier.CreateSubscription()

	ctx, cancelFn := context.WithTimeout(context.Background(), deadline)
	defer cancelFn()

	api.events.WithContext(ctx)

	pendingTxSub, cancelSubs, err := api.events.SubscribePendingTxs()
	if err != nil {
		return nil, err
	}

	go func(txsCh <-chan coretypes.ResultEvent) {
		defer cancelSubs()

		for {
			select {
			case ev := <-txsCh:
				data, ok := ev.Data.(tmtypes.EventDataTx)
				if !ok {
					err = fmt.Errorf("invalid event data %T, expected %s", ev.Data, tmtypes.EventTx)
					pendingTxSub.err <- err
					return
				}

				txHash, err := types.GetTxHash(api.clientCtx.TxConfig.TxDecoder(), tmtypes.Tx(data.Tx))
				if err != nil {
					pendingTxSub.err <- err
					return
				}

				// To keep the original behaviour, send a single tx hash in one notification.
				// TODO(rjl493456442) Send a batch of tx hashes in one notification
				err = notifier.Notify(rpcSub.ID, txHash)
				if err != nil {
					return
				}
			case <-rpcSub.Err():
				pendingTxSub.Unsubscribe(api.events)
				return
			case <-notifier.Closed():
				pendingTxSub.Unsubscribe(api.events)
				return
			}
		}
	}(pendingTxSub.eventCh)

	return rpcSub, err
}

// NewBlockFilter creates a filter that fetches blocks that are imported into the chain.
// It is part of the filter package since polling goes with eth_getFilterChanges.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newblockfilter
func (api *PublicFilterAPI) NewBlockFilter() rpc.ID {
	headersSub, cancelSubs, err := api.events.SubscribeNewHeads()
	if err != nil {
		// wrap error on the ID
		return rpc.ID(fmt.Sprintf("error creating block filter: %s", err.Error()))
	}

	api.filtersMu.Lock()
	api.filters[headersSub.ID()] = &filter{typ: filters.BlocksSubscription, deadline: time.NewTimer(deadline), hashes: []common.Hash{}, s: headersSub}
	api.filtersMu.Unlock()

	go func(headersCh <-chan coretypes.ResultEvent, errCh <-chan error) {
		defer cancelSubs()

		for {
			select {
			case ev := <-headersCh:
				data, ok := ev.Data.(tmtypes.EventDataNewBlockHeader)
				if !ok {
					err = fmt.Errorf("invalid event data %T, expected %s", ev.Data, tmtypes.EventNewBlockHeader)
					headersSub.err <- err
					return
				}
				header, err := types.EthHeaderFromTendermint(data.Header)
				if err != nil {
					headersSub.err <- err
					return
				}
				// override dynamicly miner address
				sdkCtx, err := api.backend.GetEVMContext().GetSdkContextWithHeader(&data.Header)
				if err != nil {
					headersSub.err <- err
					return
				}

				validator, err := api.backend.GetEVMKeeper().GetCoinbaseAddress(sdkCtx)
				if err != nil {
					headersSub.err <- err
					return
				}
				header.Coinbase = validator

				api.filtersMu.Lock()
				if f, found := api.filters[headersSub.ID()]; found {
					f.hashes = append(f.hashes, header.Hash)
				}
				api.filtersMu.Unlock()
			case <-errCh:
				api.filtersMu.Lock()
				delete(api.filters, headersSub.ID())
				api.filtersMu.Unlock()
				return
			}
		}
	}(headersSub.eventCh, headersSub.Err())

	return headersSub.ID()
}

// NewHeads send a notification each time a new (header) block is appended to the chain.
func (api *PublicFilterAPI) NewHeads(ctx context.Context) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	api.events.WithContext(ctx)
	rpcSub := notifier.CreateSubscription()

	headersSub, cancelSubs, err := api.events.SubscribeNewHeads()
	if err != nil {
		return &rpc.Subscription{}, err
	}

	go func(headersCh <-chan coretypes.ResultEvent) {
		defer cancelSubs()

		for {
			select {
			case ev := <-headersCh:
				data, ok := ev.Data.(tmtypes.EventDataNewBlockHeader)
				if !ok {
					err = fmt.Errorf("invalid event data %T, expected %s", ev.Data, tmtypes.EventNewBlockHeader)
					headersSub.err <- err
					return
				}

				header, err := types.EthHeaderFromTendermint(data.Header)
				if err != nil {
					headersSub.err <- err
					return
				}
				// override dynamicly miner address
				sdkCtx, err := api.backend.GetEVMContext().GetSdkContextWithHeader(&data.Header)
				if err != nil {
					headersSub.err <- err
					return
				}

				validator, err := api.backend.GetEVMKeeper().GetCoinbaseAddress(sdkCtx)
				if err != nil {
					headersSub.err <- err
					return
				}
				header.Coinbase = validator

				err = notifier.Notify(rpcSub.ID, header)
				if err != nil {
					headersSub.err <- err
					return
				}
			case <-rpcSub.Err():
				headersSub.Unsubscribe(api.events)
				return
			case <-notifier.Closed():
				headersSub.Unsubscribe(api.events)
				return
			}
		}
	}(headersSub.eventCh)

	return rpcSub, err
}

// Logs creates a subscription that fires for all new log that match the given filter criteria.
func (api *PublicFilterAPI) Logs(ctx context.Context, crit filters.FilterCriteria) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	api.events.WithContext(ctx)
	rpcSub := notifier.CreateSubscription()

	logsSub, cancelSubs, err := api.events.SubscribeLogs(crit)
	if err != nil {
		return &rpc.Subscription{}, err
	}

	go func(logsCh <-chan coretypes.ResultEvent) {
		defer cancelSubs()

		for {
			select {
			case ev := <-logsCh:
				_, isMsgEthereumTx := ev.Events[fmt.Sprintf("%s.%s", evmtypes.EventTypeEthereumTx, evmtypes.AttributeKeyEthereumTxHash)]
				api.logger.Debug("\x1b[32m------ logs tx type is evm from sub: %t\x1b[0m\n", isMsgEthereumTx)
				if !isMsgEthereumTx {
					continue
				}

				// get transaction result data
				dataTx, ok := ev.Data.(tmtypes.EventDataTx)
				if !ok {
					err = fmt.Errorf("invalid event data %T, expected %s", ev.Data, tmtypes.EventTx)
					logsSub.err <- err
					return
				}
				api.logger.Debug("\x1b[32m------ logs tx dataTx: %+v\x1b[0m\n", dataTx)

				txResponse, err := evmtypes.DecodeTxResponse(dataTx.TxResult.Result.Data)
				if err != nil {
					logsSub.err <- err
					return
				}
				api.logger.Debug("\x1b[32m------ logs tx response: %+v\x1b[0m\n", txResponse)
				api.logger.Debug("\x1b[32m------ logs crit: %+v\x1b[0m\n", crit)

				matchedLogs := FilterLogs(evmtypes.LogsToEthereum(txResponse.Logs), crit.FromBlock, crit.ToBlock, crit.Addresses, crit.Topics)
				api.logger.Debug("\x1b[32m------ logs matchedLogs: %+v\x1b[0m\n", matchedLogs)
				for _, log := range matchedLogs {
					err = notifier.Notify(rpcSub.ID, log)
					if err != nil {
						logsSub.err <- err
						return
					}
				}
			case <-rpcSub.Err(): // client send an unsubscribe request
				logsSub.Unsubscribe(api.events)
				return
			case <-notifier.Closed(): // connection dropped
				logsSub.Unsubscribe(api.events)
				return
			}
		}
	}(logsSub.eventCh)

	return rpcSub, err
}

// NewFilter creates a new filter and returns the filter id. It can be
// used to retrieve logs when the state changes. This method cannot be
// used to fetch logs that are already stored in the state.
//
// Default criteria for the from and to block are "latest".
// Using "latest" as block number will return logs for mined blocks.
// Using "pending" as block number returns logs for not yet mined (pending) blocks.
// In case logs are removed (chain reorg) previously returned logs are returned
// again but with the removed property set to true.
//
// In case "fromBlock" > "toBlock" an error is returned.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_newfilter
func (api *PublicFilterAPI) NewFilter(criteria filters.FilterCriteria) (rpc.ID, error) {
	if len(api.filters) >= int(api.backend.RPCFilterCap()) {
		return rpc.ID(""), fmt.Errorf("error creating filter: max limit reached")
	}

	var (
		filterID = rpc.ID("")
		err      error
	)

	logsSub, cancelSubs, err := api.events.SubscribeLogs(criteria)
	if err != nil {
		return rpc.ID(""), err
	}

	filterID = logsSub.ID()

	api.filtersMu.Lock()
	api.filters[filterID] = &filter{typ: filters.LogsSubscription, crit: criteria, deadline: time.NewTimer(deadline), hashes: []common.Hash{}, s: logsSub}
	api.filtersMu.Unlock()

	go func(eventCh <-chan coretypes.ResultEvent) {
		defer cancelSubs()

		for {
			select {
			case ev := <-eventCh:
				_, isMsgEthereumTx := ev.Events[fmt.Sprintf("%s.%s", evmtypes.EventTypeEthereumTx, evmtypes.AttributeKeyEthereumTxHash)]
				if !isMsgEthereumTx {
					continue
				}

				dataTx, ok := ev.Data.(tmtypes.EventDataTx)
				if !ok {
					err = fmt.Errorf("invalid event data %T, expected EventDataTx", ev.Data)
					logsSub.err <- err
					return
				}

				txResponse, err := evmtypes.DecodeTxResponse(dataTx.TxResult.Result.Data)
				if err != nil {
					logsSub.err <- err
					return
				}

				logs := FilterLogs(evmtypes.LogsToEthereum(txResponse.Logs), criteria.FromBlock, criteria.ToBlock, criteria.Addresses, criteria.Topics)

				api.filtersMu.Lock()
				if f, found := api.filters[filterID]; found {
					f.logs = append(f.logs, logs...)
				}
				api.filtersMu.Unlock()
			case <-logsSub.Err():
				api.filtersMu.Lock()
				delete(api.filters, filterID)
				api.filtersMu.Unlock()
				return
			}
		}
	}(logsSub.eventCh)

	return filterID, err
}

// GetLogs returns logs matching the given argument that are stored within the state.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getlogs
func (api *PublicFilterAPI) GetLogs(ctx context.Context, crit filters.FilterCriteria) ([]*ethtypes.Log, error) {
	var filter *Filter
	if crit.BlockHash != nil {
		// Block filter requested, construct a single-shot filter
		filter = NewBlockFilter(api.logger, api.backend, crit)
	} else {
		// Convert the RPC block numbers into internal representations
		begin := rpc.LatestBlockNumber.Int64()
		if crit.FromBlock != nil {
			begin = crit.FromBlock.Int64()
		}
		end := rpc.LatestBlockNumber.Int64()
		if crit.ToBlock != nil {
			end = crit.ToBlock.Int64()
		}
		// Construct the range filter
		filter = NewRangeFilter(api.logger, api.backend, begin, end, crit.Addresses, crit.Topics)
	}

	// Run the filter and return all the logs
	logs, err := filter.Logs(ctx, int(api.backend.RPCLogsCap()), int64(api.backend.RPCBlockRangeCap()))
	if err != nil {
		return nil, err
	}

	return returnLogs(logs), err
}

// UninstallFilter removes the filter with the given filter id.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_uninstallfilter
func (api *PublicFilterAPI) UninstallFilter(id rpc.ID) bool {
	api.filtersMu.Lock()
	f, found := api.filters[id]
	if found {
		delete(api.filters, id)
	}
	api.filtersMu.Unlock()

	if !found {
		return false
	}
	f.s.Unsubscribe(api.events)
	return true
}

// GetFilterLogs returns the logs for the filter with the given id.
// If the filter could not be found an empty array of logs is returned.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getfilterlogs
func (api *PublicFilterAPI) GetFilterLogs(ctx context.Context, id rpc.ID) ([]*ethtypes.Log, error) {
	api.filtersMu.Lock()
	f, found := api.filters[id]
	api.filtersMu.Unlock()

	if !found {
		return returnLogs(nil), fmt.Errorf("filter %s not found", id)
	}

	if f.typ != filters.LogsSubscription {
		return returnLogs(nil), fmt.Errorf("filter %s doesn't have a LogsSubscription type: got %d", id, f.typ)
	}

	var filter *Filter
	if f.crit.BlockHash != nil {
		// Block filter requested, construct a single-shot filter
		filter = NewBlockFilter(api.logger, api.backend, f.crit)
	} else {
		// Convert the RPC block numbers into internal representations
		begin := rpc.LatestBlockNumber.Int64()
		if f.crit.FromBlock != nil {
			begin = f.crit.FromBlock.Int64()
		}
		end := rpc.LatestBlockNumber.Int64()
		if f.crit.ToBlock != nil {
			end = f.crit.ToBlock.Int64()
		}
		// Construct the range filter
		filter = NewRangeFilter(api.logger, api.backend, begin, end, f.crit.Addresses, f.crit.Topics)
	}
	// Run the filter and return all the logs
	logs, err := filter.Logs(ctx, int(api.backend.RPCLogsCap()), int64(api.backend.RPCBlockRangeCap()))
	if err != nil {
		return nil, err
	}
	return returnLogs(logs), nil
}

// GetFilterChanges returns the logs for the filter with the given id since
// last time it was called. This can be used for polling.
//
// For pending transaction and block filters the result is []common.Hash.
// (pending)Log filters return []Log.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getfilterchanges
func (api *PublicFilterAPI) GetFilterChanges(id rpc.ID) (interface{}, error) {
	api.filtersMu.Lock()
	defer api.filtersMu.Unlock()

	f, found := api.filters[id]
	if !found {
		return nil, fmt.Errorf("filter %s not found", id)
	}

	if !f.deadline.Stop() {
		// timer expired but filter is not yet removed in timeout loop
		// receive timer value and reset timer
		<-f.deadline.C
	}
	f.deadline.Reset(deadline)

	switch f.typ {
	case filters.PendingTransactionsSubscription, filters.BlocksSubscription:
		hashes := f.hashes
		f.hashes = nil
		return returnHashes(hashes), nil
	case filters.LogsSubscription, filters.MinedAndPendingLogsSubscription:
		logs := make([]*ethtypes.Log, len(f.logs))
		copy(logs, f.logs)
		f.logs = []*ethtypes.Log{}
		return returnLogs(logs), nil
	default:
		return nil, fmt.Errorf("invalid filter %s type %d", id, f.typ)
	}
}

// Syncing provides information when this nodes starts synchronising with the OneLedger network and when it's finished.
func (api *PublicFilterAPI) Syncing() (*rpc.Subscription, error) {
	return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
}
