package cmd

import (
	"fmt"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/spf13/cobra"
)

var accountLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long:  `This command logs you out by deleting your API key from the local configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := auth.DeleteAPIKey()
		if err != nil {
			fmt.Printf("Error deleting API key: %v\n", err)
		} else {
			fmt.Println("API key deleted successfully!")
		}
	},
}
