package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list subcommand
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all orders",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		orders, _, err := apiClient.OrdersAPI.
			ListOrders(ctx).
			Execute()
		if err != nil {
			log.Fatalf("Error listing orders: %v", err)
		}

		printJSON(orders)
	},
}
