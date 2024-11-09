package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var accountLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long:  `This command logs you out by deleting your API key from the local configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteAPIKey()
		if err != nil {
			fmt.Printf("Error deleting API key: %v\n", err)
		} else {
			fmt.Println("API key deleted successfully!")
		}
	},
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
