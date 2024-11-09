package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var couponDeleteCmd = &cobra.Command{
	Use:   "delete [coupon-id]",
	Short: "Delete a coupon",
	Long:  `Delete a coupon by its ID.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, apiClient := mustGetAuth()

		_, err := apiClient.CouponCodesAPI.
			DeleteCouponCode(ctx, args[0]).
			Execute()
		if err != nil {
			return err
		}

		fmt.Printf("Deleted coupon: %s\n", args[0])
		return nil
	},
}

func init() {
	couponCmd.AddCommand(couponDeleteCmd)
}

/*
go run main.go coupon delete akpyDbB88kz2sUQBwaunD2
*/
