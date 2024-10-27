package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var uploadMiniprogramCmd = &cobra.Command{
	Use:   "upload <appID>",
	Short: "Upload a miniprogram",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		appID := args[0] // Retrieve the appID argument

		// Get authentication context and client
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		miniprogram, _, err := apiClient.MiniprogramAPI.
			GetMiniprogram(ctx, appID).
			Execute()
		if err != nil {
			log.Fatalf("Error deleting miniprogram: %v", err)
		}

		printJSON(miniprogram)

		log.Println("Uploading miniprogram with ID:", appID)
	},
}

// go run main.go miniprogram upload <appID>
