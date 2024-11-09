package cmd

import "github.com/spf13/cobra"

var accountBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get account balance",
	Long:  `This command retrieves the account balance.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	accountCmd.AddCommand(accountBalanceCmd)
}
