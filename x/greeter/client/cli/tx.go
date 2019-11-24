package cli

import (
	//"fmt"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	gtypes "github.com/cosmos/sdk-tutorials/hellochain/x/greeter/internal/types"
)

// GetTxCmd returns the parent transaction command for the greeting module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	greetingTxCmd := &cobra.Command{
		Use: 												"greeter",
		Short: 											"greeter transaction subcommands",
		DisableFlagParsing: 				true,
		SuggestionMinimumDistance: 	2,
		RunE:												client.ValidateCmd,
	}

	greetingTxCmd.AddCommand(client.PostCommands(
		GetCmdSayHello(cdc),
	)...)

	return greetingTxCmd
}

// GetCmdSayHello returns the tx say command for the greeter module
func GetCmdSayHello(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "say [body] [addr]",
		Short: "send a greeting to another user. Usage: say [body] [address]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Get the sending account address provided to the --form flag
			sender := cliCtx.GetFromAddress()
			body := args[0]

			// Construct the receiving address from the procided argument
			recipient, err := sdk.AccAddressFromBech32(args[1])

			if err != nil {
				return err
			}

			// used to construct, sign and encode the transction (Tx) to send our greeting message
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := gtypes.NewMsgGreet(sender, body, recipient)
			err = msg.Validatebasic()
			if err != nil {
				return err
			}

			// Build, sign and broadcast our transaction containing our greeting message.
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}



