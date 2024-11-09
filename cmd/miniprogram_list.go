package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var listMiniprogramCmd = &cobra.Command{
	Use:   "list",
	Short: "List all miniprograms",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := mustGetAuth()

		miniprograms, _, err := apiClient.MiniprogramAPI.
			ListMiniprograms(ctx).
			Execute()
		if err != nil {
			log.Fatalf("Error listing miniprograms: %v", err)
		}

		prettyprint.JSON(miniprograms)

		if len(miniprograms.Data) > 0 {
			var items []string
			for _, mp := range miniprograms.Data {
				if mp.ActiveRelease != nil {
					items = append(items, *mp.ActiveRelease.AppInfo.Name)
				}
			}

			prompt := promptui.Select{
				Label: "Select Miniprogram",
				Items: items,
			}

			_, result, err := prompt.Run()
			if err != nil {
				log.Fatalf("Prompt failed: %v", err)
			}

			log.Printf("You selected: %s", result)
		}

	},
}

// go run main.go miniprogram list
