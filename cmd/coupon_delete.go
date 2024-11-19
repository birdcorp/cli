package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/birdcorp/cli/pkg/auth"
	"github.com/spf13/cobra"
)

var couponDeleteCmd = &cobra.Command{
	Use:   "delete [coupon-id]",
	Short: "Delete a coupon",
	Long:  `Delete a coupon by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		resp, err := apiClient.CouponCodesAPI.
			DeleteCouponCode(ctx, args[0]).
			Execute()

		if err != nil {
			log.Println("Error deleting coupon:", err)
			if resp != nil && resp.Body != nil {
				body, _ := io.ReadAll(resp.Body)
				log.Printf("Response body: %s\n", string(body))
			}
			return
		}

		fmt.Printf("Deleted coupon: %s\n", args[0])
	},
}

func init() {
	couponCmd.AddCommand(couponDeleteCmd)
}

/*
go run main.go coupon delete akpyDbB88kz2sUQBwaunD2
*/
