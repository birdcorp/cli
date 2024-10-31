package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// deleteMiniprogramCmd represents the delete command
var deleteMiniprogramCmd = &cobra.Command{
	Use:   "delete <appID>",
	Short: "Delete a miniprogram",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		appID := args[0] // Get the app ID from the arguments
		log.Println("Deleting miniprogram with ID:", appID)

		_, err := apiClient.MiniprogramAPI.
			DeleteMiniprogram(ctx, appID).
			Execute()
		if err != nil {
			log.Fatalf("Error deleting miniprogram: %v", err)
		}

		log.Println("Miniprogram deleted successfully.")
	},
}

// go run main.go miniprogram delete <appID>
