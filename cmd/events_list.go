package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var listEventsCmd = &cobra.Command{
	Use:   "list",
	Short: "List events",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		events, _, err := apiClient.EventsAPI.
			ListEvents(ctx).
			Execute()
		if err != nil {
			log.Fatalf("Error listing events: %v", err)
		}

		prettyprint.JSON(events)
	},
}

// go run main.go events list
