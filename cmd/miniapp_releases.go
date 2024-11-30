package cmd

import "github.com/spf13/cobra"

var miniappReleasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Manage miniapp releases",
}

func init() {
	miniappCmd.AddCommand(miniappReleasesCmd)
}
