package sds

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/mock"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
	"github.com/stratosnet/stratos-chain/x/pot"
	pottypes "github.com/stratosnet/stratos-chain/x/pot/types"
	"github.com/stratosnet/stratos-chain/x/register"
	regtypes "github.com/stratosnet/stratos-chain/x/register/types"
	"github.com/stratosnet/stratos-chain/x/sds/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	dbm "github.com/tendermint/tm-db"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var (
	paramSpecificMinedReward = sdk.NewInt(160000000000)
	paramSpecificEpoch       = sdk.NewInt(100)

	ppNodeOwner1   = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	ppNodeOwner2   = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	ppNodeOwner3   = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	ppNodeOwner4   = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	ppNodeOwnerNew = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())

	ppNodePubKey1   = ed25519.GenPrivKey().PubKey()
	ppNodeAddr1     = sdk.AccAddress(ppNodePubKey1.Address())
	ppInitialStake1 = sdk.NewInt(100000000)

	ppNodePubKey2   = ed25519.GenPrivKey().PubKey()
	ppNodeAddr2     = sdk.AccAddress(ppNodePubKey2.Address())
	ppInitialStake2 = sdk.NewInt(100000000)

	ppNodePubKey3   = ed25519.GenPrivKey().PubKey()
	ppNodeAddr3     = sdk.AccAddress(ppNodePubKey3.Address())
	ppInitialStake3 = sdk.NewInt(100000000)

	ppNodePubKey4   = ed25519.GenPrivKey().PubKey()
	ppNodeAddr4     = sdk.AccAddress(ppNodePubKey4.Address())
	ppInitialStake4 = sdk.NewInt(100000000)

	ppNodePubKeyNew = ed25519.GenPrivKey().PubKey()
	ppNodeAddrNew   = sdk.AccAddress(ppNodePubKeyNew.Address())
	ppNodeStakeNew  = sdk.NewInt(100000000)
)

func TestRandomPurchasedUoz(t *testing.T) {
	/********************* initialize mock app *********************/
	SetConfig()
	mApp, k, _, registerKeeper, _ := getMockAppPrepay(t)
	accs := setupAccounts(mApp)
	mock.SetGenesis(mApp, accs)
	//mock.CheckBalance(t, mApp, foundationAccAddr, foundationDeposit)

	header := abci.Header{Height: mApp.LastBlockHeight() + 1}
	ctx := mApp.BaseApp.NewContext(true, header)

	initialStakeTotal := sdk.NewInt(43000000000000)
	registerKeeper.SetInitialGenesisStakeTotal(ctx, initialStakeTotal)

	// setup resource nodes
	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")

	resourceNodeStake := sdk.NewInt(19000000000000)
	resouceNodeTokens := make([]sdk.Int, 0)
	numNodes := 10
	for i := 0; i < numNodes; i++ {
		resouceNodeTokens = append(resouceNodeTokens, resourceNodeStake)
	}

	log.Printf("Before: initial stake supply is %v \n\n", initialStakeTotal)
	initialUOzonePrice := registerKeeper.GetInitialUOzonePrice(ctx)
	log.Printf("Before: initial uozone price is %v \n\n", initialUOzonePrice)
	initOzoneLimit := initialStakeTotal.ToDec().Quo(initialUOzonePrice).TruncateInt()
	registerKeeper.SetRemainingOzoneLimit(ctx, initOzoneLimit)
	log.Printf("Before: remaining ozone limit is %v \n\n", registerKeeper.GetRemainingOzoneLimit(ctx))
	for i, val := range resouceNodeTokens {
		tmpResourceNode := regtypes.NewResourceNode("sds://resourceNode"+strconv.Itoa(i+1), ppNodePubKey1, ppNodeOwner1, regtypes.NewDescription("sds://resourceNode"+strconv.Itoa(i+1), "", "", "", ""), "storage", time)
		tmpResourceNode.Tokens = val
		tmpResourceNode.Status = sdk.Bonded
		tmpResourceNode.OwnerAddress = accs[i%5].GetAddress()
		ozoneLimitChange, _ := registerKeeper.AddResourceNodeStake(ctx, tmpResourceNode, sdk.NewCoin("ustos", val))
		log.Printf("Add resourceNode #%v(stake=%v), ozone limit increases by %v, remaining ozone limit is %v", i, resourceNodeStake, ozoneLimitChange, registerKeeper.GetRemainingOzoneLimit(ctx))
		// doPrepay
		randomPurchase := sdk.NewInt(int64(rand.Float64() * 100 * 1000000000))
		purchased, _ := k.Prepay(ctx, accs[i%5].GetAddress(), sdk.NewCoins(sdk.NewCoin("ustos", randomPurchase)))
		log.Printf("%v Uoz purchased by %v ustos, remaining ozone limit drops to %v", purchased, randomPurchase, registerKeeper.GetRemainingOzoneLimit(ctx))
	}
}

