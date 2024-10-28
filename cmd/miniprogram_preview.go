package cmd

import (
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/open"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var miniprogramPreviewCmd = &cobra.Command{
	Use:   "create-preview <appID> --url <url>",
	Short: "Preview a miniprogram",
	Args:  cobra.NoArgs, // No arguments are required
	Run: func(cmd *cobra.Command, args []string) {

		// Get the URL from the flags
		url, _ := cmd.Flags().GetString("url")
		name, _ := cmd.Flags().GetString("name")

		// Get authentication context and client
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		preview, _, err := apiClient.MiniprogramAPI.
			CreateMiniProgramPreview(ctx).
			CreateMiniProgramPreviewRequest(birdsdk.CreateMiniProgramPreviewRequest{
				Url:  url,
				Name: name,
			}).
			Execute()
		if err != nil {
			log.Fatalf("Error creating miniprogram preview: %v", err)
		}

		prettyprint.JSON(preview)

		if preview.Link != nil {
			open.Browser(*preview.Link)
		}
	},
}

func init() {
	miniprogramPreviewCmd.Flags().String("url", "", "URL for the miniprogram preview")
	miniprogramPreviewCmd.Flags().String("name", "", "Name for the miniprogram preview")
	miniprogramPreviewCmd.MarkFlagRequired("url")
}

// go run main.go miniprogram create-preview <appID> --url <url>
//

/*
go run main.go miniprogram create-preview \
 --url https://miniprogram-developer.onrender.com/ \
 --name "Miniprogram Developer"
*/
