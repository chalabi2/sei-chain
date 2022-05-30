package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
)

// SetLongBook set a specific longBook in the store
func (k Keeper) SetLongBook(ctx sdk.Context, contractAddr string, longBook types.LongBook) {
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		types.OrderBookPrefix(
			true, contractAddr, longBook.Entry.PriceDenom, longBook.Entry.AssetDenom,
		),
	)
	b := k.cdc.MustMarshal(&longBook)
	store.Set(GetKeyForLongBook(longBook), b)
}

func (k Keeper) GetLongBookByPrice(ctx sdk.Context, contractAddr string, price sdk.Dec, priceDenom types.Denom, assetDenom types.Denom) (val types.LongBook, found bool) {
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		types.OrderBookPrefix(
			true, contractAddr, priceDenom, assetDenom,
		),
	)
	b := store.Get(GetKeyForPrice(price))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveLongBookByPrice(ctx sdk.Context, contractAddr string, price sdk.Dec, priceDenom types.Denom, assetDenom types.Denom) {
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		types.OrderBookPrefix(
			true, contractAddr, priceDenom, assetDenom,
		),
	)
	store.Delete(GetKeyForPrice(price))
}

// GetAllLongBook returns all longBook
func (k Keeper) GetAllLongBook(ctx sdk.Context, contractAddr string) (list []types.LongBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ContractKeyPrefix(types.LongBookKey, contractAddr))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LongBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetAllLongBookForPair(ctx sdk.Context, contractAddr string, priceDenom types.Denom, assetDenom types.Denom) (list []types.OrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.OrderBookPrefix(true, contractAddr, priceDenom, assetDenom))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LongBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, &val)
	}

	return
}

func GetKeyForLongBook(longBook types.LongBook) []byte {
	return GetKeyForPrice(longBook.Entry.Price)
}

func GetKeyForPrice(price sdk.Dec) []byte {
	key, err := price.Marshal()
	if err != nil {
		panic(err)
	}
	return key
}
