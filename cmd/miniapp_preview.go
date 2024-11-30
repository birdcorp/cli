package cmd

import (
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/open"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var miniappPreviewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview a miniapp",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for URL
		urlPrompt := promptui.Prompt{
			Label:    "URL",
			Validate: func(s string) error { return nil },
		}
		url, err := urlPrompt.Run()
		if err != nil {
			log.Fatalf("Error getting URL input: %v", err)
		}

		// Prompt for name
		namePrompt := promptui.Prompt{
			Label:    "Name (optional)",
			Validate: func(s string) error { return nil },
		}
		name, err := namePrompt.Run()
		if err != nil {
			log.Fatalf("Error getting name input: %v", err)
		}

		ctx, apiClient := auth.MustGetAuth()

		preview, resp, err := apiClient.MiniprogramAPI.
			CreateMiniProgramPreview(ctx).
			CreateMiniProgramPreviewRequest(birdsdk.CreateMiniProgramPreviewRequest{
				Url:  url,
				Name: name,
			}).
			Execute()
		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		if preview.Link != nil {
			open.Browser(*preview.Link)
		}
	},
}

func init() {
	// No flags needed since we're using prompts
}

// go run main.go miniapp create-preview <appID> --url <url>
//

/*
go run main.go miniapp create-preview \
 --url https://miniapp-developer.onrender.com/ \
 --name "Miniprogram Developer"
*/
