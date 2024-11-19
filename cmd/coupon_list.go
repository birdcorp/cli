package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/formatting"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var couponListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all coupons",
	Long:  `List all available coupons in your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		coupons, _, err := apiClient.CouponCodesAPI.
			ListCouponCodes(ctx).
			Execute()
		if err != nil {
			log.Fatalf("Error listing coupons: %v", err)
			return
		}

		fmt.Printf("\n%s\n\n", color.CyanString("🎟️  Coupons"))
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tCode\tType\tDiscount\tRemaining\tExpiry")

		for _, coupon := range coupons.Data {
			discount := ""

			// Determine discount type
			if *coupon.Type.Ptr() == "percentage" {
				discount = fmt.Sprintf("%.2f%%", *coupon.DiscountPercent)
			} else if *coupon.Type.Ptr() == "fixed_amount" {
				discount = fmt.Sprintf("$%.2f", *coupon.DiscountAmount)
			}

			// Calculate relative expiry time
			expiryDate := coupon.ExpiryDate // Assuming `ExpiryDate` is of type time.Time
			var relativeTime string
			if expiryDate != nil && !expiryDate.IsZero() { // Check if expiryDate is not zero
				relativeTime = formatting.FormatRelativeTimeWithExpired(*expiryDate) // Directly use expiryDate
			} else {
				relativeTime = "No expiry"
			}

			// Print coupon details
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%.2f\t%s\n",
				*coupon.Id, *coupon.Code, *coupon.Type.Ptr(), discount, float64(*coupon.Remaining), relativeTime)
		}

		w.Flush()
	},
}

func init() {
	couponCmd.AddCommand(couponListCmd)
}

/*
go run main.go coupon list
*/
