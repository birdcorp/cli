package cmd

import (
	"fmt"
	"log"

	"github.com/birdcorp/cli/pkg/miniprogram"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var miniprogramReleasesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List miniprogram releases",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		config, err := miniprogram.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		releases, _, err := apiClient.MiniprogramAPI.
			ListMiniprogramReleases(ctx, config.AppInfo.AppID).
			Execute()
		if err != nil {
			fmt.Println("Error listing miniprogram releases:", err)
			return
		}

		prettyprint.JSON(releases)
	},
}

func init() {
	miniprogramReleasesCmd.AddCommand(miniprogramReleasesListCmd)
}

//	go run main.go miniprograms releases list
