package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "zkproof",
		Short:                      "Querying commands for the zkproof module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryProof(),
	)

	return cmd
}

func CmdQueryProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proof [claim-type] --creator [creator-address]",
		Short: "Query a zkproof by claim type and creator address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			creator, err := cmd.Flags().GetString("creator")
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryProofRequest{
				Creator:   creator,
				ClaimType: args[0],
			}

			res, err := queryClient.Proof(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String("creator", "", "Creator (Bech32 address)")
	_ = cmd.MarkFlagRequired("creator")

	return cmd
}
