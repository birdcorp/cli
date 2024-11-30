package cmd

import (
	"context"
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/auth"

	"github.com/birdcorp/cli/pkg/printer"
	"github.com/spf13/cobra"
)

func handleCouponGet(ctx context.Context, apiClient *birdsdk.APIClient, identifier string) {
	coupon, resp, err := apiClient.CouponCodesAPI.
		GetCouponCode(ctx, identifier).
		Execute()
	if err != nil {
		if resp != nil {
			//prettyprint.JSON(resp.Body)
		}
		log.Fatalf("Error getting coupon: %v", err)
	}

	printer.CouponDetails(coupon)
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
				//prettyprint.JSON(resp.Body)
				log.Fatalf("Error getting order (HTTP %d): %v", resp.StatusCode, err)
			}
		} else {
			log.Fatalf("Error connecting to API: %v", err)
		}
	}
	if order == nil {
		log.Fatalf("Unexpected error: Order response was empty")
	}
	printer.Order(order)
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
