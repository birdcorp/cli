package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "birdcli",
	Short: "Bird CLI - Manage BirdPay merchant services",
	Long: `
Bird CLI is a command-line interface for managing BirdPay merchant services.

Before using the CLI, you'll need to authenticate with your BirdPay ` + color.YellowString("API key") + `:

  $ birdcli ` + color.YellowString("login") + `

Once logged in, you can manage your account, orders, webhooks, and more using the available commands.
`,
	// Display the help message when no subcommands are provided.
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the RootCmd and handles errors.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands to the RootCmd.
	RootCmd.AddCommand(getCmd) // Example command; replace with actual commands.

	// Set a custom help template for detailed usage instructions.
	RootCmd.SetHelpTemplate(`
{{.Long}}

` + color.BlueString("Usage:") + `
  birdcli [command]

` + color.GreenString("Available Commands:") + `
  birdcli ` + color.YellowString("version") + `        🔢 Print the CLI version
  birdcli ` + color.YellowString("login") + `          🔐 Authenticate with BirdPay
  birdcli ` + color.YellowString("logout") + `         🚫 Logout from the CLI
  birdcli ` + color.YellowString("account") + `        👨‍💼 Get merchant account info
  birdcli ` + color.YellowString("get") + `            🔍 Get a resource by id, order, coupon, event, etc.
  birdcli ` + color.YellowString("coupon") + `         🏷️  Manage coupons 
  birdcli ` + color.YellowString("event") + `          📡 Manage events
  birdcli ` + color.YellowString("help") + `           💡 Get help for any command
  birdcli ` + color.YellowString("miniapp") + `        📲 Manage mini-apps
  birdcli ` + color.YellowString("order") + `          🛒 Manage orders
  birdcli ` + color.YellowString("webhook") + `        🪝 Manage webhooks

` + color.CyanString("Flags:") + `
  -h, --help   Display help for birdcli or a command

Use "birdcli [command] --help" for more information about a command.
`)

	// Set a custom usage template for consistent formatting.
	RootCmd.SetUsageTemplate(`
` + color.BlueString("Usage:") + `
  birdcli [command]

` + color.GreenString("Available Commands:") + `
{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding}} {{.Short}}{{end}}{{end}}

` + color.CyanString("Flags:") + `
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Use "birdcli [command] --help" for more information about a command.
`)
}
