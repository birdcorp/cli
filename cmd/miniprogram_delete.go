package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/spf13/cobra"
)

// deleteMiniprogramCmd represents the delete command
var deleteMiniprogramCmd = &cobra.Command{
	Use:   "delete <appID>",
	Short: "Delete a miniprogram",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		appID := args[0] // Get the app ID from the arguments
		log.Println("Deleting miniprogram with ID:", appID)

		resp, err := apiClient.MiniprogramAPI.
			DeleteMiniprogram(ctx, appID).
			Execute()

		if err != nil {
			log.Fatalf("Error deleting miniprogram: %v", err)
		}

		if resp.StatusCode != 204 {
			log.Fatalf("Unexpected status code: %d", resp.StatusCode)
		}

		log.Println("Miniprogram deleted successfully.")
	},
}

// go run main.go miniprogram delete <appID>
