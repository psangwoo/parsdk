package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// Keeper of the oracle store
type Keeper struct {
	allowedAddressStoreKey sdk.StoreKey
	exchangeRateStoreKey   sdk.StoreKey
	cdc                    codec.BinaryMarshaler
}

// NewKeeper creates an oracle keeper
func NewKeeper(cdc codec.BinaryMarshaler, allowedAddresskey sdk.StoreKey, exchangeRatekey sdk.StoreKey) Keeper {
	keeper := Keeper{
		allowedAddressStoreKey: allowedAddresskey,
		exchangeRateStoreKey:   exchangeRatekey,
		cdc:                    cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetAllowedAddress(ctx sdk.Context) (allowedAddress string) {
	store := ctx.KVStore(k.allowedAddressStoreKey)
	b := store.Get(types.ParamStoreKeyAllowedAddress)
	if b == nil {
		panic("stored allowed address should not have been nil")
	}

	return string(b)
}

func (k Keeper) SetAllowedAddress(ctx sdk.Context, allowedAddress string) {
	store := ctx.KVStore(k.exchangeRateStoreKey)
	store.Set(types.ParamStoreKeyAllowedAddress, []byte(allowedAddress))
}

func (k Keeper) GetExchangeRate(ctx sdk.Context) (exchangeRate sdk.Coin) {
	store := ctx.KVStore(k.exchangeRateStoreKey)
	b := store.Get(types.ExchangeRateKey)
	if b == nil {
		panic("stored exchange rate should not have been nil")
	}

	k.cdc.MustUnmarshalBinaryBare(b, &exchangeRate)
	return
}

func (k Keeper) SetExchangeRate(ctx sdk.Context, exchangeRate sdk.Coin) {
	store := ctx.KVStore(k.exchangeRateStoreKey)
	b := k.cdc.MustMarshalBinaryBare(&exchangeRate)
	store.Set(types.ExchangeRateKey, b)
}

func (k Keeper) DeleteExchangeRate(ctx sdk.Context) {
	store := ctx.KVStore(k.exchangeRateStoreKey)
	store.Delete(types.ExchangeRateKey)
}