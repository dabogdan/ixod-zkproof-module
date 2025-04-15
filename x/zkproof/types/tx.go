package types

import (
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zkproof",
		Short: "zkproof transaction subcommands",
	}

	// Add your commands like SubmitProof here:
	cmd.AddCommand(
		CmdSubmitProof(),
	)

	return cmd
}

func CmdSubmitProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proof [claim-type] [proof-data]",
		Short: "Submit a zkProof",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// You can parse args here and build the MsgSubmitProof
			return nil
		},
	}
	return cmd
}

// func (msg *MsgSubmitProof) GetSigners() []sdk.AccAddress {
// 	creator, err := sdk.AccAddressFromBech32(msg.Creator)
// 	if err != nil {
// 		panic(err) // or return nil if you prefer graceful failure
// 	}
// 	return []sdk.AccAddress{creator}
// }
