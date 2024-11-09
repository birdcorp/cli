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

	miniprogramCmd.AddCommand(createMiniprogramCmd)
	miniprogramCmd.AddCommand(listMiniprogramCmd)
	miniprogramCmd.AddCommand(deleteMiniprogramCmd)
	miniprogramCmd.AddCommand(publishMiniprogramCmd)
	miniprogramCmd.AddCommand(getMiniprogramCmd)
	miniprogramCmd.AddCommand(miniprogramPreviewCmd)
	miniprogramCmd.AddCommand(miniprogramReleasesListCmd)

	// Add the orders command to the root command
	RootCmd.AddCommand(miniprogramCmd)
}

// go run main.go miniprogram list
// go run main.go miniprogram upload <appID>
// go run main.go miniprogram releases list
