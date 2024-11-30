package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	birdsdk "github.com/birdcorp/bird-go-sdk"

	"github.com/fatih/color"
)

func AccountInfo(account *birdsdk.Account) {
	fmt.Printf("\n%s\n\n", color.CyanString("📋  Account Information"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Basic Info Section
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Name:"), color.WhiteString(account.GetName()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Email:"), color.CyanString(account.GetEmail()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Phone:"), color.WhiteString(account.GetPhone()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("URL:"), color.BlueString(account.GetUrl()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Address:"), color.WhiteString(account.GetAddress()))

	// Branding Section
	fmt.Fprintf(w, "\n%s\n", color.YellowString("🎨  Branding"))
	fmt.Fprintf(w, "%s\t%s %s\n", color.HiBlackString("Brand Color:"), color.HiWhiteString(account.GetBrandColor()), color.New(color.BgHiWhite).Add(color.FgBlack).Sprintf("  "))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Logo:"), color.BlueString(account.GetLogo()))

	// Business Address Section
	businessAddr := account.GetBusinessAddress()
	fmt.Fprintf(w, "\n%s\n", color.YellowString("🏢  Business Address"))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Street:"), color.WhiteString(businessAddr.GetLine1()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("City:"), color.WhiteString(businessAddr.GetCity()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("State:"), color.WhiteString(businessAddr.GetState()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Postal Code:"), color.WhiteString(businessAddr.GetPostalCode()))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Country:"), color.WhiteString(businessAddr.GetCountry()))

	w.Flush()
	fmt.Println()
}
