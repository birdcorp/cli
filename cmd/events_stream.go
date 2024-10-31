package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var streamEventsCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream events",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, err := getAPIKey() // Assuming getAPIKey() is defined elsewhere in your code
		if err != nil {
			log.Fatalf("Error getting API key: %v", err)
		}
		log.Println("Streaming events.")

		// Set the request URL and headers
		req, err := http.NewRequest("GET", "https://api.birdwallet.xyz/v2/events/stream", nil)
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-API-KEY", apiKey)

		// Execute the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		// Check if the response status is 200 OK
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Failed to connect, status code: %d", resp.StatusCode)
		}

		// Read and parse each event line by line
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading response: %v", err)
			}

			// Check if the line contains "data: "
			if strings.HasPrefix(line, "data: ") {
				// Remove "data: " prefix
				eventJSON := strings.TrimPrefix(line, "data: ")
				// Define the struct to hold the parsed event data
				var event birdsdk.WebhookEvent

				// Unmarshal the event JSON
				if err := json.Unmarshal([]byte(eventJSON), &event); err != nil {
					log.Printf("Error unmarshaling event JSON: %v", err)
					continue
				}

				color.Set(color.FgGreen) // Set text color to green with a black background
				fmt.Println(event.Type)  // Print the event type
				color.Unset()            // Reset to default color

				fmt.Println("Raw event:")
				prettyprint.JSON(event)
				fmt.Println("") // Print a blank newline
				fmt.Println("") // Print a blank newline
			}
		}

	},
}

// go run main.go events stream
