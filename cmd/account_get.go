package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getApiKeyCmd = &cobra.Command{
	Use:   "get-api-key",
	Short: "Get the API key",
	Long:  `This command retrieves the API key from the local configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, err := getAPIKey()
		if err != nil {
			fmt.Printf("Error retrieving API key: %v\n", err)
		} else {
			fmt.Printf("API key: %s\n", apiKey)
		}
	},
}
