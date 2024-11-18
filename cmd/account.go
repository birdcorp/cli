package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	// Add the auth command to the root command
	RootCmd.AddCommand(accountCmd)

	RootCmd.AddCommand(accountLoginCmd)
	RootCmd.AddCommand(accountLogoutCmd)
}

func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".birdcli")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}
	return configDir, nil
}

func saveAPIKey(apiKey string) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	config := map[string]string{
		"API_KEY": apiKey,
	}

	file, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "config.json")
	return os.WriteFile(configFile, file, 0600)
}

func getAPIKey() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}

	configFile := filepath.Join(configDir, "config.json")
	file, err := os.ReadFile(configFile)
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
	apiKey, err := getAPIKey()
	if err != nil {
		fmt.Printf("\x1b[31m❌ API key not found\x1b[0m\n")
		fmt.Printf("\x1b[33m💡 Please set the API key using the following command:\x1b[0m\n")
		fmt.Printf("\x1b[36m\t$ birdcli login\x1b[0m\n")
		os.Exit(1)
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
