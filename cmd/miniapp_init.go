package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/birdcorp/cli/pkg/auth"
	miniapp "github.com/birdcorp/cli/pkg/miniapp"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	defaultBuildDir     = "./build"
	defaultIconURL      = "https://dlkosrb2bmrzf.cloudfront.net/miniapps/blank-icon.png"
	defaultIconPath     = "./app-icon.png"
	defaultURL          = "https://example.com"
	defaultBgColor      = "#FFFFFF"
	defaultFgColor      = "#000000"
	defaultNavBgColor   = "#F0F0F0"
	defaultNavTextColor = "dark"
	defaultTags         = "tag1,tag2"
)

func promptUser(message string) string {
	prompt := promptui.Prompt{
		Label: color.CyanString(message),
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("%s Prompt failed %v", color.RedString("✗"), err)
	}
	return result
}

func promptBuildDir() string {
	buildDirOptions := []string{"./build", "./dist", "Custom"}
	prompt := promptui.Select{
		Label: color.CyanString("Select the directory containing your built miniapp files (e.g. 'build')"),
		Items: buildDirOptions,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("%s Prompt failed %v", color.RedString("✗"), err)
	}

	if result == "Custom" {
		return promptUser("Enter custom build directory path")
	}
	return result
}

var createMiniappCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new miniapp project",
	Long:  `Initialize a new miniapp project with configuration and required files.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		config, _ := miniapp.GetConfig()

		if config != nil {
			prompt := promptui.Prompt{
				Label:     color.YellowString("Miniapp config already exists. Do you want to overwrite it?"),
				IsConfirm: true,
			}

			result, err := prompt.Run()
			if err != nil {
				log.Fatalf("%s Miniapp initialization cancelled", color.RedString("✗"))
			}

			if result != "y" {
				fmt.Println(color.YellowString("Operation cancelled."))
				return
			}
		}

		// Get or prompt for required fields
		name := getRequiredFlag(cmd, "name", "Enter a name for your miniapp (e.g. 'My Store App')")
		description := getRequiredFlag(cmd, "description", "Enter a description for your miniapp")

		// Get build directory from flag or prompt
		buildDir := cmd.Flag("build-directory").Value.String()
		if buildDir == "" {
			buildDir = promptBuildDir()
		}

		// Check if icon file exists, download default if it doesn't
		if _, err := os.Stat(defaultIconPath); os.IsNotExist(err) {
			fmt.Printf("%s Downloading default icon...\n", color.BlueString("ℹ"))
			if err := downloadDefaultIcon(); err != nil {
				log.Fatalf("%s Error downloading default icon: %v", color.RedString("✗"), err)
			}
		}

		// Create miniapp
		response, httpRes, err := apiClient.MiniprogramAPI.
			CreateMiniprogram(ctx).
			Execute()

		if err != nil {
			handleAPIError(err, httpRes)
			os.Exit(1)
		}

		// Create local config file
		appID := response.GetId()
		if _, err := miniapp.CreateConfig(appID, name, description, buildDir, defaultIconPath); err != nil {
			log.Fatalf("%s Error creating config file: %v", color.RedString("✗"), err)
		}

		fmt.Printf("%s Miniapp initialized!\n", color.GreenString("✓"))
	},
}

func getRequiredFlag(cmd *cobra.Command, flagName, prompt string) string {
	value, err := cmd.Flags().GetString(flagName)
	if err != nil {
		log.Fatalf("%s Error getting %s flag: %v", color.RedString("✗"), flagName, err)
	}
	if value == "" {
		value = promptUser(prompt)
	}
	return value
}

func getOptionalFlag(cmd *cobra.Command, flagName, defaultValue string) string {
	value, err := cmd.Flags().GetString(flagName)
	if err != nil {
		log.Fatalf("%s Error getting %s flag: %v", color.RedString("✗"), flagName, err)
	}
	if value == "" {
		return defaultValue
	}
	return value
}

func downloadDefaultIcon() error {
	resp, err := http.Get(defaultIconURL)
	if err != nil {
		return fmt.Errorf("failed to download icon: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(defaultIconPath)
	if err != nil {
		return fmt.Errorf("failed to create icon file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to save icon file: %w", err)
	}
	return nil
}

func init() {
	createMiniappCmd.Flags().String("name", "", "Name of the miniapp")
	createMiniappCmd.Flags().String("description", "", "Description of the miniapp")
	createMiniappCmd.Flags().String("url", defaultURL, "URL of the miniapp")
	createMiniappCmd.Flags().String("background-color", defaultBgColor, "Background color code")
	createMiniappCmd.Flags().String("foreground-color", defaultFgColor, "Foreground color code")
	createMiniappCmd.Flags().String("nav-background-color", defaultNavBgColor, "Navigation background color")
	createMiniappCmd.Flags().String("nav-text-color", defaultNavTextColor, "Navigation text color (dark/light)")
	createMiniappCmd.Flags().String("tags", defaultTags, "Comma-separated list of tags")
	createMiniappCmd.Flags().String("build-directory", "", "Build directory path")
	createMiniappCmd.Flags().String("icon-image", "", "Icon image path")
}

/*
Example usage:
go run main.go miniapp init \
  --name "My App" \
  --description "My awesome miniapp" \
  --build-directory "./build" \
  --icon-image "./app-icon.png"
*/
