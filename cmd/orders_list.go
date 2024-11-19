package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/formatting"
	"github.com/spf13/cobra"
)

// listCmd represents the list subcommand
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all orders",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		orders, _, err := apiClient.OrdersAPI.
			ListOrders(ctx).
			Execute()
		if err != nil {
			log.Fatalf("Error listing orders: %v", err)
		}

		if orders == nil || orders.Data == nil {
			log.Println("No orders found or unable to retrieve orders data.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tStatus\tType\tTotal\tCreated At\tUpdated At")

		for _, order := range orders.Data {
			fmt.Fprintf(w, "%s\t%s\t%s\t%.2f\t%s\t%s\n",
				order.Id,
				order.Status,
				*order.Type.Get().Ptr(),
				formatting.ParseFloat(order.Total.GetValue()),
				formatting.FormatRelativeTime(order.CreatedAt),
				formatting.FormatRelativeTime(order.UpdatedAt),
			)
		}

		w.Flush()
	},
}

// go run main.go orders list
