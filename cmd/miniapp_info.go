package cmd

import (
	"fmt"
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/miniapp"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var getMiniappInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get a miniapp info",
	Args:  nil,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		config, err := miniapp.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		miniapp, resp, err := apiClient.MiniprogramAPI.
			GetMiniprogram(ctx, config.AppInfo.AppID).
			Execute()
		if err != nil {
			fmt.Printf("%s Failed to get miniapp:\n", color.RedString("❌"))

			printer.HandleAPIFailure(resp)
			return
		}
		printer.Miniapp(miniapp)

		/*
			releases, _, err := apiClient.MiniprogramAPI.
				ListMiniprogramReleases(ctx, config.AppInfo.AppID).
				Execute()
			if err != nil {
				log.Fatalf("Error getting miniapp releases: %v", err)
			}
		*/

	},
}

// go run main.go miniapp info
