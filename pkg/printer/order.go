package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/formatting"
	"github.com/fatih/color"
)

func Order(order *birdsdk.Order) {
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
