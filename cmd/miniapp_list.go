package cmd

import (
	"fmt"
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var listMiniappCmd = &cobra.Command{
	Use:   "list",
	Short: "List all miniapps",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		miniapps, _, err := apiClient.MiniprogramAPI.
			ListMiniprograms(ctx).
			Execute()
		if err != nil {
			fmt.Printf("%s Failed to list miniapp:\n", color.RedString("❌"))
			return
		}

		if len(miniapps.Data) > 0 {
			var items []string
			for _, mp := range miniapps.Data {
				if mp.ActiveRelease != nil {
					items = append(items, *mp.ActiveRelease.AppInfo.Name)
				}
			}

			prompt := promptui.Select{
				Label: "Select Miniapp",
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

// go run main.go miniapp list
