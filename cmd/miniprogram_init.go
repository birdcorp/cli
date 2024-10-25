package cmd

import (
	"log"

	"github.com/birdcorp/cli/pkg/miniprogram"
	"github.com/spf13/cobra"
)

var initMiniprogramCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a miniprogram",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, err := getAuth()
		if err != nil {
			log.Fatal(err)
		}

		miniprogram.InitConfig()

		log.Println("Initializing miniprogram.")
	},
}

// go run main.go miniprogram init
