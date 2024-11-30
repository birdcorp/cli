package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/birdcorp/cli/pkg/formatting"
	"github.com/fatih/color"
)

func WebhookList(webhooks []birdsdk.Webhook) {
	fmt.Printf("\n%s\n\n", color.CyanString("📋 Webhooks"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Created\tID\tURL\tSecret")

	for _, webhook := range webhooks {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			formatting.FormatRelativeTime(webhook.GetCreatedAt()),
			webhook.Id,
			webhook.Url,
			webhook.Secret)
	}

	w.Flush()
	fmt.Println()
}

func Webhook(webhook *birdsdk.Webhook) {
	fmt.Printf("\n%s\n\n", color.CyanString("🔔 Webhook Details"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("ID:"), color.GreenString(webhook.Id))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("URL:"), color.CyanString(webhook.Url))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Secret:"), color.YellowString(webhook.Secret))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Created:"), color.BlueString(formatting.FormatRelativeTime(webhook.GetCreatedAt())))
	fmt.Fprintf(w, "%s\t%s\n", color.HiBlackString("Updated:"), color.BlueString(formatting.FormatRelativeTime(webhook.GetUpdatedAt())))
	w.Flush()
	fmt.Println()
}
