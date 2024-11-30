package cmd

import (
	"log"
	"time"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/printer"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage account",
	Long:  `Set, get, or delete the API key.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Fetching account information..."
		s.Start()

		account, _, err := apiClient.AccountAPI.
			GetAccount(ctx).
			Execute()
		if err != nil {
			s.Stop()
			log.Fatal(err)
		}
		s.Stop()

		printer.AccountInfo(account)

	},
}

func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// go run main.go auth me

func init() {

	// Add the auth command to the root command
	RootCmd.AddCommand(accountCmd)

	RootCmd.AddCommand(accountLoginCmd)
	RootCmd.AddCommand(accountLogoutCmd)
}
