package cmd

import (
	"fmt"
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/spf13/cobra"
)

// createCmd represents the create subcommand
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new order",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		/*
		* Create order
		 */
		orderPayload := birdsdk.OrderPayload{
			Total: &birdsdk.Total{
				Value:    "1.99",
				Currency: birdsdk.USD,
			},
			LineItems: []birdsdk.LineItem{
				{
					Label:        "Test Item",
					Value:        "1.99",
					Status:       birdsdk.PtrString("final"),
					Type:         birdsdk.PtrString("item"),
					ThumbnailUrl: birdsdk.PtrString("https://theiphoneway.com/wp-content/uploads/2023/10/iPhone-14-Pro-Max-9907.jpg"),
				},
				{
					Label:  "Tax",
					Value:  "0.00",
					Status: birdsdk.PtrString("pending"),
					Type:   birdsdk.PtrString("tax"),
				},
				{
					Label:  "Shipping",
					Value:  "0.00",
					Status: birdsdk.PtrString("pending"),
					Type:   birdsdk.PtrString("shipping"),
				},
			},
			RequiredShippingFields: []string{"name", "postalAddress", "phone", "email"},
			RequiredBillingFields:  []string{"name", "postalAddress", "phone", "email"},
		}

		response, _, err := apiClient.OrdersAPI.
			CreateOrder(ctx).
			OrderPayload(orderPayload).
			Execute()

		if err != nil {
			log.Fatal(err)
		}

		printJSON(response)

		// Logic to create an order
		fmt.Println("Creating a new order...")
	},
}
