package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/birdcorp/cli/pkg/miniprogram"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func promptUser(message string) string {
	prompt := promptui.Prompt{
		Label: message,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v", err)
	}
	return result
}

var createMiniprogramCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a miniprogram",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		// Prompt for `name`
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal("Error getting name flag:", err)
		}
		if name == "" {
			name = promptUser("Enter miniprogram name")
		}

		// Prompt for `description`
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			log.Fatal("Error getting description flag:", err)
		}
		if description == "" {
			description = promptUser("Enter miniprogram description")
		}

		buildDir, err := cmd.Flags().GetString("build-directory")
		if err != nil {
			log.Fatal("Error getting build-directory flag:", err)
		}
		if buildDir == "" {
			buildDir = promptUser("Enter build directory path example: ./build")
		}

		appIcon, err := cmd.Flags().GetString("icon-image")
		if err != nil {
			log.Fatal("Error getting icon-image flag:", err)
		}
		if appIcon == "" {
			appIcon = promptUser("Enter icon file path example: ./app-icon.png")
		}

		response, httpRes, err := apiClient.MiniprogramAPI.
			CreateMiniprogram(ctx).
			Execute()
		if err != nil {
			if httpRes != nil {
				if httpRes.Body != nil {
					body, err := io.ReadAll(httpRes.Body)
					if err == nil {
						fmt.Println(string(body)) // Print the HTTP response body for error details
					} else {
						fmt.Println("Error reading response body:", err)
					}
				}
			}
			log.Fatalf("Error creating miniprogram: %v", err)
		}

		prettyprint.JSON(response)

		appID := response.GetId()

		// Create local config file
		_, err = miniprogram.CreateConfig(appID, name, description, buildDir, appIcon)
		if err != nil {
			log.Fatalf("Error creating config file: %v", err)
		}

		fmt.Println("config-miniprogram.json created successfully!")
	},
}

func init() {
	createMiniprogramCmd.Flags().String("name", "", "Name of the miniprogram")
	createMiniprogramCmd.Flags().String("description", "", "Description of the miniprogram")
	createMiniprogramCmd.Flags().String("url", "https://example.com", "URL of the miniprogram")
	createMiniprogramCmd.Flags().String("background-color", "#FFFFFF", "Background color code")
	createMiniprogramCmd.Flags().String("foreground-color", "#000000", "Foreground color code")
	createMiniprogramCmd.Flags().String("nav-background-color", "#F0F0F0", "Navigation background color")
	createMiniprogramCmd.Flags().String("nav-text-color", "dark", "Navigation text color (dark/light)")
	createMiniprogramCmd.Flags().String("tags", "tag1,tag2", "Comma-separated list of tags")
	createMiniprogramCmd.Flags().String("build-directory", "", "Build directory path")
	createMiniprogramCmd.Flags().String("icon-image", "", "Icon image path")
}

/*
go run main.go miniprogram create \
  --name "My App" \
  --url "https://myapp.com" \
  --description "My awesome miniprogram" \
  --background-color "#FFFFFF" \
  --foreground-color "#000000" \
  --icon-image "https://example.com/icon.png" \
  --nav-background-color "#F0F0F0" \
  --nav-text-color "dark" \
  --tags "tag1,tag2"
*/

// go run main.go miniprogram create

// ./fixtures/app/build

// ./fixtures/app-icon.png
