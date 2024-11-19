package cmd

import (
	"fmt"
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var accountLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Set the API key",
	Long:  `This command sets the API key and saves it to a local configuration file.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label: "API Key",
		}

		apiKey, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}

		err = auth.SaveAPIKey(apiKey)
		if err != nil {
			fmt.Printf("Error saving API key: %v\n", err)
		} else {
			fmt.Println("API key saved successfully!")
		}

	},
}