func TestPurchasedUoz(t *testing.T) {
	/********************* initialize mock app *********************/
	SetConfig()
	mApp, k, _, registerKeeper, _ := getMockAppPrepay(t)
	accs := setupAccounts(mApp)
	mock.SetGenesis(mApp, accs)
	//mock.CheckBalance(t, mApp, foundationAccAddr, foundationDeposit)

	header := abci.Header{Height: mApp.LastBlockHeight() + 1}
	ctx := mApp.BaseApp.NewContext(true, header)

	initialStakeTotal := sdk.NewInt(43000000000000)
	registerKeeper.SetInitialGenesisStakeTotal(ctx, initialStakeTotal)

	// setup resource nodes
	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")

	resourceNodeStake := sdk.NewInt(19000000000000)
	resouceNodeTokens := make([]sdk.Int, 0)
	numNodes := 10
	for i := 0; i < numNodes; i++ {
		resouceNodeTokens = append(resouceNodeTokens, resourceNodeStake)
	}

	log.Printf("Before: initial stake supply is %v \n\n", initialStakeTotal)
	initialUOzonePrice := registerKeeper.GetInitialUOzonePrice(ctx)
	log.Printf("Before: initial uozone price is %v \n\n", initialUOzonePrice)
	initOzoneLimit := initialStakeTotal.ToDec().Quo(initialUOzonePrice).TruncateInt()
	registerKeeper.SetRemainingOzoneLimit(ctx, initOzoneLimit)
	log.Printf("Before: remaining ozone limit is %v \n\n", registerKeeper.GetRemainingOzoneLimit(ctx))
	for i, val := range resouceNodeTokens {
		tmpResourceNode := regtypes.NewResourceNode("sds://resourceNode"+strconv.Itoa(i+1), ppNodePubKey1, ppNodeOwner1, regtypes.NewDescription("sds://resourceNode"+strconv.Itoa(i+1), "", "", "", ""), "storage", time)
		tmpResourceNode.Tokens = val
		tmpResourceNode.Status = sdk.Bonded
		tmpResourceNode.OwnerAddress = accs[i%5].GetAddress()
		ozoneLimitChange, _ := registerKeeper.AddResourceNodeStake(ctx, tmpResourceNode, sdk.NewCoin("ustos", val))
		log.Printf("Add resourceNode #%v(stake=%v), ozone limit increases by %v, remaining ozone limit is %v", i, resourceNodeStake, ozoneLimitChange, registerKeeper.GetRemainingOzoneLimit(ctx))
		// doPrepay
		purchased, _ := k.Prepay(ctx, accs[i%5].GetAddress(), sdk.NewCoins(sdk.NewCoin("ustos", sdk.NewInt(10000000000))))
		log.Printf("%v Uoz purchased by 10 stos, remaining ozone limit drops to %v", purchased, registerKeeper.GetRemainingOzoneLimit(ctx))
	}
}

