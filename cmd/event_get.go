package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var getEventCmd = &cobra.Command{
	Use:   "get <eventID>",
	Short: "Get an event",
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		eventID := args[0] // Retrieve the appID argument

		ctx, apiClient := auth.MustGetAuth()

		event, _, err := apiClient.EventsAPI.
			GetEvent(ctx, eventID).
			Execute()
		if err != nil {
			log.Fatalf("Error getting event: %v", err)
		}

		prettyprint.JSON(event)
	},
}

// go run main.go events get <eventID>
