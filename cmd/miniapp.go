package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var miniappCmd = &cobra.Command{
	Use:   "miniapp",
	Short: "Manage miniapp",
	Long:  `Create, get, and list miniapp.`,
}

func init() {
	// Add the subcommands to the orders command
	miniappCmd.AddCommand(createMiniappCmd)
	miniappCmd.AddCommand(listMiniappCmd)
	miniappCmd.AddCommand(deleteMiniappCmd)
	miniappCmd.AddCommand(publishMiniappCmd)
	miniappCmd.AddCommand(getMiniappCmd)
	miniappCmd.AddCommand(miniappPreviewCmd)
	miniappCmd.AddCommand(getMiniappInfoCmd)
	miniappCmd.AddCommand(miniappReleasesListCmd)

	// Add the orders command to the root command
	RootCmd.AddCommand(miniappCmd)

	miniappCmd.SetHelpTemplate(`
$ birdcli ` + color.YellowString("miniapp") + `

📲 Create, manage and publish miniapps for your Bird store.

` + color.GreenString("COMMANDS:") + `
  birdcli ` + color.YellowString("miniapp preview") + `         📱 Create a preview link for a miniapp staging URL
  birdcli ` + color.YellowString("miniapp init") + `            🔧 Initialize miniapp config
  birdcli ` + color.YellowString("miniapp create") + `          ➕ Create a miniapp
  birdcli ` + color.YellowString("miniapp delete <id>") + `     🗑️  Delete a miniapp
  birdcli ` + color.YellowString("miniapp list") + `            📋 List all miniapps for your merchantID account
  birdcli ` + color.YellowString("miniapp publish") + `         🚀 Publish a miniapp release for production review
  birdcli ` + color.YellowString("miniapp get <id>") + `        🔍 Get a miniapp
  birdcli ` + color.YellowString("miniapp info <id>") + `       ℹ️  Get miniapp info
  birdcli ` + color.YellowString("miniapp releases list") + `   📦 List miniapp releases

Use "birdcli [command] --help" for more information about a command.
`)
}