func TestOzoneLimitChange(t *testing.T) {
	/********************* initialize mock app *********************/
	SetConfig()
	mApp, k, _, registerKeeper, _ := getMockApp(t)
	accs := setupAccounts(mApp)
	mock.SetGenesis(mApp, accs)

	header := abci.Header{Height: mApp.LastBlockHeight() + 1}
	ctx := mApp.BaseApp.NewContext(true, header)

	initialStakeTotal := sdk.NewInt(43000000000000)
	registerKeeper.SetInitialGenesisStakeTotal(ctx, initialStakeTotal)

	// setup resource nodes
	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")

	resourceNodeStake := sdk.NewInt(19000000000000)
	resouceNodeTokens := make([]sdk.Int, 0)
	numNodes := 10
	for i := 0; i < numNodes; i++ {
		resouceNodeTokens = append(resouceNodeTokens, resourceNodeStake)
	}

	//init pp nodes.
	log.Printf("Before: initial stake supply is %v \n\n", initialStakeTotal)
	initialUOzonePrice := registerKeeper.GetInitialUOzonePrice(ctx)
	log.Printf("Before: initial uozone price is %v \n\n", initialUOzonePrice)
	//initOzoneLimit := initialStakeTotal.ToDec().Quo(initialUOzonePrice.ToDec()).TruncateInt()
	//registerKeeper.SetRemainingOzoneLimit(ctx, initOzoneLimit)
	log.Printf("Before: remaining ozone limit is %v \n\n", registerKeeper.GetRemainingOzoneLimit(ctx))
	for i, val := range resouceNodeTokens {
		tmpResourceNode := regtypes.NewResourceNode("sds://resourceNode"+strconv.Itoa(i+1), ppNodePubKey1, ppNodeOwner1, regtypes.NewDescription("sds://resourceNode"+strconv.Itoa(i+1), "", "", "", ""), "storage", time)
		tmpResourceNode.Tokens = val
		tmpResourceNode.Status = sdk.Bonded
		tmpResourceNode.OwnerAddress = accs[i%5].GetAddress()
		ozoneLimitChange, _ := registerKeeper.AddResourceNodeStake(ctx, tmpResourceNode, sdk.NewCoin("ustos", val))
		log.Printf("Add resourceNode #%v(stake=%v), ozone limit increases by %v, remaining ozone limit is %v", i, resourceNodeStake, ozoneLimitChange, registerKeeper.GetRemainingOzoneLimit(ctx))
		// doPrepay
		purchased, _ := k.Prepay(ctx, accs[i%5].GetAddress(), sdk.NewCoins(sdk.NewCoin("ustos", sdk.NewInt(10000000000))))
		log.Printf("%v Uoz purchased by 10 stos", purchased)
	}
}

