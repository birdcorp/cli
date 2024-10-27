package miniprogram

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Version            string   `json:"version"`
	BackgroundColor    string   `json:"background_color"`
	BuildCommand       string   `json:"build_command"`
	BuildDirectory     string   `json:"build_directory"`
	Description        string   `json:"description"`
	ForegroundColor    string   `json:"foreground_color"`
	IconImage          string   `json:"icon_image"`
	Name               string   `json:"name"`
	NavBackgroundColor string   `json:"nav_background_color"`
	NavTextColor       string   `json:"nav_text_color"`
	Tags               []string `json:"tags"`
	Testers            []string `json:"testers"`
}

var defaultConfig = Config{
	Name:               "my app name",
	Version:            "1.0.0",
	BackgroundColor:    "#ffffff",
	BuildCommand:       "npm run build",
	BuildDirectory:     "build",
	Description:        "Some description about the miniprogram",
	ForegroundColor:    "#000000",
	IconImage:          "https://payments-webapp-assets-stage.s3.us-west-2.amazonaws.com/miniprograms/blank-icon.png",
	NavBackgroundColor: "transparent",
	NavTextColor:       "dark",
	Tags:               []string{},
	Testers:            []string{},
}

// CreateConfig creates a miniprogram configuration file with default values.
func InitConfig() error {
	file, err := os.Create("miniprogram-config.json") // Create the JSON file
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after we're done

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Format the JSON output for better readability

	if err := encoder.Encode(defaultConfig); err != nil {
		return fmt.Errorf("error encoding config to JSON: %w", err)
	}

	fmt.Println("miniprogram-config.json created successfully!")
	return nil
}
