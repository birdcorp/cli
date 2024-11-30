package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ordersCmd represents the orders command
var eventsCmd = &cobra.Command{
	Use:   "event",
	Short: "Manage events",
	Long:  `Stream events.`,
}

func init() {
	// Add the subcommands to the orders command
	eventsCmd.AddCommand(streamEventsCmd)
	eventsCmd.AddCommand(listEventsCmd)
	// Add the orders command to the root command
	RootCmd.AddCommand(eventsCmd)

	eventsCmd.SetHelpTemplate(`
$ birdcli ` + color.YellowString("events") + `

📡 Stream, retrieve, and list events from the Bird API to monitor order status changes, webhook deliveries, and system events.

` + color.GreenString("COMMANDS:") + `
  birdcli ` + color.YellowString("event stream") + `        📺 Stream events
  birdcli ` + color.YellowString("event get <id>") + `      🔍 Get event by id
  birdcli ` + color.YellowString("event list") + `          📋 List events

Use "birdcli [command] --help" for more information about a command.
`)
}
