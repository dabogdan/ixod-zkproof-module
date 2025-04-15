package zkproof

import (
	"encoding/json"

	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/keeper"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ixofoundation/ixo-blockchain/v5/x/zkproof/client/cli"
	"github.com/spf13/cobra"
)

type AppModuleBasic struct{}

var (
	_ module.AppModuleBasic = AppModuleBasic{}
	_ module.AppModule      = AppModule{}
)

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	bz, err := json.Marshal(types.DefaultGenesis())
	if err != nil {
		panic(err)
	}
	return bz
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ interface{}, bz json.RawMessage) error {
	var data types.GenesisState
	return json.Unmarshal(bz, &data)
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(
	clientCtx sdkclient.Context,
	mux *runtime.ServeMux,
) {
	// ToDo: implement
}

// AppModule implements the AppModule interface.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{keeper: k}
}

func (am AppModule) Name() string {
	return types.ModuleName
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func (am AppModule) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

func (AppModule) IsAppModule()        {}
func (AppModule) IsOnePerModuleType() {}
