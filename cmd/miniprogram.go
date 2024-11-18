package cmd

import (
	"github.com/spf13/cobra"
)

// ordersCmd represents the orders command
var miniprogramCmd = &cobra.Command{
	Use:   "miniprograms",
	Short: "Manage miniprogram",
	Long:  `Create, get, and list miniprogram.`,
}

func init() {
	// Add the subcommands to the orders command

	miniprogramCmd.AddCommand(createMiniprogramCmd)
	miniprogramCmd.AddCommand(listMiniprogramCmd)
	miniprogramCmd.AddCommand(deleteMiniprogramCmd)
	miniprogramCmd.AddCommand(publishMiniprogramCmd)
	miniprogramCmd.AddCommand(getMiniprogramCmd)
	miniprogramCmd.AddCommand(miniprogramPreviewCmd)
	miniprogramCmd.AddCommand(getMiniprogramInfoCmd)
	miniprogramCmd.AddCommand(miniprogramReleasesListCmd)

	// Add the orders command to the root command
	RootCmd.AddCommand(miniprogramCmd)

	miniprogramCmd.SetHelpTemplate(`
$ birdcli miniprograms

📲 The miniprograms command allows you to manage your Bird miniprograms.
You can create, publish, preview and manage your miniprograms. This includes
uploading new versions, managing releases, and getting build information.

COMMANDS:
  birdcli miniprograms preview       📱 Preview a miniprogram

  birdcli miniprograms init          🔧 Initialize miniprogram config
  birdcli miniprograms create        ➕ Create a miniprogram
  birdcli miniprograms delete <id>   🗑️  Delete a miniprogram
  birdcli miniprograms list          📋 List miniprograms
  birdcli miniprograms publish <id>  🚀 Publish a miniprogram
  birdcli miniprograms get <id>      🔍 Get a miniprogram
  birdcli miniprograms info <id>     ℹ️  Get miniprogram info
  birdcli miniprograms releases list 📦 List miniprogram releases

Use "birdcli [command] --help" for more information about a command.
`)
}

// go run main.go miniprogram list
// go run main.go miniprogram upload <appID>
// go run main.go miniprogram releases list
