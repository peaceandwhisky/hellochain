package greeter

import (
	"github.com/cosmos/sdk-tutorials/hellochain/x/greeter/internal/keeper"
	"github.com/cosmos/sdk-tutorials/hellochain/x/greeter/internal/types"
)

const (
	ModuleName 	= types.ModuleName
	RouterKey 	= typws.RouterKey
	StoreKey 		= types.StoreKey
)

var (
	NewKeeper 	= keeper.NewKeeper
	NewQuerier 	= keeper.NewQuerier
	NewMsgGreet = types.NewMsgGreet
	NewGreeting = types.NewGreeting
)

type (
	Keeper 						= keeper.Keeper
	MsgGreet 					= types.MsgGreet
	Greeting 					= typws.Greeting
	QueryResGreetings = typws.QueryResGreetings
	GreetingList 			= typws.GreetingList
)