func TestSdsMsgs(t *testing.T) {

	/********************* initialize mock app *********************/
	SetConfig()
	mApp, keeper, _, _, _ := getMockApp(t)
	accs := setupAccounts(mApp)
	mock.SetGenesis(mApp, accs)
	//mock.CheckBalance(t, mApp, foundationAccAddr, foundationDeposit)

	///********************* create fileUpload msg *********************/
	log.Print("====== Testing MsgFileUpload ======")
	fileHash, _ := hex.DecodeString(testFileHashHex)
	fileUploadMsg := types.NewMsgUpload(fileHash, sdsAccAddr1, spNodeAddrIdx1, sdsAccAddr2)
	headerUpload := abci.Header{Height: mApp.LastBlockHeight() + 1}
	mock.SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, headerUpload, []sdk.Msg{fileUploadMsg}, []uint64{18}, []uint64{0}, true, true, sdsAccPrivKey1)
	coin := sdk.NewCoin(DefaultDenom, spNodeInitialStakeIdx1)
	mock.CheckBalance(t, mApp, spNodeAddrIdx1, sdk.Coins{coin})
	///********************* create prepay msg *********************/
	log.Print("====== Testing MsgPrepay ======")
	coinToPrepay := sdk.NewCoin(DefaultDenom, prepayAmt)
	prepayMsg := types.NewMsgPrepay(sdsAccAddr3, sdk.NewCoins(coinToPrepay))
	headerPrepay := abci.Header{Height: mApp.LastBlockHeight() + 1}
	mock.SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, headerPrepay, []sdk.Msg{prepayMsg}, []uint64{20}, []uint64{0}, true, true, sdsAccPrivKey3)
	newBalanceInt := sdsAccBal3.Sub(prepayAmt)
	newBalanceCoin := sdk.NewCoin(DefaultDenom, newBalanceInt)
	mock.CheckBalance(t, mApp, sdsAccAddr3, sdk.NewCoins(newBalanceCoin))

	///********************* test uozPrice *********************/
	header := abci.Header{Height: mApp.LastBlockHeight() + 1}
	ctx := mApp.BaseApp.NewContext(true, header)
	log.Print("====== Testing uozPrice ======\n\n")
	initS := sdk.NewInt(43000)
	initLt := sdk.NewInt(43000)
	initPt := sdk.NewInt(0)

	keeper.RegisterKeeper.SetInitialGenesisStakeTotal(ctx, initS)
	keeper.PotKeeper.SetTotalUnissuedPrepay(ctx, initPt)
	keeper.RegisterKeeper.SetRemainingOzoneLimit(ctx, initLt)

	log.Printf("==== init stake total is %v", keeper.RegisterKeeper.GetInitialGenesisStakeTotal(ctx))
	log.Printf("==== init prepay is %v", keeper.PotKeeper.GetTotalUnissuedPrepay(ctx))
	log.Printf("==== ozone limit is %v\n\n", keeper.RegisterKeeper.GetRemainingOzoneLimit(ctx))

	numPrepay := 5
	prepaySeq := make([]sdk.Int, 0)
	for i := 0; i < numPrepay; i++ {
		prepaySeq = append(prepaySeq, sdk.NewInt(19000))
	}

	for i, val := range prepaySeq {
		S := keeper.RegisterKeeper.GetInitialGenesisStakeTotal(ctx)
		Pt := keeper.PotKeeper.GetTotalUnissuedPrepay(ctx)
		Lt := keeper.RegisterKeeper.GetRemainingOzoneLimit(ctx)

		uozPriceBefore := S.ToDec().Add(Pt.ToDec()).Quo(Lt.ToDec()).TruncateInt()

		uozPurchased := Lt.ToDec().
			Mul(val.ToDec()).
			Quo((S.
				Add(Pt).
				Add(val)).ToDec()).
			TruncateInt()

		uozPriceAfter := S.ToDec().Add(Pt.ToDec()).Add(val.ToDec()).Quo(Lt.ToDec().Sub(uozPurchased.ToDec())).TruncateInt()

		Pt = Pt.Add(val)
		Lt = Lt.Sub(uozPurchased)
		keeper.PotKeeper.SetTotalUnissuedPrepay(ctx, Pt)
		keeper.RegisterKeeper.SetRemainingOzoneLimit(ctx, Lt)
		log.Printf("---- prepay #%v: %v ustos----", i, val)
		log.Printf("uozPriceBefore is %v", uozPriceBefore)
		log.Printf("uozPurchased is %v", uozPurchased)
		log.Printf("uozPriceAfter is %v", uozPriceAfter)
		log.Printf("New Pt is %v", Pt)
		log.Printf("New Lt is %v", Lt)
		log.Print("\n")
	}
}

