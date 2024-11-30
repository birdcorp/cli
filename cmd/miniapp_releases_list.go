package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/miniapp"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/spf13/cobra"
)

var miniappReleasesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List miniapp releases",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		config, err := miniapp.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		releases, resp, err := apiClient.MiniprogramAPI.
			ListMiniprogramReleases(ctx, config.AppInfo.AppID).
			Execute()
		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		printer.MiniappReleases(releases.Data)
	},
}

func init() {
	miniappReleasesCmd.AddCommand(miniappReleasesListCmd)
}

//	go run main.go miniapp releases list
