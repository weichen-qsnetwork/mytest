package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stratosnet/stratos-chain/x/pot/types"
)

func (k Keeper) SetFoundationAccount(ctx sdk.Context, acc sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(acc)
	store.Set(types.FoundationAccountKey, b)
}

func (k Keeper) GetFoundationAccount(ctx sdk.Context) (acc sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.FoundationAccountKey)
	if b == nil {
		panic("Stored foundation account should not have been nil")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &acc)
	return
}

func (k Keeper) SetInitialUOzonePrice(ctx sdk.Context, price sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(price)
	store.Set(types.InitialUOzonePriceKey, b)
}

func (k Keeper) GetInitialUOzonePrice(ctx sdk.Context) (price sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.InitialUOzonePriceKey)
	if b == nil {
		panic("Stored initial uOzone price should not have been nil")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &price)
	return
}

func (k Keeper) SetMatureEpoch(ctx sdk.Context, matureEpoch sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(matureEpoch)
	store.Set(types.MatureEpochKey, b)
}

func (k Keeper) GetMatureEpoch(ctx sdk.Context) (matureEpoch sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MatureEpochKey)
	if b == nil {
		panic("Stored mature epoch should not have been nil")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &matureEpoch)
	return
}

func (k Keeper) setMinedTokens(ctx sdk.Context, minedToken sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(minedToken)
	store.Set(types.MinedTokensKey, b)
}

func (k Keeper) getMinedTokens(ctx sdk.Context) (minedToken sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinedTokensKey)
	if b == nil {
		return sdk.ZeroInt()
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &minedToken)
	return
}

func (k Keeper) SetTotalUnissuedPrepay(ctx sdk.Context, totalUnissuedPrepay sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(totalUnissuedPrepay)
	store.Set(types.TotalUnissuedPrepayKey, b)
}

func (k Keeper) GetTotalUnissuedPrepay(ctx sdk.Context) (totalUnissuedPrepay sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.TotalUnissuedPrepayKey)
	if b == nil {
		return sdk.ZeroInt()
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &totalUnissuedPrepay)
	return
}

func (k Keeper) setRewardAddressPool(ctx sdk.Context, addressList []sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(addressList)
	store.Set(types.RewardAddressPoolKey, b)
}

func (k Keeper) getRewardAddressPool(ctx sdk.Context) (addressList []sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.RewardAddressPoolKey)
	if b == nil {
		return nil
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &addressList)
	return
}

func (k Keeper) setLastMaturedEpoch(ctx sdk.Context, epoch sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(epoch)
	store.Set(types.LastMaturedEpochKey, b)
}

func (k Keeper) getLastMaturedEpoch(ctx sdk.Context) (epoch sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.LastMaturedEpochKey)
	if b == nil {
		return sdk.ZeroInt()
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &epoch)
	return
}

func (k Keeper) setIndividualReward(ctx sdk.Context, acc sdk.AccAddress, epoch sdk.Int, value sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	store.Set(types.GetIndividualRewardKey(acc, epoch), b)
}

func (k Keeper) getIndividualReward(ctx sdk.Context, acc sdk.AccAddress, epoch sdk.Int) (value sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetIndividualRewardKey(acc, epoch))
	if b == nil {
		return sdk.ZeroInt()
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &value)
	return
}

func (k Keeper) setMatureTotalReward(ctx sdk.Context, acc sdk.AccAddress, value sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	store.Set(types.GetMatureTotalRewardKey(acc), b)
}

func (k Keeper) getMatureTotalReward(ctx sdk.Context, acc sdk.AccAddress) (value sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetMatureTotalRewardKey(acc))
	if b == nil {
		return sdk.ZeroInt()
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &value)
	return
}

func (k Keeper) setImmatureTotalReward(ctx sdk.Context, acc sdk.AccAddress, value sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(value)
	store.Set(types.GetImmatureTotalRewardKey(acc), b)
}

func (k Keeper) getImmatureTotalReward(ctx sdk.Context, acc sdk.AccAddress) (value sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetImmatureTotalRewardKey(acc))
	if b == nil {
		return sdk.ZeroInt()
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &value)
	return
}