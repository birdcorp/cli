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
		printer.HandleAPIFailure(resp)
		return
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
