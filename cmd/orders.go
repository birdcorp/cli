package cmd

import "github.com/spf13/cobra"

// ordersCmd represents the orders command
var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Manage orders",
	Long:  `Create, get, and list orders.`,
}

func init() {
	// Add the subcommands to the orders command
	ordersCmd.AddCommand(createCmd)
	ordersCmd.AddCommand(getCmd)
	ordersCmd.AddCommand(listCmd)

	// Add the orders command to the root command
	RootCmd.AddCommand(ordersCmd)
}

// go run main.go orders create
// go run main.go orders get <id>
// go run main.go orders list
