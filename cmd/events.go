package cmd

import "github.com/spf13/cobra"

// ordersCmd represents the orders command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Manage events",
	Long:  `Stream events.`,
}

func init() {
	// Add the subcommands to the orders command
	eventsCmd.AddCommand(streamEventsCmd)
	eventsCmd.AddCommand(getEventCmd)
	eventsCmd.AddCommand(listEventsCmd)
	// Add the orders command to the root command
	RootCmd.AddCommand(eventsCmd)

	eventsCmd.SetHelpTemplate(`
$ birdcli events

📡 The events command allows you to manage and monitor events from the Bird API.
You can stream events in real-time, retrieve specific events by ID, or list
historical events. This is useful for monitoring order status changes,
webhook deliveries, and other system events.

COMMANDS:
  birdcli events stream        📺 Stream events
  birdcli events get <id>      🔍 Get event by id
  birdcli events list          📋 List events

Use "birdcli [command] --help" for more information about a command.
`)
}
