package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list subcommand
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all orders",
	Run: func(cmd *cobra.Command, args []string) {
		if !checkAPIKey() {
			fmt.Println("API key is not set. Please set it using the auth command.")
			return
		}

		// Logic to list all orders
		fmt.Println("Listing all orders...")
	},
}
