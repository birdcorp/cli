package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/formatting"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list subcommand
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all orders",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Fetching orders..."
		s.Start()

		orders, _, err := apiClient.OrdersAPI.
			ListOrders(ctx).
			Execute()
		if err != nil {
			s.Stop()
			log.Fatalf("Error listing orders: %v", err)
		}

		s.Stop()

		if orders == nil || orders.Data == nil {
			log.Println("No orders found or unable to retrieve orders data.")
			return
		}

		fmt.Printf("\n%s\n\n", color.CyanString("📦 Orders"))
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tStatus\tType\tTotal\tCreated\tLast Updated")

		for _, order := range orders.Data {

			fmt.Fprintf(w, "%s\t%s\t%s\t$%.2f\t%s\t%s\n",
				order.GetId(),
				order.GetStatus(),
				*order.Type.Get().Ptr(),
				formatting.ParseFloat(order.Total.GetValue()),
				formatting.FormatRelativeTime(order.GetCreatedAt()),
				formatting.FormatRelativeTime(order.GetUpdatedAt()),
			)
		}

		w.Flush()
		fmt.Println()
	},
}

// go run main.go orders list
