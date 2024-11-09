package miniprogram

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppInfo struct {
	AppID       string   `json:"appID"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type Appearance struct {
	BackgroundColor    string `json:"backgroundColor"`
	ForegroundColor    string `json:"foregroundColor"`
	NavBackgroundColor string `json:"navBackgroundColor"`
	NavTextColor       string `json:"navTextColor"`
}

type Build struct {
	BuildDirectory string `json:"buildDirectory"`
}

type Assets struct {
	AppIcon string `json:"appIcon"`
}

type Users struct {
	Testers []string `json:"testers"`
}

type Configuration struct {
	DefaultLanguage   string `json:"defaultLanguage"`
	PrivacyPolicyUrl  string `json:"privacyPolicyUrl"`
	TermsOfServiceUrl string `json:"termsOfServiceUrl"`
}

type Config struct {
	AppInfo       AppInfo       `json:"appInfo"`
	Appearance    Appearance    `json:"appearance"`
	Build         Build         `json:"build"`
	Assets        Assets        `json:"assets"`
	Users         Users         `json:"users"`
	Configuration Configuration `json:"configuration"`
}

// GetConfig reads and parses the local miniprogram configuration file
func GetConfig() (*Config, error) {
	file, err := os.Open("miniprogram-config.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %w", err)
	}

	return &config, nil
}

// CreateConfig creates a new Config struct with the provided values
func CreateConfig(appID string, name string, description string, buildDir string, appIcon string) (*Config, error) {
	// Create config from payload
	config := &Config{
		AppInfo: AppInfo{
			AppID:       appID,
			Name:        name,
			Version:     "1.0.0",
			Description: description,
			Tags:        []string{},
		},
		Appearance: Appearance{
			BackgroundColor:    "#FFFFFF",
			ForegroundColor:    "#000000",
			NavBackgroundColor: "#F0F0F0",
			NavTextColor:       "dark",
		},
		Build: Build{
			BuildDirectory: buildDir,
		},
		Assets: Assets{
			AppIcon: appIcon,
		},
		Users: Users{
			Testers: []string{},
		},
		Configuration: Configuration{
			DefaultLanguage:   "en",
			PrivacyPolicyUrl:  "https://myapp.com/privacy",
			TermsOfServiceUrl: "https://myapp.com/terms",
		},
	}

	file, err := os.Create("miniprogram-config.json") // Create the JSON file
	if err != nil {
		return nil, fmt.Errorf("error creating config file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after we're done

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Format the JSON output for better readability

	if err := encoder.Encode(config); err != nil {
		return nil, fmt.Errorf("error encoding config to JSON: %w", err)
	}

	return config, nil
}
