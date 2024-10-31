package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var getMiniprogramCmd = &cobra.Command{
	Use:   "get <appID>",
	Short: "Get a miniprogram",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		appID := args[0] // Retrieve the appID argument

		ctx, apiClient := mustGetAuth()

		miniprogram, _, err := apiClient.MiniprogramAPI.
			GetMiniprogram(ctx, appID).
			Execute()
		if err != nil {
			log.Fatalf("Error getting miniprogram: %v", err)
		}

		prettyprint.JSON(miniprogram)
	},
}

// go run main.go miniprogram init