func getMockApp(t *testing.T) (*mock.App, Keeper, bank.Keeper, register.Keeper, pot.Keeper) {
	mApp := mock.NewApp()

	RegisterCodec(mApp.Cdc)
	supply.RegisterCodec(mApp.Cdc)
	staking.RegisterCodec(mApp.Cdc)
	register.RegisterCodec(mApp.Cdc)

	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
	keyRegister := sdk.NewKVStoreKey(register.StoreKey)
	keyPot := sdk.NewKVStoreKey(pot.StoreKey)
	keySds := sdk.NewKVStoreKey(StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)

	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyStaking, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyRegister, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyPot, sdk.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	feeCollector := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
	notBondedPool := supply.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
	bondPool := supply.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)

	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollector.GetAddress().String()] = true
	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	blacklistedAddrs[bondPool.GetAddress().String()] = true

	bankKeeper := bank.NewBaseKeeper(mApp.AccountKeeper, mApp.ParamsKeeper.Subspace(bank.DefaultParamspace), blacklistedAddrs)
	maccPerms := map[string][]string{
		auth.FeeCollectorName:     {"fee_collector"},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
	}
	supplyKeeper := supply.NewKeeper(mApp.Cdc, keySupply, mApp.AccountKeeper, bankKeeper, maccPerms)
	stakingKeeper := staking.NewKeeper(mApp.Cdc, keyStaking, supplyKeeper, mApp.ParamsKeeper.Subspace(staking.DefaultParamspace))
	registerKeeper := register.NewKeeper(mApp.Cdc, keyRegister, mApp.ParamsKeeper.Subspace(register.DefaultParamSpace), mApp.AccountKeeper, bankKeeper)
	potKeeper := pot.NewKeeper(mApp.Cdc, keyPot, mApp.ParamsKeeper.Subspace(DefaultParamSpace), auth.FeeCollectorName, bankKeeper, supplyKeeper, mApp.AccountKeeper, stakingKeeper, registerKeeper)
	keeper := NewKeeper(mApp.Cdc, keySds, bankKeeper, registerKeeper, potKeeper)

	mApp.Router().AddRoute(staking.RouterKey, staking.NewHandler(stakingKeeper))
	mApp.Router().AddRoute(RouterKey, NewHandler(keeper))
	mApp.SetEndBlocker(getEndBlocker(keeper))
	mApp.SetInitChainer(getInitChainer(mApp, keeper, mApp.AccountKeeper, supplyKeeper,
		[]supplyexported.ModuleAccountI{feeCollector, notBondedPool, bondPool}, stakingKeeper, registerKeeper, potKeeper))

	err = mApp.CompleteSetup(keySds, keyStaking, keySupply, keyRegister, keyPot)
	require.NoError(t, err)

	return mApp, keeper, bankKeeper, registerKeeper, potKeeper
}

func getMockAppPrepay(t *testing.T) (*mock.App, Keeper, bank.Keeper, register.Keeper, pot.Keeper) {
	mApp := mock.NewApp()

	RegisterCodec(mApp.Cdc)
	supply.RegisterCodec(mApp.Cdc)
	staking.RegisterCodec(mApp.Cdc)
	register.RegisterCodec(mApp.Cdc)

	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
	keyRegister := sdk.NewKVStoreKey(register.StoreKey)
	keyPot := sdk.NewKVStoreKey(pot.StoreKey)
	keySds := sdk.NewKVStoreKey(StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)

	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keySupply, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyStaking, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyRegister, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyPot, sdk.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	//ctx := sdk.NewContext(ms, abci.Header{ChainID: "foochainid"}, false, tmtLog.NewNopLogger())

	feeCollector := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
	notBondedPool := supply.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
	bondPool := supply.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)

	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[feeCollector.GetAddress().String()] = true
	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
	blacklistedAddrs[bondPool.GetAddress().String()] = true

	//accountKeeper := auth.NewAccountKeeper(mApp.Cdc, keyAcc, mApp.ParamsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(mApp.AccountKeeper, mApp.ParamsKeeper.Subspace(bank.DefaultParamspace), blacklistedAddrs)
	maccPerms := map[string][]string{
		auth.FeeCollectorName:     {"fee_collector"},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
	}
	supplyKeeper := supply.NewKeeper(mApp.Cdc, keySupply, mApp.AccountKeeper, bankKeeper, maccPerms)
	stakingKeeper := staking.NewKeeper(mApp.Cdc, keyStaking, supplyKeeper, mApp.ParamsKeeper.Subspace(staking.DefaultParamspace))
	registerKeeper := register.NewKeeper(mApp.Cdc, keyRegister, mApp.ParamsKeeper.Subspace(register.DefaultParamSpace), mApp.AccountKeeper, bankKeeper)
	potKeeper := pot.NewKeeper(mApp.Cdc, keyPot, mApp.ParamsKeeper.Subspace(DefaultParamSpace), auth.FeeCollectorName, bankKeeper, supplyKeeper, mApp.AccountKeeper, stakingKeeper, registerKeeper)
	keeper := NewKeeper(mApp.Cdc, keySds, bankKeeper, registerKeeper, potKeeper)

	mApp.Router().AddRoute(staking.RouterKey, staking.NewHandler(stakingKeeper))
	mApp.Router().AddRoute(RouterKey, NewHandler(keeper))
	mApp.SetEndBlocker(getEndBlocker(keeper))
	mApp.SetInitChainer(getInitChainerTestPurchase(mApp, keeper, mApp.AccountKeeper, supplyKeeper,
		[]supplyexported.ModuleAccountI{feeCollector, notBondedPool, bondPool}, stakingKeeper, registerKeeper, potKeeper))

	err = mApp.CompleteSetup(keySds, keyStaking, keySupply, keyRegister, keyPot)
	require.NoError(t, err)

	return mApp, keeper, bankKeeper, registerKeeper, potKeeper
}

