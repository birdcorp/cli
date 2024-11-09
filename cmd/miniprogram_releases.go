package cmd

import "github.com/spf13/cobra"

var miniprogramReleasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Manage miniprogram releases",
}

func init() {
	miniprogramCmd.AddCommand(miniprogramReleasesCmd)
}
