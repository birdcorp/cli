package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/filesystem"
	"github.com/birdcorp/cli/pkg/miniprogram"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var publishMiniprogramCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a miniprogram",
	Args:  nil, // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {

		ctx, apiClient := auth.MustGetAuth()

		config, err := miniprogram.GetConfig()
		if err != nil {
			log.Fatalf("Error getting miniprogram config: %v", err)
			return
		}

		// Zip the build folder to upload
		sourceFile := config.Build.BuildDirectory // The file you want to zip
		destinationZip := "./release.zip"         // The name of the output zip file

		err = filesystem.ZipDir(sourceFile, destinationZip)
		if err != nil {
			log.Fatal("Error zipping file:", err)
		} else {
			fmt.Println("File successfully zipped!")
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
			fmt.Println("Error uploading file:", err)
			return
		}

		defer resp.Body.Close()

		fmt.Println("File uploaded successfully!")
		// Delete the temporary zip file
		err = os.Remove(destinationZip)
		if err != nil {
			fmt.Printf("Warning: Failed to delete temporary zip file %s: %v\n", destinationZip, err)
		}

		prettyprint.JSON(release)

	},
}

// go run main.go miniprogram publish
