package cmd

import (
	"fmt"
	"log"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var webhookURL string

var createWebhookCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a webhook",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		// If URL not provided via flag, prompt for it
		if webhookURL == "" {
			prompt := promptui.Prompt{
				Label: "Webhook URL",
				Validate: func(input string) error {
					if input == "" {
						return fmt.Errorf("URL cannot be empty")
					}
					return nil
				},
			}

			result, err := prompt.Run()
			if err != nil {
				log.Fatalf("Prompt failed: %v", err)
			}
			webhookURL = result
		}

		log.Println("Creating webhook.")

		payload := birdsdk.CreateWebhookRequest{
			Url: webhookURL,
		}

		webhook, resp, err := apiClient.WebhooksAPI.
			CreateWebhook(ctx).
			CreateWebhookRequest(payload).
			Execute()

		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		printer.Webhook(webhook)

	},
}

func init() {
	createWebhookCmd.Flags().StringVarP(&webhookURL, "url", "u", "", "URL for the webhook")
}

// go run main.go webhook create --url https://www.example.com
