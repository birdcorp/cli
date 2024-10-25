package miniprogram

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppID              string   `json:"app_id"`
	AppSecret          string   `json:"app_secret"`
	BackgroundColor    string   `json:"background_color"`
	BuildCommand       string   `json:"build_command"`
	BuildDirectory     string   `json:"build_directory"`
	CodeImage          string   `json:"code_image"`
	Description        string   `json:"description"`
	ForegroundColor    string   `json:"foreground_color"`
	IconImage          string   `json:"icon_image"`
	Name               string   `json:"name"`
	NavBackgroundColor string   `json:"nav_background_color"`
	NavTextColor       string   `json:"nav_text_color"`
	Tags               []string `json:"tags"`
	Testers            []string `json:"testers"`
	URL                string   `json:"url"`
}

var defaultConfig = Config{
	AppID:              "",
	AppSecret:          "",
	BackgroundColor:    "#ffffff",
	BuildCommand:       "npm run build",
	BuildDirectory:     "build",
	CodeImage:          "",
	Description:        "Some description about the miniprogram",
	ForegroundColor:    "#000000",
	IconImage:          "https://payments-webapp-assets-stage.s3.us-west-2.amazonaws.com/miniprograms/blank-icon.png",
	Name:               "",
	NavBackgroundColor: "transparent",
	NavTextColor:       "dark",
	Tags:               []string{},
	Testers:            []string{},
	URL:                "https://www.my-staging-link-example.com",
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
