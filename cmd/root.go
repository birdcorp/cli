package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd
var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "root comand",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available Commands:")

		cmd.Root().Commands()

		for _, c := range cmd.Root().Commands() {
			fmt.Printf("  %s: %s\n", c.Use, c.Short)
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(-1)
	}
}
