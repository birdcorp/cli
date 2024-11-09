package cmd

import (
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var couponListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all coupons",
	Long:  `List all available coupons in your account.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, apiClient := mustGetAuth()

		coupons, _, err := apiClient.CouponCodesAPI.
			ListCouponCodes(ctx).
			Execute()
		if err != nil {
			return err
		}

		prettyprint.JSON(coupons)
		return nil
	},
}

func init() {
	couponCmd.AddCommand(couponListCmd)
}

/*
go run main.go coupon list
*/
