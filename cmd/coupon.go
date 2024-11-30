package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var couponCmd = &cobra.Command{
	Use:   "coupon",
	Short: "Manage coupons",
	Long:  `Create, list, and delete coupons.`,
}

func init() {
	RootCmd.AddCommand(couponCmd)

	couponCmd.SetHelpTemplate(`
$ birdcli ` + color.YellowString("coupon") + `

🎟️  Create, list, and manage discount coupons for your Bird store.

` + color.GreenString("COMMANDS:") + `
  birdcli ` + color.YellowString("coupon create") + `       ✨ Create a new coupon
  birdcli ` + color.YellowString("coupon list") + `         📋 List all coupons
  birdcli ` + color.YellowString("coupon get <id>") + `     🔍 Get coupon details
  birdcli ` + color.YellowString("coupon delete <id>") + `  🗑️  Delete a coupon

Use "birdcli [command] --help" for more information about a command.
`)
}