// getInitChainer initializes the chainer of the mock app and sets the genesis
// state. It returns an empty ResponseInitChain.
func getInitChainer(mapp *mock.App, keeper Keeper, accountKeeper auth.AccountKeeper, supplyKeeper supply.Keeper,
	blacklistedAddrs []supplyexported.ModuleAccountI, stakingKeeper staking.Keeper, registerKeeper register.Keeper, potKeeper pot.Keeper) sdk.InitChainer {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		// set module accounts
		for _, macc := range blacklistedAddrs {
			supplyKeeper.SetModuleAccount(ctx, macc)
		}

		mapp.InitChainer(ctx, req)

		var lastResourceNodeStakes []register.LastResourceNodeStake
		lastResourceNodeStakes = append(lastResourceNodeStakes, register.LastResourceNodeStake{Address: addrRes1, Stake: initialStakeRes1})
		lastResourceNodeStakes = append(lastResourceNodeStakes, register.LastResourceNodeStake{Address: addrRes2, Stake: initialStakeRes2})
		lastResourceNodeStakes = append(lastResourceNodeStakes, register.LastResourceNodeStake{Address: addrRes3, Stake: initialStakeRes3})
		lastResourceNodeStakes = append(lastResourceNodeStakes, register.LastResourceNodeStake{Address: addrRes4, Stake: initialStakeRes4})
		lastResourceNodeStakes = append(lastResourceNodeStakes, register.LastResourceNodeStake{Address: addrRes5, Stake: initialStakeRes5})

		var lastIndexingNodeStakes []register.LastIndexingNodeStake
		lastIndexingNodeStakes = append(lastIndexingNodeStakes, register.LastIndexingNodeStake{Address: addrIdx1, Stake: initialStakeIdx1})
		lastIndexingNodeStakes = append(lastIndexingNodeStakes, register.LastIndexingNodeStake{Address: addrIdx2, Stake: initialStakeIdx2})
		lastIndexingNodeStakes = append(lastIndexingNodeStakes, register.LastIndexingNodeStake{Address: addrIdx3, Stake: initialStakeIdx3})
		lastIndexingNodeStakes = append(lastIndexingNodeStakes, register.LastIndexingNodeStake{Address: spNodeAddrIdx1, Stake: spNodeInitialStakeIdx1})

		resourceNodes := setupAllResourceNodes()
		indexingNodes := setupAllIndexingNodes()

		registerGenesis := register.NewGenesisState(register.DefaultParams(), lastResourceNodeStakes, resourceNodes, lastIndexingNodeStakes, indexingNodes, initialUOzonePrice)

		register.InitGenesis(ctx, registerKeeper, registerGenesis)

		// set module accounts
		for _, macc := range blacklistedAddrs {
			supplyKeeper.SetModuleAccount(ctx, macc)
		}

		stakingGenesis := staking.NewGenesisState(staking.NewParams(staking.DefaultUnbondingTime, staking.DefaultMaxValidators, staking.DefaultMaxEntries, 0, DefaultDenom), nil, nil)

		totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000)))
		supplyKeeper.SetSupply(ctx, supply.NewSupply(totalSupply))

		// set module accounts
		for _, macc := range blacklistedAddrs {
			supplyKeeper.SetModuleAccount(ctx, macc)
		}

		validators := staking.InitGenesis(ctx, stakingKeeper, accountKeeper, supplyKeeper, stakingGenesis)

		//preset
		registerKeeper.SetRemainingOzoneLimit(ctx, remainingOzoneLimit)
		potKeeper.SetTotalUnissuedPrepay(ctx, totalUnissuedPrepay)

		//pot genesis data load
		pot.InitGenesis(ctx, potKeeper, pot.NewGenesisState(pottypes.DefaultParams()))

		// init bank genesis
		keeper.BankKeeper.SetSendEnabled(ctx, true)

		return abci.ResponseInitChain{
			Validators: validators,
		}
	}
}

