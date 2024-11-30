package miniapp

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DefaultVersion      = "1.0.0"
	DefaultBgColor      = "#FFFFFF"
	DefaultFgColor      = "#000000"
	DefaultNavBgColor   = "#F0F0F0"
	DefaultNavTextColor = "dark"
	DefaultLanguage     = "en"
	DefaultPrivacyURL   = "https://myapp.com/privacy"
	DefaultTermsURL     = "https://myapp.com/terms"
	ConfigFileName      = "miniapp.config.json"
)

type Config struct {
	AppInfo struct {
		AppID       string   `json:"appID"`
		Name        string   `json:"name"`
		Version     string   `json:"version"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	} `json:"appInfo"`
	Appearance struct {
		BackgroundColor    string `json:"backgroundColor"`
		ForegroundColor    string `json:"foregroundColor"`
		NavBackgroundColor string `json:"navBackgroundColor"`
		NavTextColor       string `json:"navTextColor"`
	} `json:"appearance"`
	Build struct {
		BuildDirectory string `json:"buildDirectory"`
	} `json:"build"`
	Assets struct {
		AppIcon string `json:"appIcon"`
	} `json:"assets"`
	Users struct {
		Testers []string `json:"testers"`
	} `json:"users"`
	Configuration struct {
		DefaultLanguage   string `json:"defaultLanguage"`
		PrivacyPolicyUrl  string `json:"privacyPolicyUrl"`
		TermsOfServiceUrl string `json:"termsOfServiceUrl"`
	} `json:"configuration"`
}

// GetConfig reads and parses the local miniprogram configuration file
func GetConfig() (*Config, error) {
	file, err := os.Open(ConfigFileName)
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
		AppInfo: struct {
			AppID       string   `json:"appID"`
			Name        string   `json:"name"`
			Version     string   `json:"version"`
			Description string   `json:"description"`
			Tags        []string `json:"tags"`
		}{
			AppID:       appID,
			Name:        name,
			Version:     DefaultVersion,
			Description: description,
			Tags:        []string{},
		},
		Appearance: struct {
			BackgroundColor    string `json:"backgroundColor"`
			ForegroundColor    string `json:"foregroundColor"`
			NavBackgroundColor string `json:"navBackgroundColor"`
			NavTextColor       string `json:"navTextColor"`
		}{
			BackgroundColor:    DefaultBgColor,
			ForegroundColor:    DefaultFgColor,
			NavBackgroundColor: DefaultNavBgColor,
			NavTextColor:       DefaultNavTextColor,
		},
		Build: struct {
			BuildDirectory string `json:"buildDirectory"`
		}{
			BuildDirectory: buildDir,
		},
		Assets: struct {
			AppIcon string `json:"appIcon"`
		}{
			AppIcon: appIcon,
		},
		Users: struct {
			Testers []string `json:"testers"`
		}{
			Testers: []string{},
		},
		Configuration: struct {
			DefaultLanguage   string `json:"defaultLanguage"`
			PrivacyPolicyUrl  string `json:"privacyPolicyUrl"`
			TermsOfServiceUrl string `json:"termsOfServiceUrl"`
		}{
			DefaultLanguage:   DefaultLanguage,
			PrivacyPolicyUrl:  DefaultPrivacyURL,
			TermsOfServiceUrl: DefaultTermsURL,
		},
	}

	file, err := os.Create(ConfigFileName) // Create the JSON file
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
