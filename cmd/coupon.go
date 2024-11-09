package cmd

import (
	"github.com/spf13/cobra"
)

var couponCmd = &cobra.Command{
	Use:   "coupon",
	Short: "Manage coupons",
	Long:  `Create, list, and delete coupons.`,
}

func init() {
	RootCmd.AddCommand(couponCmd)
}
