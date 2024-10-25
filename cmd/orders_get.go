package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get subcommand
var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an order by ID",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		if !checkAPIKey() {
			fmt.Println("API key is not set. Please set it using the auth command.")
			return
		}

		id := args[0] // Get the order ID from the argument
		// Logic to get the order by ID
		fmt.Printf("Getting order with ID: %s\n", id)
	},
}
