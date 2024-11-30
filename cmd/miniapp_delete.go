package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/spf13/cobra"
)

// deleteMiniappCmd represents the delete command
var deleteMiniappCmd = &cobra.Command{
	Use:   "delete <appID>",
	Short: "Delete a miniapp",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		appID := args[0]
		log.Println("Deleting miniapp with ID:", appID)

		resp, err := apiClient.MiniprogramAPI.
			DeleteMiniprogram(ctx, appID).
			Execute()

		if err != nil {
			log.Fatalf("Error deleting miniapp: %v", err)
		}

		if resp.StatusCode != 204 {
			log.Fatalf("Unexpected status code: %d", resp.StatusCode)
		}

		log.Println("Miniapp deleted successfully.")
	},
}

// go run main.go miniapp delete <appID>
