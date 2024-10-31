package cmd

import (
	"fmt"
	"io"
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var webhookURL string

var createWebhookCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a webhook",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		log.Println("Creating webhook.")

		webhook, httpRes, err := apiClient.WebhooksAPI.
			CreateWebhook(ctx).
			CreateWebhookRequest(birdsdk.CreateWebhookRequest{
				Url: webhookURL,
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
			log.Fatalf("Error creating webhook: %v", err)
		}

		prettyprint.JSON(webhook)

	},
}

func init() {
	createWebhookCmd.Flags().StringVarP(&webhookURL, "url", "u", "", "URL for the webhook (required)")
	createWebhookCmd.MarkFlagRequired("url")
}

// go run main.go webhook create --url https://www.example.com
