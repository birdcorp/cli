package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook subcommand
var webhookCmd = &cobra.Command{
	Use:   "webhook",
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

	webhookCmd.SetHelpTemplate(`
$ birdcli ` + color.YellowString("webhooks") + `

🔔 Create, manage and delete webhooks.

` + color.GreenString("COMMANDS:") + `
  birdcli ` + color.YellowString("webhooks create") + `        ➕ Create a new webhook
  birdcli ` + color.YellowString("webhooks list") + `          📋 List all webhooks
  birdcli ` + color.YellowString("webhooks delete <id>") + `   🗑️  Delete a webhook

Use "birdcli [command] --help" for more information about a command.
`)
}
