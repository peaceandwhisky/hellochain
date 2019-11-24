package hellochain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/sdk-tutorials/hellochain/starter"

	//import greeter types
	"github.com/cosmos/sdk-tutorials/hellochain/x/greeter"
)

const appName = "hellochain"

var (
	// ModuleBasics holds the AppModuleBasic struct of all modules included in the app
	ModuleBasics = starter.ModuleBasics
)

// Add the keeper and its key to our app struct
type helloChainApp struct {
	*starter.AppStarter // helloChainApp extends starter.AppStarter
	greeterKey		*sdk.KVStoreKey // the store key for the greeter module
	greeterKeeper greeter.Keeper // the keeper for the greeter module
}

// NewHelloChainApp returns a fully constructed SDK application
func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Application {

	// pass greeter's AppModuleBasic to be included int the ModuleBasicsManager
	appStarter := starter.NewAppStarter(appName, logger, db, greeter.AppModuleBasic{})

	// create the key for greeter's store
	greeterKey := sdk.NewKVStoreKey(greeter.StoreKey)

	// construct the keeper
	greeterKeeper := greeter.NewKeeper(greeterKey, appStarter.Cdc)

	// compose our app with greeter
	var app = &helloChainApp{
		appStarter,
		greeterKey,
		greeterKeeper,
	}

	// Add greeter's complete AppModule to the ModuleManager
	greeterMod := greeter.NewAppModule(greeterKeeper)
	app.Mm.Modules[greeterMod.Name()] = greeterMod

	// create a subspace for greeter's data in the main store.
	app.MountStore(greeterKey, sdk.StoreTypeDB)

	// do some final configuration...
	app.InitializeStarter()

	return app
}


