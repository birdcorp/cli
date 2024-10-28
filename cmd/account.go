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
}

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

var deleteApiKeyCmd = &cobra.Command{
	Use:   "delete-api-key",
	Short: "Delete the API key",
	Long:  `This command deletes the API key from the local configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteAPIKey()
		if err != nil {
			fmt.Printf("Error deleting API key: %v\n", err)
		} else {
			fmt.Println("API key deleted successfully!")
		}
	},
}

var meAuthCmd = &cobra.Command{
	Use:   "me",
	Short: "Get the current user",
	Long:  `This command retrieves the current user from the API.`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}
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
	accountCmd.AddCommand(meAuthCmd)
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

func deleteAPIKey() error {
	// Check if the config file exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		return fmt.Errorf("config.json file does not exist")
	}

	// Read the existing configuration
	file, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}

	var config map[string]string
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	// Remove the API key if it exists
	delete(config, "API_KEY")

	// Write the updated configuration back to the file
	updatedFile, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", updatedFile, 0644)
}

// checkAPIKey verifies if the API key is set in the configuration file.
func checkAPIKey() bool {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return false
	}

	var config map[string]string
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return false
	}

	_, exists := config["API_KEY"]
	return exists
}

func printJSON(resp interface{}) {
	b, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Println("unable to decode response: ", err)
		return
	}

	fmt.Println(string(b))
}

// getAuthContext retrieves the API key and creates an authentication context for the API client
func getAuth() (context.Context, *birdsdk.APIClient, error) {
	apiKey, err := getAPIKey() // Assuming getAPIKey() is defined elsewhere in your code
	if err != nil {
		return nil, nil, fmt.Errorf("error getting API key: %w", err)
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

	return ctx, apiClient, nil
}
