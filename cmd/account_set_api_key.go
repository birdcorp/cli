package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var apiKeyCmd = &cobra.Command{
	Use:   "set-api-key [API_KEY]",
	Short: "Set the API key",
	Long:  `This command sets the API key and saves it to a local configuration file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := args[0]
		err := saveAPIKey(apiKey)
		if err != nil {
			fmt.Printf("Error saving API key: %v\n", err)
		} else {
			fmt.Println("API key saved successfully!")
		}
	},
}
