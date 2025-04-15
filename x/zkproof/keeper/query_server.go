package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Proof(goCtx context.Context, req *types.QueryProofRequest) (*types.QueryProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	proof, found := k.GetProof(ctx, creatorAddr, req.ClaimType)
	if !found {
		return nil, status.Error(codes.NotFound, "proof not found")
	}

	return &types.QueryProofResponse{
		Creator:     proof.Creator,
		ClaimType:   proof.ClaimType,
		ProofData:   proof.ProofHash,
		PublicInput: proof.PublicInput,
	}, nil
}
