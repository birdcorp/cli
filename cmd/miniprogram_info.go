package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/miniprogram"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var getMiniprogramInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get a miniprogram info",
	Args:  nil, // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		config, err := miniprogram.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		miniprogram, _, err := apiClient.MiniprogramAPI.
			GetMiniprogram(ctx, config.AppInfo.AppID).
			Execute()
		if err != nil {
			log.Fatalf("Error getting miniprogram: %v", err)
		}
		prettyprint.JSON(miniprogram)

		releases, _, err := apiClient.MiniprogramAPI.
			ListMiniprogramReleases(ctx, config.AppInfo.AppID).
			Execute()
		if err != nil {
			log.Fatalf("Error getting miniprogram releases: %v", err)
		}

		prettyprint.JSON(releases)
	},
}

// go run main.go miniprogram info
