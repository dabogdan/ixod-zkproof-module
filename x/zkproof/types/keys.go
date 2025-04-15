package types

import "fmt"

const (
	ModuleName = "zkproof"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
	MemStoreKey = "mem_zkproof"
)

func GetProofKey(creator string, claimType string) []byte {
	return []byte(fmt.Sprintf("proof:%s:%s", creator, claimType))
}
