package cmd

import (
	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/spf13/cobra"
)

var getMiniappCmd = &cobra.Command{
	Use:   "get <appID>",
	Short: "Get a miniapp",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appID := args[0]

		ctx, apiClient := auth.MustGetAuth()

		miniapp, resp, err := apiClient.MiniprogramAPI.
			GetMiniprogram(ctx, appID).
			Execute()
		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		printer.Miniapp(miniapp)
	},
}

// go run main.go miniapps get <appID>
