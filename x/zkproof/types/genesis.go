package types

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// )

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

// ValidateGenesis validates the genesis state
func ValidateGenesis(_ GenesisState) error {
	return nil
}

// GenesisState defines the module's genesis state
type GenesisState struct{}
