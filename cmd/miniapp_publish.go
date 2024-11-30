package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/fs"
	"github.com/birdcorp/cli/pkg/miniapp"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var publishMiniappCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a miniapp",
	Args:  nil,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, apiClient := auth.MustGetAuth()

		config, err := miniapp.GetConfig()
		if err != nil {
			fmt.Printf("%s\n", color.RedString("Missing miniapp config. Please run 'birdcli miniapp init' to create one."))
			os.Exit(1)
			return
		}

		// Zip the build folder to upload
		sourceFile := config.Build.BuildDirectory // The file you want to zip
		destinationZip := "./release.zip"         // The name of the output zip file

		err = fs.ZipDir(sourceFile, destinationZip)
		if err != nil {
			log.Fatal("Error zipping file:", err)
		}

		// Example usage: Upload a file
		file, err := os.Open(destinationZip)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Get app icon file
		iconFile, err := os.Open(config.Assets.AppIcon)
		if err != nil {
			fmt.Println("Error opening icon file:", err)
			return
		}
		defer iconFile.Close()

		release, resp, err := apiClient.MiniprogramAPI.
			CreateMiniprogramRelease(ctx, config.AppInfo.AppID).
			Build(file).
			AppIcon(iconFile).
			Name(config.AppInfo.Name).
			Version(config.AppInfo.Version).
			Description(config.AppInfo.Description).
			BackgroundColor(config.Appearance.BackgroundColor).
			ForegroundColor(config.Appearance.ForegroundColor).
			NavBackgroundColor(config.Appearance.NavBackgroundColor).
			NavTextColor(config.Appearance.NavTextColor).
			Execute()

		if err != nil {
			printer.HandleAPIFailure(resp)
			return
		}

		defer resp.Body.Close()

		fmt.Println("File uploaded successfully!")
		// Delete the temporary zip file
		err = os.Remove(destinationZip)
		if err != nil {
			fmt.Printf("Warning: Failed to delete temporary zip file %s: %v\n", destinationZip, err)
		}

		fmt.Println(release.Message)
	},
}

// go run main.go miniapp publish
