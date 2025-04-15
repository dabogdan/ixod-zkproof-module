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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestQueryServer_Proof(t *testing.T) {
	params.SetAddressPrefixes()

	// Setup store
	storeKey := storetypes.NewKVStoreKey("zkproof")
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, cms.LoadLatestVersion())

	// Create context and keeper
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	k := keeper.NewKeeper(storeKey)
	queryServer := k // implements types.QueryServer

	// Store a test proof
	creator := "ixo1xvc8q4jtg03dtaf7e9uwwt7tsqjh0487e2lsyg"
	claimType := "age-verification"
	proofHash := "zk-proof-hash"
	publicInput := "age>18"

	proof := types.NewProof(creator, claimType, proofHash, publicInput)
	k.SetProof(ctx, proof)

	// Valid query
	resp, err := queryServer.Proof(ctx, &types.QueryProofRequest{
		Creator:   creator,
		ClaimType: claimType,
	})
	require.NoError(t, err)
	require.Equal(t, creator, resp.Creator)
	require.Equal(t, claimType, resp.ClaimType)
	require.Equal(t, proofHash, resp.ProofData)
	require.Equal(t, publicInput, resp.PublicInput)

	// Nil request
	_, err = queryServer.Proof(ctx, nil)
	require.ErrorContains(t, err, "empty request")

	// Invalid address
	_, err = queryServer.Proof(ctx, &types.QueryProofRequest{
		Creator:   "not-an-address",
		ClaimType: claimType,
	})
	require.Error(t, err)
	require.Equal(t, sdkerrors.ErrInvalidAddress.ABCICode(), sdkerrors.ErrInvalidAddress.ABCICode()) // optional check

	// Not found
	other := sdk.AccAddress([]byte("somekey")).String()
	_, err = queryServer.Proof(ctx, &types.QueryProofRequest{
		Creator:   other,
		ClaimType: "some-other-claim",
	})
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
	require.Equal(t, "invalid address", st.Message())

}
