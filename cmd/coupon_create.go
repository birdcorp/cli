package cmd

import (
	"fmt"
	"io"
	"log"
	"time"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/spf13/cobra"
)

var (
	couponCode        string
	couponType        string
	couponAmount      float32
	couponExpiryDate  time.Time
	couponDiscountAmt float32
)

var couponCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new coupon",
	Long:  `Create a new coupon with specified amount and expiry date.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, apiClient := mustGetAuth()

		log.Println("Creating new fixed amount coupon code...")

		coupon, resp, err := apiClient.CouponCodesAPI.
			CreateCouponCode(ctx).
			CreateCouponRequest(birdsdk.CreateCouponRequest{
				Code:           couponCode,
				Type:           birdsdk.CouponType(couponType),
				AmountIssued:   int32(couponAmount),
				ExpiryDate:     couponExpiryDate,
				DiscountAmount: &couponDiscountAmt,
			}).
			Execute()

		if err != nil {
			body, _ := io.ReadAll(resp.Body)
			log.Printf("Response body: %s\n", string(body))
			return fmt.Errorf("failed to create coupon: %v", err)
		}

		log.Println("Successfully created fixed amount coupon code. Response:")
		prettyprint.JSON(coupon)
		return nil
	},
}

func init() {
	couponCmd.AddCommand(couponCreateCmd)

	couponCreateCmd.Flags().StringVar(&couponCode, "code", "", "Coupon code")
	couponCreateCmd.Flags().StringVar(&couponType, "type", "fixed_amount", "Coupon type (fixed_amount)")
	couponCreateCmd.Flags().Float32Var(&couponAmount, "amount", 0, "Amount issued")
	couponCreateCmd.Flags().Float32Var(&couponDiscountAmt, "discount", 0, "Discount amount")

	couponCreateCmd.MarkFlagRequired("code")
	couponCreateCmd.MarkFlagRequired("amount")
	couponCreateCmd.MarkFlagRequired("discount")
}

/*

go run main.go coupon create --code FIXED10 --amount 10 --discount 10

*/
