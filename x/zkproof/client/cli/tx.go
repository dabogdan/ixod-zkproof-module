package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"
)

func CmdSubmitProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proof [claim-type] [proof-hash]",
		Short: "Submit a zkProof claim",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			claimType := args[0]
			proofHash := args[1]
			publicInput := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitProof(
				clientCtx.GetFromAddress().String(),
				claimType,
				proofHash,
				publicInput,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "zkproof",
		Short:                      "zkproof transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(CmdSubmitProof())

	return txCmd
}
