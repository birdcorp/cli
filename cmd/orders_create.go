package cmd

import (
	"encoding/json"
	"log"
	"strings"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/open"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var (
	totalValue         string
	currency           string
	lineItemsJSON      string
	requiredShipFields string
	requiredBillFields string
)

// createCmd represents the create subcommand
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new order",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		// Parse the required fields into slices
		shipFields := strings.Split(requiredShipFields, ",")
		billFields := strings.Split(requiredBillFields, ",")

		// Parse the line items JSON string
		var lineItems []birdsdk.LineItem
		if err := json.Unmarshal([]byte(lineItemsJSON), &lineItems); err != nil {
			log.Fatalf("Failed to parse line items JSON: %v", err)
		}

		// Construct the order payload
		orderPayload := birdsdk.OrderPayload{
			Total: &birdsdk.Total{
				Value:    totalValue,
				Currency: birdsdk.Currency(currency),
			},
			LineItems:              lineItems,
			RequiredShippingFields: shipFields,
			RequiredBillingFields:  billFields,
		}

		// Call the API to create the order
		response, httpRes, err := apiClient.OrdersAPI.
			CreateOrder(ctx).
			OrderPayload(orderPayload).
			Execute()
		if err != nil {
			if httpRes != nil {
				log.Printf("Error creating order: %s\nHTTP Status: %s", err.Error(), httpRes.Status)
				var errorResponse map[string]interface{}
				if decodeErr := json.NewDecoder(httpRes.Body).Decode(&errorResponse); decodeErr == nil {
					prettyprint.JSON(errorResponse)
				} else {
					log.Println("Could not decode error response:", decodeErr)
				}
			} else {
				log.Printf("Error creating order: %s", err.Error())
			}
			log.Fatal("Order creation failed")
		}

		prettyprint.JSON(response)

		open.Browser(response.Link)
	},
}

func init() {
	createCmd.Flags().StringVar(&totalValue, "total-value", "1.99", "Total amount for the order")
	createCmd.Flags().StringVar(&currency, "currency", "USD", "Currency for the order")
	createCmd.Flags().StringVar(&lineItemsJSON, "line-items", "[]", "JSON string of line items") // New flag
	createCmd.Flags().StringVar(&requiredShipFields, "required-shipping-fields", "name,postalAddress,phone,email", "Comma-separated list of required shipping fields")
	createCmd.Flags().StringVar(&requiredBillFields, "required-billing-fields", "name,postalAddress,phone,email", "Comma-separated list of required billing fields")

	// Add createCmd to the root command
}

/*
go run main.go orders create \
  --total-value "10.99" \
  --currency "USD" \
  --line-items '[
    {
      "label": "Item1",
      "value": "5.99",
      "status": "final",
      "type": "item"
    },
    {
      "label": "Item2",
      "value": "5.00",
      "status": "pending",
      "type": "tax"
    },
    {
      "label": "Shipping",
      "value": "0.00",
      "status": "pending",
      "type": "shipping"
    }
  ]' \
  --required-shipping-fields "name,postalAddress,phone,email" \
  --required-billing-fields "name,postalAddress,phone,email"
*/
