package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "birdcli",
	Short: "Bird CLI - A command line interface for managing BirdPay merchant services",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(getCmd)

	RootCmd.SetHelpTemplate(`
Before using the CLI, you'll need to login:

$ birdcli login

Usage:
  birdcli [command]

COMMANDS:
  birdcli account        👨‍💼 Manage account
  birdcli coupon         🏷️  Manage coupons 
  birdcli events         📡 Manage events
  birdcli help           💡 Get help for any command
  birdcli login          🔐 Set the API key
  birdcli logout         🚫 Logout from the CLI
  birdcli miniprograms   📲 Manage miniprograms
  birdcli orders         🛒 Manage orders
  birdcli webhooks       🪝 Manage webhooks

Use "birdcli [command] --help" for more information about a command.
`)

	RootCmd.SetUsageTemplate(`
Usage:
  birdcli [command]

Available Commands:
{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding}} {{.Short}}{{end}}{{end}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Use "birdcli [command] --help" for more information about a command.
`)
}
