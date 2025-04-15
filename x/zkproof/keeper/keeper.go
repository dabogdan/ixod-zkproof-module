package keeper

import (
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	storetypes "cosmossdk.io/store/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
}

func NewKeeper(storeKey storetypes.StoreKey) *Keeper {
	return &Keeper{storeKey: storeKey}
}

func (k Keeper) SetProof(ctx sdk.Context, proof types.Proof) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProofKey(proof.Creator, proof.ClaimType)
	bz := k.MustMarshalProof(proof)
	store.Set(key, bz)
}

func (k Keeper) GetProof(ctx sdk.Context, creator sdk.AccAddress, claimType string) (types.Proof, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProofKey(creator.String(), claimType)

	bz := store.Get(key)
	if bz == nil {
		return types.Proof{}, false
	}

	return k.MustUnmarshalProof(bz), true
}

func (k Keeper) MustMarshalProof(proof types.Proof) []byte {
	return types.Amino.MustMarshalJSON(&proof)
}

func (k Keeper) MustUnmarshalProof(bz []byte) types.Proof {
	var proof types.Proof
	types.Amino.MustUnmarshalJSON(bz, &proof)
	return proof
}
