package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/miniprogram"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var getMiniprogramInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get a miniprogram info",
	Args:  nil, // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		config, err := miniprogram.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		miniprogram, resp, err := apiClient.MiniprogramAPI.
			GetMiniprogram(ctx, config.AppInfo.AppID).
			Execute()
		if err != nil {
			fmt.Printf("%s Failed to get miniprogram:\n", color.RedString("❌"))

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading response body: %v", err)
			}
			fmt.Println("Response Body:", string(body))
			return
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
