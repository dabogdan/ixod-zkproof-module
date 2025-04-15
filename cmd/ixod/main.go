package main

import (
	"fmt"
	"os"
	"path/filepath"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ixofoundation/ixo-blockchain/v5/app/params"
	"github.com/ixofoundation/ixo-blockchain/v5/cmd/ixod/cmd"
)

func main() {
	params.SetAddressPrefixes()

	// Get the main CLI root command with all Cosmos preconfigured commands
	rootCmd := cmd.NewRootCmd()

	// Provide a fallback default home directory
	homeDir := filepath.Join(os.Getenv("HOME"), ".ixod")

	if err := svrcmd.Execute(rootCmd, "IXO", homeDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
