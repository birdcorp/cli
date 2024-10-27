package cmd

import (
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/open"
	"github.com/spf13/cobra"
)

var miniprogramPreviewCmd = &cobra.Command{
	Use:   "create-preview <appID> --url <url>",
	Short: "Preview a miniprogram",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		appID := args[0] // Retrieve the appID argument

		// Get the URL from the flags
		url, _ := cmd.Flags().GetString("url")

		// Get authentication context and client
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		miniprogram, _, err := apiClient.MiniprogramAPI.
			CreateMiniProgramPreview(ctx, appID).
			CreateMiniProgramPreviewRequest(birdsdk.CreateMiniProgramPreviewRequest{
				Url: url,
			}).
			Execute()
		if err != nil {
			log.Fatalf("Error creating miniprogram preview: %v", err)
		}

		printJSON(miniprogram)

		if miniprogram.Link != nil {
			open.Browser(*miniprogram.Link)
		}

		log.Println("Previewing miniprogram with ID:", appID)
	},
}

func init() {
	miniprogramPreviewCmd.Flags().String("url", "", "URL for the miniprogram preview")
	miniprogramPreviewCmd.MarkFlagRequired("url")
}

// go run main.go miniprogram create-preview <appID> --url <url>
// go run main.go miniprogram create-preview 17e9a4eaf2afd428 --url https://miniprogram-developer.onrender.com/
