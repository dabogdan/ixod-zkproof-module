package zkproof

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.zkproof.v1beta.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "SubmitProof",
					Use:       "submit-proof [claim-type] [proof-data]",
					Short:     "Submit a zkProof claim",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "claim_type"},
						{ProtoField: "proof_data"},
						{ProtoField: "public_input"},
					},
				},
			},
		},
	}
}
