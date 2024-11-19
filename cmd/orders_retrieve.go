package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

// getCmd represents the get subcommand
var getCmd = &cobra.Command{
	Use:   "retrieve <id>",
	Short: "Retrieve an order by ID",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		orderID := args[0] // Retrieve the orderID argument

		ctx, apiClient := auth.MustGetAuth()

		order, _, err := apiClient.OrdersAPI.
			GetOrder(ctx, orderID).
			Execute()
		if err != nil {
			log.Fatalf("Error listing orders: %v", err)
		}

		prettyprint.JSON(order)
	},
}
