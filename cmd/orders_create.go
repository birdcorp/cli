package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/open"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create subcommand
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new order",
	Long: `Create a new order with specified total amount, currency, line items and required fields.
	
Example:
  bird orders create --total-value "10.99" --currency "USD" --line-items '[{"label":"Item1","value":"10.99"}]'`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		// Parse and validate the line items JSON
		var lineItems []birdsdk.LineItem
		lineItemsStr := viper.GetString("line-items")
		if err := json.Unmarshal([]byte(lineItemsStr), &lineItems); err != nil {
			pterm.Error.Printf("Invalid line items JSON format: %v\n", err)
			log.Fatal("Please provide valid JSON for line items")
		}

		if len(lineItems) == 0 {
			pterm.Warning.Println("No line items provided")
		}

		// Validate total matches sum of line items
		totalValue := viper.GetString("total-value")
		total, err := strconv.ParseFloat(totalValue, 64)
		if err != nil {
			pterm.Error.Printf("Invalid total value format: %v\n", err)
			log.Fatal("Please provide a valid numeric total value")
		}

		var lineItemsTotal float64
		for _, item := range lineItems {
			value, err := strconv.ParseFloat(item.Value, 64)
			if err != nil {
				pterm.Error.Printf("Invalid line item value format: %v\n", err)
				log.Fatal("Please provide valid numeric values for line items")
			}
			lineItemsTotal += value
		}

		// Compare totals with 2 decimal precision
		if fmt.Sprintf("%.2f", total) != fmt.Sprintf("%.2f", lineItemsTotal) {
			pterm.Error.Printf("Total value (%.2f) does not match sum of line items (%.2f)\n", total, lineItemsTotal)
			log.Fatal("Please ensure total matches sum of line items")
		}

		// Validate currency
		currency := birdsdk.Currency(viper.GetString("currency"))

		// Construct order payload
		orderPayload := birdsdk.OrderPayload{
			Total: &birdsdk.Total{
				Value:    viper.GetString("total-value"),
				Currency: currency,
			},
			LineItems:              lineItems,
			RequiredShippingFields: strings.Split(viper.GetString("required-shipping-fields"), ","),
			RequiredBillingFields:  strings.Split(viper.GetString("required-billing-fields"), ","),
		}

		response, resp, err := apiClient.OrdersAPI.
			CreateOrder(ctx).
			OrderPayload(orderPayload).
			Execute()

		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		// Fetch and display the created order
		order, resp, err := apiClient.OrdersAPI.
			GetOrder(ctx, response.Id).
			Execute()
		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		printer.Order(order)

		// Prompt user to open order link in browser
		prompt := promptui.Prompt{
			Label:     color.CyanString("Press Enter to open order link in browser"),
			IsConfirm: true,
		}
		if _, err := prompt.Run(); err == nil {
			if err := open.Browser(response.Link); err != nil {
				pterm.Warning.Printf("Failed to open order link in browser: %v\n", err)
			}
		}
	},
}

func init() {
	createCmd.Flags().String("total-value", "1.99", "Total amount for the order (e.g., \"10.99\")")
	createCmd.Flags().String("currency", "USD", "Currency code for the order (e.g., USD, EUR, GBP)")
	createCmd.Flags().String("line-items", "[]", "JSON array of line items")
	createCmd.Flags().String("required-shipping-fields", "name,postalAddress,phone,email", "Comma-separated list of required shipping fields")
	createCmd.Flags().String("required-billing-fields", "name,postalAddress,phone,email", "Comma-separated list of required billing fields")

	for _, flag := range []string{
		"total-value",
		"currency",
		"line-items",
		"required-shipping-fields",
		"required-billing-fields",
	} {
		if err := viper.BindPFlag(flag, createCmd.Flags().Lookup(flag)); err != nil {
			log.Fatalf("Error binding flag %s: %v", flag, err)
		}
	}
}

// handleAPIError handles API errors with detailed output
func handleAPIError(err error, httpRes *http.Response) {
	if httpRes != nil {
		pterm.Error.Printf("API Error: %s\nHTTP Status: %s\n", err.Error(), httpRes.Status)
		var errorResponse map[string]interface{}
		if decodeErr := json.NewDecoder(httpRes.Body).Decode(&errorResponse); decodeErr == nil {
			//prettyprint.JSON(errorResponse)
		} else {
			pterm.Error.Println("Could not decode error response:", decodeErr)
		}
	} else {
		pterm.Error.Printf("Error: %s\n", err.Error())
	}
}

/*
Example usage:
go run main.go order create \
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
