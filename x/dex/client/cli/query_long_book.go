package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
	"github.com/spf13/cobra"
)

func CmdListLongBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-long-book",
		Short: "list all longBook",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLongBookRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LongBookAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowLongBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-long-book [contract address] [price] [price denom] [asset denom]",
		Short: "shows a longBook",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			contractAddr := args[0]
			price, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}
			priceDenom := types.Denom(types.Denom_value[args[2]])
			assetDenom := types.Denom(types.Denom_value[args[3]])

			params := &types.QueryGetLongBookRequest{
				Price:        price,
				ContractAddr: contractAddr,
				PriceDenom:   priceDenom,
				AssetDenom:   assetDenom,
			}
			res, err := queryClient.LongBook(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
