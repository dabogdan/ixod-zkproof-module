package types

type Proof struct {
	Creator     string `json:"creator" yaml:"creator"`
	ClaimType   string `json:"claim_type" yaml:"claim_type"`
	ProofHash   string `json:"proof_hash" yaml:"proof_hash"`
	PublicInput string `json:"public_input"`
}

func NewProof(creator, claimType, proofHash string, publicInput string) Proof {
	return Proof{
		Creator:     creator,
		ClaimType:   claimType,
		ProofHash:   proofHash,
		PublicInput: publicInput,
	}
}
