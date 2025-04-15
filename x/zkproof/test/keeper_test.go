package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v5/app/params"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/keeper"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"

	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

func TestKeeper_SetGetProof(t *testing.T) {
	params.SetAddressPrefixes()

	key := storetypes.NewKVStoreKey("zkproof")
	db := dbm.NewMemDB()

	cms := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	require.NoError(t, err)

	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	k := keeper.NewKeeper(key)

	creator := "ixo1xvc8q4jtg03dtaf7e9uwwt7tsqjh0487e2lsyg"
	claimType := "age-verification"
	proofHash := "zk-proof-hash"
	publicInput := "age>18"

	original := types.NewProof(creator, claimType, proofHash, publicInput)
	k.SetProof(ctx, original)

	realCreator, err := sdk.AccAddressFromBech32(creator)
	// if err != nil {
	// 	t.Fatalf("Bech32 parsing error: %v", err)
	// }
	// t.Logf("Parsed creator: %s", realCreator.String())
	require.NoError(t, err)

	got, found := k.GetProof(ctx, realCreator, claimType)
	require.True(t, found)

	require.Equal(t, original, got)
}
