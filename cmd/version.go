package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("birdcli version %s\n", version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

// go build -ldflags="-X 'github.com/birdcorp/cli/cmd.version=v1.2.3'" -o birdcli
