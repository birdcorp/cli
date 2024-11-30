package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ordersCmd represents the orders command
var ordersCmd = &cobra.Command{
	Use:   "order",
	Short: "Manage orders",
	Long:  `Create, get, and list orders.`,
}

func init() {
	// Add the subcommands to the orders command
	ordersCmd.AddCommand(createCmd)
	ordersCmd.AddCommand(getOrderCmd)
	ordersCmd.AddCommand(listCmd)

	// Add the orders command to the root command
	RootCmd.AddCommand(ordersCmd)

	ordersCmd.SetHelpTemplate(`
$ birdcli ` + color.YellowString("order") + `

🛒 Create, retrieve, and list orders to manage your merchant transactions.

` + color.GreenString("COMMANDS:") + `
  birdcli ` + color.YellowString("order create") + `        📝 Create a new order
  birdcli ` + color.YellowString("order get <id>") + `      🔍 Get order by id
  birdcli ` + color.YellowString("order list") + `          📋 List orders

Use "birdcli [command] --help" for more information about a command.
`)
}
