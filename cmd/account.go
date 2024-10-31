package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage account",
	Long:  `Set, get, or delete the API key.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		account, _, err := apiClient.AccountAPI.
			GetAccount(ctx).
			Execute()
		if err != nil {
			log.Fatal(err)
		}

		prettyprint.JSON(account)
	},
}

// go run main.go auth me

func init() {
	accountCmd.AddCommand(apiKeyCmd)
	accountCmd.AddCommand(getApiKeyCmd)
	accountCmd.AddCommand(deleteApiKeyCmd) // Add the delete-api-key command

	// Add the auth command to the root command
	RootCmd.AddCommand(accountCmd)
}

func saveAPIKey(apiKey string) error {
	config := map[string]string{
		"API_KEY": apiKey,
	}

	file, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", file, 0644)
}

func getAPIKey() (string, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return "", err
	}

	var config map[string]string
	err = json.Unmarshal(file, &config)
	if err != nil {
		return "", err
	}

	apiKey, exists := config["API_KEY"]
	if !exists {
		return "", fmt.Errorf("API key not found in config")
	}

	return apiKey, nil
}

// getAuthContext retrieves the API key and creates an authentication context for the API client
func mustGetAuth() (context.Context, *birdsdk.APIClient) {
	apiKey, err := getAPIKey() // Assuming getAPIKey() is defined elsewhere in your code
	if err != nil {
		fmt.Println("API key not found. Please set the API key using the following command:")
		fmt.Println("\tbirdcli account set-api-key <your-api-key>")
		return nil, nil
	}

	config := birdsdk.NewConfiguration()
	apiClient := birdsdk.NewAPIClient(config)

	ctx := context.WithValue(
		context.Background(),
		birdsdk.ContextAPIKeys,
		map[string]birdsdk.APIKey{
			"ApiKeyAuth": {
				Key: apiKey,
			},
		},
	)

	return ctx, apiClient
}
