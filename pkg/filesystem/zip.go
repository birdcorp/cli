package filesystem

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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
