package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/auth"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
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
	Run: func(cmd *cobra.Command, args []string) {
		ctx, apiClient := auth.MustGetAuth()

		// Get code from flag or prompt
		code, err := cmd.Flags().GetString("code")
		if err != nil {
			log.Println("Error retrieving coupon code flag:", err)
			return
		}
		if code == "" {
			prompt := promptui.Prompt{
				Label: "Enter coupon code",
			}
			code, err = prompt.Run()
			if err != nil {
				log.Println("Error during coupon code prompt:", err)
				return
			}
		}

		// Get type from flag or prompt
		cType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Println("Error retrieving coupon type flag:", err)
			return
		}
		if cType == "" {
			prompt := promptui.Select{
				Label: "Select coupon type",
				Items: []string{"fixed_amount", "percentage"},
			}
			_, cType, err = prompt.Run()
			if err != nil {
				log.Println("Error during coupon type prompt:", err)
				return
			}
		}

		var discountPrompt string

		switch cType {
		case "fixed_amount":
			discountPrompt = "Enter discount amount"
		case "percentage":
			discountPrompt = "Enter discount percentage"
		}

		// Get discount from flag or prompt
		discount, err := cmd.Flags().GetFloat32("discount")
		if err != nil {
			log.Println("Error retrieving discount flag:", err)
			return
		}
		if discount == 0 {
			prompt := promptui.Prompt{
				Label: discountPrompt,
			}
			discountStr, err := prompt.Run()
			if err != nil {
				log.Println("Error during discount prompt:", err)
				return
			}
			fmt.Sscanf(discountStr, "%f", &discount)
		}

		// Get amount from flag or prompt
		amount, err := cmd.Flags().GetFloat32("amount")
		if err != nil {
			log.Println("Error retrieving amount flag:", err)
			return
		}
		if amount == 0 {
			prompt := promptui.Prompt{
				Label: "Enter amount issued",
			}
			amountStr, err := prompt.Run()
			if err != nil {
				log.Println("Error during amount prompt:", err)
				return
			}
			fmt.Sscanf(amountStr, "%f", &amount)
		}

		// Get expiry date from flag or prompt
		expiryStr, err := cmd.Flags().GetString("expiry")
		if err != nil {
			log.Println("Error retrieving expiry date flag:", err)
			return
		}
		var expiryDate time.Time
		if expiryStr == "" {
			prompt := promptui.Prompt{
				Label: "Enter expiry date (YYYY-MM-DD)",
			}
			expiryStr, err = prompt.Run()
			if err != nil {
				log.Println("Error during expiry date prompt:", err)
				return
			}
		}
		expiryDate, err = time.Parse("2006-01-02", expiryStr)
		if err != nil {
			log.Printf("Invalid date format. Please use YYYY-MM-DD: %v\n", err)
			return
		}

		var coupon *birdsdk.CouponCode
		var resp *http.Response

		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = " Creating coupon..."
		s.Start()

		switch cType {
		case "fixed_amount":
			coupon, resp, err = apiClient.CouponCodesAPI.
				CreateCouponCode(ctx).
				CreateCouponRequest(birdsdk.CreateCouponRequest{
					Code:           code,
					Type:           birdsdk.CouponType(cType),
					AmountIssued:   int32(amount),
					ExpiryDate:     expiryDate,
					DiscountAmount: &discount,
				}).
				Execute()
		case "percentage":
			coupon, resp, err = apiClient.CouponCodesAPI.
				CreateCouponCode(ctx).
				CreateCouponRequest(birdsdk.CreateCouponRequest{
					Code:            code,
					Type:            birdsdk.CouponType(cType),
					AmountIssued:    int32(amount),
					ExpiryDate:      expiryDate,
					DiscountPercent: &discount,
				}).
				Execute()
		}

		s.Stop()

		if err != nil {
			log.Println("Error creating coupon code:", err)
			if resp != nil && resp.Body != nil {
				body, _ := io.ReadAll(resp.Body)
				log.Printf("Response body: %s\n", string(body))
			}
			return
		}

		fmt.Printf("\n%s\n\n", color.GreenString("✅ Coupon Created"))

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		switch cType {
		case "fixed_amount":
			fmt.Fprintln(w, "Code\tDiscount Amount\tExpiry Date\tID\tRemaining\tType\tAmount Issued")
			fmt.Fprintf(w, "%s\t%.2f\t%s\t%s\t%.2f\t%s\t%.2f\n",
				*coupon.Code,
				*coupon.DiscountAmount,
				coupon.ExpiryDate.Format("2006-01-02"),
				*coupon.Id,
				float64(*coupon.Remaining),
				*coupon.Type.Ptr(),
				float64(*coupon.AmountIssued),
			)
		case "percentage":
			fmt.Fprintln(w, "Code\tDiscount Percent\tExpiry Date\tID\tRemaining\tType\tAmount Issued")
			fmt.Fprintf(w, "%s\t%.2f\t%s\t%s\t%.2f\t%s\t%.2f\n",
				*coupon.Code,
				*coupon.DiscountPercent,
				coupon.ExpiryDate.Format("2006-01-02"),
				*coupon.Id,
				float64(*coupon.Remaining),
				*coupon.Type.Ptr(),
				float64(*coupon.AmountIssued),
			)
		}

		w.Flush()
	},
}

func init() {
	couponCmd.AddCommand(couponCreateCmd)

	couponCreateCmd.Flags().String("code", "", "Coupon code")
	couponCreateCmd.Flags().String("type", "", "Coupon type")
	couponCreateCmd.Flags().Float32("amount", 0, "Amount issued")
	couponCreateCmd.Flags().Float32("discount", 0, "Discount amount")
	couponCreateCmd.Flags().String("expiry", "", "Expiry date (YYYY-MM-DD)")
}

/*

go run main.go coupon create --code FIXED10 --amount 10 --discount 10 --expiry 2024-12-31

*/
