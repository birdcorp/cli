package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var webhookID string

var deleteWebhookCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a webhook",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		_, err = apiClient.WebhooksAPI.
			DeleteWebhook(ctx, webhookID).
			Execute()
		if err != nil {
			log.Fatalf("Error deleting webhook: %v", err)
		}

		fmt.Println("Webhook deleted successfully")
	},
}

func init() {

	deleteWebhookCmd.Flags().StringVarP(&webhookID, "id", "i", "", "ID of the webhook to delete (required)")
	deleteWebhookCmd.MarkFlagRequired("id")
}

// go run main.go webhook delete --id 1234567890
