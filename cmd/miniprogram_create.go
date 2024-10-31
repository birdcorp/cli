package cmd

import (
	"fmt"
	"io"
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var createMiniprogramCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a miniprogram",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		miniprogram, httpRes, err := apiClient.MiniprogramAPI.
			CreateMiniprogram(ctx).
			MiniprogramCreatePayload(birdsdk.MiniprogramCreatePayload{
				Name:               "my go miniprogram",
				Description:        "description",
				Url:                "https://example.com",          // Provide a valid URL
				BackgroundColor:    "#FFFFFF",                      // Example color code
				ForegroundColor:    "#000000",                      // Example color code
				IconImage:          "https://example.com/icon.png", // Provide a valid image URL
				NavBackgroundColor: "#F0F0F0",                      // Example color code
				NavTextColor:       birdsdk.DARK,                   // dark / light
				Tags:               []string{"tag1", "tag2"},       // Uncomment and provide tags if needed
			}).
			Execute()
		if err != nil {
			if httpRes != nil {
				if httpRes.Body != nil {
					body, err := io.ReadAll(httpRes.Body)
					if err == nil {
						fmt.Println(string(body)) // Print the HTTP response body for error details
					} else {
						fmt.Println("Error reading response body:", err)
					}
				}
			}
			log.Fatalf("Error creating miniprogram: %v", err)
		}

		prettyprint.JSON(miniprogram)

	},
}

// go run main.go miniprogram create
