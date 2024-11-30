package cmd

import (
	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/spf13/cobra"
)

var listWebhooksCmd = &cobra.Command{
	Use:   "list",
	Short: "List webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		webhooks, resp, err := apiClient.WebhooksAPI.
			ListWebhooks(ctx).
			Execute()
		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		printer.WebhookList(webhooks.Data)
	},
}

// go run main.go webhook list
