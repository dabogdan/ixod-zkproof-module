package test

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v5/app/params"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/keeper"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

func TestMsgServer(t *testing.T) {
	params.SetAddressPrefixes()

	// store setup
	storeKey := storetypes.NewKVStoreKey("zkproof")
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, cms.LoadLatestVersion())

	//create context and keeper
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	k := keeper.NewKeeper(storeKey)
	msgServer := keeper.NewMsgServerImpl(*k)

	// dummy input
	creator := "ixo1xvc8q4jtg03dtaf7e9uwwt7tsqjh0487e2lsyg"
	claimType := "age-verification"
	proofHash := "zk-proof-hash"
	publicInput := "age>18"

	msg := &types.MsgSubmitProof{
		Creator:     creator,
		ClaimType:   claimType,
		ProofData:   proofHash,
		PublicInput: publicInput,
	}

	resp, err := msgServer.SubmitProof(ctx, msg)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)

	addr, err := sdk.AccAddressFromBech32(creator)
	require.NoError(t, err)

	stored, found := k.GetProof(ctx, addr, claimType)
	require.True(t, found)
	require.Equal(t, proofHash, stored.ProofHash)
	require.Equal(t, publicInput, stored.PublicInput)
	require.Equal(t, creator, stored.Creator)
	require.Equal(t, claimType, stored.ClaimType)
}
