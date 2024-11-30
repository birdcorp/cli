package cmd

import (
	"fmt"
	"io"
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

		webhook, httpRes, err := apiClient.WebhooksAPI.
			CreateWebhook(ctx).
			CreateWebhookRequest(payload).
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

		printer.Webhook(webhook)

	},
}

func init() {
	createWebhookCmd.Flags().StringVarP(&webhookURL, "url", "u", "", "URL for the webhook")
}

// go run main.go webhook create --url https://www.example.com
