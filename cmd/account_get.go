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
			fmt.Println("API key not found. Please set the API key using the following command:")
			fmt.Println("\tbirdcli account set-api-key <your-api-key>")
			return
		}

		fmt.Printf("API key: %s\n", apiKey)
	},
}

// go run main.go account get-api-key
