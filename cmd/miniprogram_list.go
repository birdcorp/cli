package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listMiniprogramCmd = &cobra.Command{
	Use:   "list",
	Short: "List all miniprograms",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic to list miniprograms
		fmt.Println("Listing all miniprograms...")
		// Here you would typically call an API or database to retrieve the list of miniprograms
	},
}
