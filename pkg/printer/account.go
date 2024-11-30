package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/ptr"
	"github.com/fatih/color"
)

func AccountInfo(account *birdsdk.Account) {
	fmt.Printf("\n%s\n\n", color.CyanString("📋  Account Information"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Basic Info Section
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Name:"), color.WhiteString(ptr.GetStringValue(account.Name)))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Email:"), color.CyanString(ptr.GetStringValue(account.Email)))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Phone:"), color.WhiteString(ptr.GetStringValue(account.Phone)))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("URL:"), color.BlueString(ptr.GetStringValue(account.Url)))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Address:"), color.WhiteString(ptr.GetStringValue(account.Address)))

	// Branding Section
	fmt.Fprintf(w, "\n%s\n", color.YellowString("🎨  Branding"))
	fmt.Fprintf(w, "%s\t%s %s\n", color.HiBlackString("Brand Color:"), color.HiWhiteString(ptr.GetStringValue(account.BrandColor)), color.New(color.BgHiWhite).Add(color.FgBlack).Sprintf("  "))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Logo:"), color.BlueString(ptr.GetStringValue(account.Logo)))

	// Business Address Section
	fmt.Fprintf(w, "\n%s\n", color.YellowString("🏢  Business Address"))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Street:"), color.WhiteString(account.BusinessAddress.Line1))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("City:"), color.WhiteString(account.BusinessAddress.City))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("State:"), color.WhiteString(account.BusinessAddress.State))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Postal Code:"), color.WhiteString(account.BusinessAddress.PostalCode))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Country:"), color.WhiteString(account.BusinessAddress.Country))

	w.Flush()
	fmt.Println()
}
