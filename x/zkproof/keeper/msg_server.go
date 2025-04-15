package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) SubmitProof(goCtx context.Context, msg *types.MsgSubmitProof) (*types.MsgSubmitProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proof := types.NewProof(msg.Creator, msg.ClaimType, msg.ProofData, msg.PublicInput)
	m.SetProof(ctx, proof)

	return &types.MsgSubmitProofResponse{
		Status: "success",
	}, nil
}
