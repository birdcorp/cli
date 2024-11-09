package cmd

import (
	"fmt"
	"io"
	"log"
	"time"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/prettyprint"
	"github.com/manifoldco/promptui"
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

		// Get code from flag or prompt
		code, err := cmd.Flags().GetString("code")
		if err != nil {
			return fmt.Errorf("error getting code flag: %v", err)
		}
		if code == "" {
			prompt := promptui.Prompt{
				Label: "Enter coupon code",
			}
			code, err = prompt.Run()
			if err != nil {
				return fmt.Errorf("prompt failed: %v", err)
			}
		}

		// Get amount from flag or prompt
		amount, err := cmd.Flags().GetFloat32("amount")
		if err != nil {
			return fmt.Errorf("error getting amount flag: %v", err)
		}
		if amount == 0 {
			prompt := promptui.Prompt{
				Label: "Enter amount issued",
			}
			amountStr, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("prompt failed: %v", err)
			}
			fmt.Sscanf(amountStr, "%f", &amount)
		}

		// Get discount from flag or prompt
		discount, err := cmd.Flags().GetFloat32("discount")
		if err != nil {
			return fmt.Errorf("error getting discount flag: %v", err)
		}
		if discount == 0 {
			prompt := promptui.Prompt{
				Label: "Enter discount amount",
			}
			discountStr, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("prompt failed: %v", err)
			}
			fmt.Sscanf(discountStr, "%f", &discount)
		}

		log.Println("Creating new fixed amount coupon code...")

		coupon, resp, err := apiClient.CouponCodesAPI.
			CreateCouponCode(ctx).
			CreateCouponRequest(birdsdk.CreateCouponRequest{
				Code:           code,
				Type:           birdsdk.CouponType("fixed_amount"),
				AmountIssued:   int32(amount),
				ExpiryDate:     couponExpiryDate,
				DiscountAmount: &discount,
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

	couponCreateCmd.Flags().String("code", "", "Coupon code")
	couponCreateCmd.Flags().Float32("amount", 0, "Amount issued")
	couponCreateCmd.Flags().Float32("discount", 0, "Discount amount")
}

/*

go run main.go coupon create --code FIXED10 --amount 10 --discount 10

*/
