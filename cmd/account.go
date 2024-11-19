package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/birdcorp/cli/pkg/ptr"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage account",
	Long:  `Set, get, or delete the API key.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		account, _, err := apiClient.AccountAPI.
			GetAccount(ctx).
			Execute()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n%s\n\n", color.CyanString("👨‍💼 Account Information"))
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "Name:\t%s\n", ptr.GetStringValue(account.Name))
		fmt.Fprintf(w, "Email:\t%s\n", ptr.GetStringValue(account.Email))
		fmt.Fprintf(w, "Phone:\t%s\n", ptr.GetStringValue(account.Phone))
		fmt.Fprintf(w, "URL:\t%s\n", ptr.GetStringValue(account.Url))
		fmt.Fprintf(w, "Address:\t%s\n", ptr.GetStringValue(account.Address))
		fmt.Fprintf(w, "Brand Color:\t%s\n", ptr.GetStringValue(account.BrandColor))
		fmt.Fprintf(w, "Logo:\t%s\n", ptr.GetStringValue(account.Logo))
		fmt.Fprintf(w, "\n🏢 Business Address:\n")
		fmt.Fprintf(w, "  Street:\t%s\n", account.BusinessAddress.Line1)
		fmt.Fprintf(w, "  City:\t%s\n", account.BusinessAddress.City)
		fmt.Fprintf(w, "  State:\t%s\n", account.BusinessAddress.State)
		fmt.Fprintf(w, "  Postal Code:\t%s\n", account.BusinessAddress.PostalCode)
		fmt.Fprintf(w, "  Country:\t%s\n", account.BusinessAddress.Country)
		w.Flush()
		fmt.Println()
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
