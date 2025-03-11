package app

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/your-repo/gpu-resource-sharing-network/blockchain/module/taskservice"
	"github.com/your-repo/gpu-resource-sharing-network/blockchain/module/paymentservice"
	"github.com/your-repo/gpu-resource-sharing-network/blockchain/module/scheduler"
)

type MyApp struct {
	*baseapp.BaseApp
	cdc        *codec.Codec
	coinsKeeper bank.Keeper
}

func NewApp() *MyApp {
	cdc := codec.New()
	app := &MyApp{
		BaseApp: baseapp.NewBaseApp("gpu-resource-sharing", nil, nil),
		cdc:     cdc,
	}

	app.RegisterModule(taskservice.NewModule())
	app.RegisterModule(paymentservice.NewModule())
	app.RegisterModule(scheduler.NewModule())
	app.RegisterModule(auth.NewModule())
	app.RegisterModule(bank.NewModule())

	return app
}

func (app *MyApp) InitGenesis() {
	// Initialize the blockchain genesis state here
	fmt.Println("Initializing Genesis")
}
