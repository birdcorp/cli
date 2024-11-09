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
}
