package cli

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/stratosnet/stratos-chain/x/sds/client/common"
	"strings"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stratosnet/stratos-chain/x/sds/types"
)

// GetQueryCmd returns the cli query commands for sds module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group sds queries under a subcommand
	sdsQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	sdsQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdQueryUploadedFile(queryRoute, cdc),
			GetCmdQueryPrepayBalance(queryRoute, cdc),
		)...,
	)

	return sdsQueryCmd
}

// GetCmdQueryUploadedFile implements the query uploaded file command.
func GetCmdQueryUploadedFile(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "upload [sender_addr]",
		Args:  cobra.RangeArgs(1, 1),
		Short: "Query uploaded file's hash by sender addr",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query uploaded file's hash by sender addr, optionally restrict to rewards from a single validator.

Example:
$ %s query sds upload st1yx3kkx9jnqeck59j744nc5qgtv4lt4dc45jcwz
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// query file by fileHash
			if len(args) == 1 {
				resp, _, err := common.QueryUploadedFile(cliCtx, queryRoute, args[0])
				if err != nil {
					return err
				}
				return cliCtx.PrintOutput(hex.EncodeToString(resp))
			}
			return nil
		},
	}
}

// GetCmdQueryPrepayBalance implements the query prepay balance command.
func GetCmdQueryPrepayBalance(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "prepay [sender_addr]",
		Args:  cobra.RangeArgs(1, 1),
		Short: "Query balance of prepayment in Volumn Pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query balance of prepayment in Volumn Pool.

Example:
$ %s query sds prepay st1yx3kkx9jnqeck59j744nc5qgtv4lt4dc45jcwz
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// query file by fileHash
			if len(args) == 1 {
				resp, _, err := common.QueryPrepayBalance(cliCtx, queryRoute, args[0])
				if err != nil {
					return err
				}
				var prepaidBalance sdk.Int
				error := prepaidBalance.UnmarshalJSON(resp)
				if error != nil {
					return error
				}
				return cliCtx.PrintOutput(prepaidBalance.String())
			}
			return nil
		},
	}
}
