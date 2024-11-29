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

var listEventsCmd = &cobra.Command{
	Use:   "list",
	Short: "List events",
	Long:  `List all events in your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		events, _, err := apiClient.EventsAPI.
			ListEvents(ctx).
			Execute()
		if err != nil {
			log.Fatalf("Error listing events: %v", err)
			return
		}

		fmt.Printf("\n%s\n\n", color.CyanString("📋 Events"))
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "Created\tID\tType\tStatus\tOrder ID\tPostal Code\tShipping Method\tCoupon Code\tUpdated")

		for _, event := range events.Data {
			var orderId, postalCode, shippingMethod, couponCode string

			if event.Data.OrderId != nil {
				orderId = *event.Data.OrderId
			}
			if event.Data.PostalCode != nil {
				postalCode = *event.Data.PostalCode
			}
			if event.Data.ShippingMethod != nil {
				shippingMethod = event.Data.ShippingMethod.GetIdentifier()
			}
			if event.Data.CouponCode != nil {
				couponCode = *event.Data.CouponCode
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				formatting.FormatRelativeTime(event.GetCreatedAt()),
				event.GetId(),
				event.GetType(),
				event.GetStatus(),
				orderId,
				postalCode,
				shippingMethod,
				couponCode,
				formatting.FormatRelativeTime(event.GetUpdatedAt()),
			)
		}

		w.Flush()
	},
}

func init() {
	eventsCmd.AddCommand(listEventsCmd)
}

/*
go run main.go events list
*/
