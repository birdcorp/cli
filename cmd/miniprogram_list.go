package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var listMiniprogramCmd = &cobra.Command{
	Use:   "list",
	Short: "List all miniprograms",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		miniprograms, _, err := apiClient.MiniprogramAPI.
			ListMiniprograms(ctx).
			Execute()
		if err != nil {
			log.Fatalf("Error listing miniprograms: %v", err)
		}

		printJSON(miniprograms)

	},
}

// go run main.go miniprogram list
