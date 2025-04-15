package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgSubmitProof = "submit_proof"

var _ sdk.Msg = &MsgSubmitProof{} // MsgSubmitProof comes from tx.pb.go

func NewMsgSubmitProof(creator, claimType, proofData string, publicInput string) *MsgSubmitProof {
	return &MsgSubmitProof{
		Creator:     creator,
		ClaimType:   claimType,
		ProofData:   proofData,
		PublicInput: publicInput,
	}
}

func (msg *MsgSubmitProof) Route() string { return RouterKey }
func (msg *MsgSubmitProof) Type() string  { return TypeMsgSubmitProof }

func (msg *MsgSubmitProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %w", err)
	}

	if msg.ClaimType == "" {
		return fmt.Errorf("claim_type cannot be empty")
	}

	if msg.ProofData == "" {
		return fmt.Errorf("proof_data cannot be empty")
	}

	if msg.PublicInput == "" {
		return fmt.Errorf("proof_data cannot be empty")
	}

	return nil
}