// getInitChainer initializes the chainer of the mock app and sets the genesis
// state. It returns an empty ResponseInitChain.
func getInitChainerTestPurchase(mapp *mock.App, keeper Keeper, accountKeeper auth.AccountKeeper, supplyKeeper supply.Keeper,
	blacklistedAddrs []supplyexported.ModuleAccountI, stakingKeeper staking.Keeper, registerKeeper register.Keeper, potKeeper pot.Keeper) sdk.InitChainer {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		// set module accounts
		for _, macc := range blacklistedAddrs {
			supplyKeeper.SetModuleAccount(ctx, macc)
		}

		mapp.InitChainer(ctx, req)

		var lastResourceNodeStakes []register.LastResourceNodeStake

		var lastIndexingNodeStakes []register.LastIndexingNodeStake
		lastIndexingNodeStakes = append(lastIndexingNodeStakes, register.LastIndexingNodeStake{Address: addrIdx1, Stake: initialStakeIdx1})

		//resourceNodes := setupAllResourceNodes()
		indexingNodes := setupAllIndexingNodes()

		registerGenesis := register.NewGenesisState(register.DefaultParams(), lastResourceNodeStakes, nil, lastIndexingNodeStakes, indexingNodes, initialUOzonePriceTestPurchase)

		register.InitGenesis(ctx, registerKeeper, registerGenesis)

		// set module accounts
		for _, macc := range blacklistedAddrs {
			supplyKeeper.SetModuleAccount(ctx, macc)
		}

		stakingGenesis := staking.NewGenesisState(staking.NewParams(staking.DefaultUnbondingTime, staking.DefaultMaxValidators, staking.DefaultMaxEntries, 0, DefaultDenom), nil, nil)

		totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000)))
		supplyKeeper.SetSupply(ctx, supply.NewSupply(totalSupply))

		// set module accounts
		for _, macc := range blacklistedAddrs {
			supplyKeeper.SetModuleAccount(ctx, macc)
		}

		validators := staking.InitGenesis(ctx, stakingKeeper, accountKeeper, supplyKeeper, stakingGenesis)

		//preset
		registerKeeper.SetRemainingOzoneLimit(ctx, remainingOzoneLimitTestPurchase)
		potKeeper.SetTotalUnissuedPrepay(ctx, totalUnissuedPrepayTestPurchase)

		//pot genesis data load
		pot.InitGenesis(ctx, potKeeper, pot.NewGenesisState(pottypes.DefaultParams()))

		// init bank genesis
		keeper.BankKeeper.SetSendEnabled(ctx, true)

		return abci.ResponseInitChain{
			Validators: validators,
		}
	}
}

// getEndBlocker returns a staking endblocker.
func getEndBlocker(keeper Keeper) sdk.EndBlocker {
	return func(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
		validatorUpdates := keeper.PotKeeper.StakingKeeper.BlockValidatorUpdates(ctx)

		return abci.ResponseEndBlock{
			ValidatorUpdates: validatorUpdates,
		}
	}
	return nil
}
