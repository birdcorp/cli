package cmd

import (
	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook subcommand
var webhookCmd = &cobra.Command{
	Use:   "webhooks",
	Short: "Manage webhooks",
}

func init() {
	// Add the subcommands to the orders command
	webhookCmd.AddCommand(createWebhookCmd)
	//webhookCmd.AddCommand(getCmd)
	webhookCmd.AddCommand(listWebhooksCmd)
	webhookCmd.AddCommand(deleteWebhookCmd)

	// Add the orders command to the root command
	RootCmd.AddCommand(webhookCmd)
}
