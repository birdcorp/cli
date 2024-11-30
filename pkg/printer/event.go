package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/fatih/color"
)

func Event(event *birdsdk.WebhookEvent) {
	// Map event types to emojis and colors for visual distinction
	emoji := "📡"
	var eventColor color.Attribute

	switch event.Type {
	case birdsdk.WEBHOOK_SHIPPING_ADDRESS_CHANGE:
		emoji = "📍"
		eventColor = color.FgBlue
	case birdsdk.WEBHOOK_SHIPPING_METHOD_CHANGE:
		emoji = "🚚"
		eventColor = color.FgYellow
	case birdsdk.WEBHOOK_COUPON_CHANGE:
		emoji = "🎟️"
		eventColor = color.FgMagenta
	case birdsdk.WEBHOOK_APPROVED:
		emoji = "✅"
		eventColor = color.FgGreen
	case birdsdk.WEBHOOK_CAPTURED:
		emoji = "💳"
		eventColor = color.FgCyan
	case birdsdk.WEBHOOK_SETTLED:
		emoji = "💰"
		eventColor = color.FgHiGreen
	default:
		eventColor = color.FgWhite
	}

	// Create styled event header with emoji and type
	eventStyle := color.New(eventColor).Add(color.Bold)
	fmt.Printf("\n%s %s\n\n",
		emoji,
		eventStyle.Sprint(string(event.Type)))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Event ID:"), color.GreenString(event.Id))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Created At:"), color.BlueString(event.CreatedAt.Format("2006-01-02 15:04:05")))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Updated At:"), color.BlueString(event.UpdatedAt.Format("2006-01-02 15:04:05")))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Webhook ID:"), color.YellowString(event.WebhookId))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Type:"), color.MagentaString(string(event.Type)))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Status:"), color.CyanString(event.Status))

	if event.Attempts != nil {
		fmt.Fprintf(w, "%s\t%d\n", color.HiBlackString("Attempts:"), *event.Attempts)
	}

	if event.Data.OrderId != nil {
		fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Order ID:"), color.GreenString(*event.Data.OrderId))
	}

	if event.Data.PostalCode != nil {
		fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Postal Code:"), *event.Data.PostalCode)
	}

	if event.Data.ShippingMethod != nil {
		fmt.Fprintf(w, "%s\t\n", color.HiBlackString("Shipping Method:"))
		fmt.Fprintf(w, "  %s\t%s\n", color.HiBlackString("Label:"), event.Data.ShippingMethod.GetLabel())
		fmt.Fprintf(w, "  %s\t%s\n", color.HiBlackString("Amount:"), color.GreenString(event.Data.ShippingMethod.GetAmount()))
		fmt.Fprintf(w, "  %s\t%s\n", color.HiBlackString("Detail:"), event.Data.ShippingMethod.GetDetail())
		fmt.Fprintf(w, "  %s\t%s\n", color.HiBlackString("Identifier:"), event.Data.ShippingMethod.GetIdentifier())
		fmt.Fprintf(w, "  %s\t%v\n", color.HiBlackString("Selected:"), event.Data.ShippingMethod.GetSelected())
	}

	if event.Data.CouponCode != nil {
		fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Coupon Code:"), color.YellowString(*event.Data.CouponCode))
	}

	w.Flush()
	fmt.Println()
}
