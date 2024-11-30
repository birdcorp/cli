package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/formatting"
	"github.com/fatih/color"
)

func CouponDetails(coupon *birdsdk.CouponCode) {
	fmt.Printf("\n%s\n\n", color.CyanString("🎟️  Coupon Details"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tCode\tType\tDiscount\tRemaining\tExpiry")

	discount := ""
	// Determine discount type
	if *coupon.Type.Ptr() == "percentage" {
		discount = fmt.Sprintf("%.2f%%", *coupon.DiscountPercent)
	} else if *coupon.Type.Ptr() == "fixed_amount" {
		discount = fmt.Sprintf("$%.2f", *coupon.DiscountAmount)
	}

	// Calculate relative expiry time
	expiryDate := coupon.ExpiryDate
	var relativeTime string
	if expiryDate != nil && !expiryDate.IsZero() {
		relativeTime = formatting.FormatRelativeTimeWithExpired(*expiryDate)
	} else {
		relativeTime = "No expiry"
	}

	// Print coupon details
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%.2f\t%s\n",
		*coupon.Id, *coupon.Code, *coupon.Type.Ptr(), discount, float64(*coupon.Remaining), relativeTime)

	w.Flush()
	fmt.Println()
}

func CouponsList(coupons []birdsdk.CouponCode) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Printf("\n%s\n\n", color.CyanString("🎟️  Coupons"))
	fmt.Fprintln(w, "ID\tCode\tType\tDiscount\tRemaining\tExpiry")

	for _, coupon := range coupons {
		discount := ""

		// Determine discount type
		if coupon.GetType() == "percentage" {
			discount = fmt.Sprintf("%.2f%%", *coupon.DiscountPercent)
		} else if coupon.GetType() == "fixed_amount" {
			discount = fmt.Sprintf("$%.2f", *coupon.DiscountAmount)
		}

		// Calculate relative expiry time
		expiryDate := coupon.GetExpiryDate()
		var relativeTime string
		if !expiryDate.IsZero() {
			relativeTime = formatting.FormatRelativeTimeWithExpired(expiryDate)
		} else {
			relativeTime = "No expiry"
		}

		// Print coupon details
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%.2f\t%s\n",
			*coupon.Id,
			color.YellowString(coupon.GetCode()),
			coupon.GetType(),
			discount,
			float64(coupon.GetRemaining()),
			relativeTime,
		)
	}

	w.Flush()
}
