package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var listWebhooksCmd = &cobra.Command{
	Use:   "list",
	Short: "List webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Creating webhook.")

		webhooks, httpRes, err := apiClient.WebhooksAPI.
			ListWebhooks(ctx).
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
			log.Fatalf("Error creating webhook: %v", err)
		}

		prettyprint.JSON(webhooks)
	},
}

func init() {

}
