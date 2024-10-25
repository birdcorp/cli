package cmd

import (
	"github.com/spf13/cobra"
)

// ordersCmd represents the orders command
var miniprogramCmd = &cobra.Command{
	Use:   "miniprogram",
	Short: "Manage miniprogram",
	Long:  `Create, get, and list miniprogram.`,
}

func init() {
	// Add the subcommands to the orders command
	miniprogramCmd.AddCommand(initMiniprogramCmd)
	miniprogramCmd.AddCommand(createMiniprogramCmd)
	miniprogramCmd.AddCommand(listMiniprogramCmd)
	miniprogramCmd.AddCommand(deleteMiniprogramCmd)

	// Add the orders command to the root command
	RootCmd.AddCommand(miniprogramCmd)
}

// go run main.go miniprogram list
