package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/formatting"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func handleCouponGet(ctx context.Context, apiClient *birdsdk.APIClient, identifier string) {
	coupon, resp, err := apiClient.CouponCodesAPI.
		GetCouponCode(ctx, identifier).
		Execute()
	if err != nil {
		if resp != nil {
			prettyprint.JSON(resp.Body)
		}
		log.Fatalf("Error getting coupon: %v", err)
	}

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

func handleOrderGet(ctx context.Context, apiClient *birdsdk.APIClient, identifier string) {
	order, resp, err := apiClient.OrdersAPI.
		GetOrder(ctx, identifier).
		Execute()
	if err != nil {
		if resp != nil {
			switch resp.StatusCode {
			case 404:
				log.Fatalf("Order not found: %s", identifier)
			case 400:
				log.Fatalf("Invalid order ID format: %s", identifier)
			case 401:
				log.Fatalf("Unauthorized: Please check your API credentials")
			case 403:
				log.Fatalf("Forbidden: You don't have permission to access this order")
			default:
				prettyprint.JSON(resp.Body)
				log.Fatalf("Error getting order (HTTP %d): %v", resp.StatusCode, err)
			}
		} else {
			log.Fatalf("Error connecting to API: %v", err)
		}
	}
	if order == nil {
		log.Fatalf("Unexpected error: Order response was empty")
	}
	fmt.Printf("\n%s\n\n", color.CyanString("📦 Order Details"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Basic order info
	fmt.Fprintln(w, "ID\tStatus\tType\tTotal\tCreated\tLast Updated")
	fmt.Fprintf(w, "%s\t%s\t%s\t$%.2f\t%s\t%s\n\n",
		order.GetId(),
		order.GetStatus(),
		*order.Type.Get().Ptr(),
		formatting.ParseFloat(order.Total.GetValue()),
		formatting.FormatRelativeTime(order.GetCreatedAt()),
		formatting.FormatRelativeTime(order.GetUpdatedAt()),
	)

	// Line items
	fmt.Fprintf(w, "%s\n", color.CyanString("Line Items:"))
	fmt.Fprintln(w, "Label\tType\tValue\tSKU ID\tStatus")

	// First display items of type "item"
	for _, item := range order.LineItems {
		if item.GetType() == "item" {
			fmt.Fprintf(w, "%s\t%s\t$%s\t%s\t%s\n",
				item.GetLabel(),
				item.GetType(),
				item.GetValue(),
				item.GetSkuId(),
				item.GetStatus(),
			)
		}
	}

	// Then display remaining items
	for _, item := range order.LineItems {
		if item.GetType() != "item" {
			fmt.Fprintf(w, "%s\t%s\t$%s\t%s\t%s\n",
				item.GetLabel(),
				item.GetType(),
				item.GetValue(),
				item.GetSkuId(),
				item.GetStatus(),
			)
		}
	}
	fmt.Fprintln(w)

	// Shipping info if available
	if order.Shipping != nil {
		fmt.Fprintf(w, "%s\n", color.CyanString("Shipping Information:"))
		fmt.Fprintf(w, "Name:\t%s\n", order.Shipping.GetName())
		fmt.Fprintf(w, "Carrier:\t%s\n", order.Shipping.GetCarrier())
		fmt.Fprintf(w, "Phone:\t%s\n", order.Shipping.GetPhone())
		fmt.Fprintf(w, "Tracking #:\t%s\n", order.Shipping.GetTrackingNumber())

		if order.Shipping.Address != nil {
			fmt.Fprintf(w, "\n%s\n", color.CyanString("Shipping Address:"))
			fmt.Fprintf(w, "Street:\t%s\n", order.Shipping.Address.GetLine1())
			if order.Shipping.Address.GetLine2() != "" {
				fmt.Fprintf(w, "\t%s\n", order.Shipping.Address.GetLine2())
			}
			fmt.Fprintf(w, "City:\t%s\n", order.Shipping.Address.GetCity())
			fmt.Fprintf(w, "State:\t%s\n", order.Shipping.Address.GetState())
			fmt.Fprintf(w, "Postal Code:\t%s\n", order.Shipping.Address.GetPostalCode())
			fmt.Fprintf(w, "Country:\t%s\n", order.Shipping.Address.GetCountry())
		}
	}

	// Billing info if available
	if order.Billing != nil {
		fmt.Fprintf(w, "\n%s\n", color.CyanString("Billing Information:"))
		if order.Billing.Name != nil {
			fmt.Fprintf(w, "Name:\t%s\n", order.Billing.GetName())
		}
		fmt.Fprintf(w, "Email:\t%s\n", order.Billing.GetEmail())
		fmt.Fprintf(w, "Phone:\t%s\n", order.Billing.GetPhone())

		if order.Billing.Address != nil {
			fmt.Fprintf(w, "\n%s\n", color.CyanString("Billing Address:"))
			fmt.Fprintf(w, "Street:\t%s\n", order.Billing.Address.GetLine1())
			if order.Billing.Address.GetLine2() != "" {
				fmt.Fprintf(w, "\t%s\n", order.Billing.Address.GetLine2())
			}
			fmt.Fprintf(w, "City:\t%s\n", order.Billing.Address.GetCity())
			fmt.Fprintf(w, "State:\t%s\n", order.Billing.Address.GetState())
			fmt.Fprintf(w, "Postal Code:\t%s\n", order.Billing.Address.GetPostalCode())
			fmt.Fprintf(w, "Country:\t%s\n", order.Billing.Address.GetCountry())
		}
	}

	w.Flush()
	fmt.Println()
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <identifier>",
	Short: "Get a resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()
		identifier := args[0]

		switch {
		case len(identifier) >= 5 && identifier[:5] == "coup_":
			handleCouponGet(ctx, apiClient, identifier)
		case len(identifier) >= 5 && identifier[:5] == "ordr_":
			handleOrderGet(ctx, apiClient, identifier)
		default:
			log.Fatalf("Unrecognized identifier prefix: %s", identifier)
		}
	},
}

// go run main.go get <identifier>
