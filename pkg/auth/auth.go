package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/spf13/afero"
)

var fs = afero.NewOsFs()

// getAuthContext retrieves the API key and creates an authentication context for the API client
func MustGetAuth() (context.Context, *birdsdk.APIClient) {
	apiKey, err := GetAPIKey()
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

func DeleteAPIKey() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "config.json")

	// Check if the config file exists
	exists, err := afero.Exists(fs, configFile)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("config.json file does not exist")
	}

	// Read the existing configuration
	file, err := afero.ReadFile(fs, configFile)
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

	return afero.WriteFile(fs, configFile, updatedFile, 0600)
}

func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".birdcli")
	if err := fs.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}
	return configDir, nil
}

func SaveAPIKey(apiKey string) error {
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
	return afero.WriteFile(fs, configFile, file, 0600)
}

func GetAPIKey() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}

	configFile := filepath.Join(configDir, "config.json")
	file, err := afero.ReadFile(fs, configFile)
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
