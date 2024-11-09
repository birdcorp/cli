package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/birdcorp/cli/pkg/miniprogram"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var publishMiniprogramCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a miniprogram",
	Args:  nil, // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {

		ctx, apiClient := mustGetAuth()

		config, err := miniprogram.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		// Zip the build folder to upload
		sourceFile := config.Build.BuildDirectory // The file you want to zip
		destinationZip := "./release.zip"         // The name of the output zip file

		err = ZipDir(sourceFile, destinationZip)
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
			BackgroundColor("#FFFFFF").
			ForegroundColor("#000000").
			NavBackgroundColor("#F0F0F0").
			NavTextColor("dark").
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

// ZipDir zips the given directory into the destination zip file.
func ZipDir(sourceDir, destinationZip string) error {
	// Create the zip file
	zipFile, err := os.Create(destinationZip)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the source directory and add files to the zip archive
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path of the file/directory for the zip archive
		relPath := strings.TrimPrefix(path, filepath.Clean(sourceDir)+string(os.PathSeparator))

		// Skip the source directory itself
		if relPath == "" {
			return nil
		}

		// Create a zip header based on the file info
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("failed to create zip header: %w", err)
		}

		// Set the name of the file in the zip archive
		header.Name = relPath

		// If it's a directory, ensure the header name ends with a "/"
		if info.IsDir() {
			header.Name += "/"
		} else {
			// Set compression method for files (not directories)
			header.Method = zip.Deflate
		}

		// Create the writer for the zip file entry
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("failed to create zip writer: %w", err)
		}

		// If it's a directory, we don't need to write data
		if !info.IsDir() {
			// Open the file to be added to the zip
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			defer file.Close()

			// Copy the file data into the zip archive
			if _, err := io.Copy(writer, file); err != nil {
				return fmt.Errorf("failed to copy file data to zip: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error zipping directory: %w", err)
	}

	return nil
}
