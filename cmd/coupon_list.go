package cmd

import (
	"log"
	"time"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var couponListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all coupons",
	Long:  `List all available coupons in your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Fetching coupons..."
		s.Start()

		coupons, _, err := apiClient.CouponCodesAPI.
			ListCouponCodes(ctx).
			Execute()
		if err != nil {
			s.Stop()
			log.Fatalf("Error listing coupons: %v", err)
			return
		}
		s.Stop()

		printer.CouponsList(coupons.Data)

	},
}

func init() {
	couponCmd.AddCommand(couponListCmd)
}

/*
go run main.go coupon list
*/
