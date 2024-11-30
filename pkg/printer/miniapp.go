package printer

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	birdsdk "github.com/birdcorp/bird-go-sdk"
	"github.com/fatih/color"
)

func Miniapp(miniapp *birdsdk.MiniProgram) {
	fmt.Printf("\n%s\n\n", color.CyanString("📱 Miniapp Details"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Basic Info
	fmt.Fprintln(w, "ID\tStatus\tCreated\tUpdated")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
		*miniapp.Id,
		*miniapp.Status,
		miniapp.CreatedAt.Format("2006-01-02"),
		miniapp.UpdatedAt.Format("2006-01-02"))
	w.Flush()
	fmt.Println()

	// Active Release Details
	if miniapp.ActiveRelease != nil {
		fmt.Printf("%s\n\n", color.CyanString("📦 Active Release"))
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		fmt.Fprintln(w, "UUID\tStatus\tVersion\tCreated")
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			*miniapp.ActiveRelease.Uuid,
			*miniapp.ActiveRelease.Status,
			*miniapp.ActiveRelease.AppInfo.Version,
			miniapp.ActiveRelease.CreatedAt.Format("2006-01-02"))
		w.Flush()
		fmt.Println()

		// App Info
		if miniapp.ActiveRelease.AppInfo != nil {
			fmt.Printf("%s\n\n", color.CyanString("ℹ️  App Info"))
			w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			fmt.Fprintf(w, "Name:\t%s\n", *miniapp.ActiveRelease.AppInfo.Name)
			fmt.Fprintf(w, "App ID:\t%s\n", *miniapp.ActiveRelease.AppInfo.AppID)
			fmt.Fprintf(w, "Description:\t%s\n", *miniapp.ActiveRelease.AppInfo.Description)
			if len(miniapp.ActiveRelease.AppInfo.Tags) > 0 {
				fmt.Fprintf(w, "Tags:\t%v\n", miniapp.ActiveRelease.AppInfo.Tags)
			}
			w.Flush()
			fmt.Println()
		}

		// Appearance
		if miniapp.ActiveRelease.Appearance != nil {
			fmt.Printf("%s\n\n", color.CyanString("🎨 Appearance"))
			w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			if miniapp.ActiveRelease.Appearance.BackgroundColor != nil {
				fmt.Fprintf(w, "Background Color:\t%s\n", *miniapp.ActiveRelease.Appearance.BackgroundColor)
			}
			if miniapp.ActiveRelease.Appearance.ForegroundColor != nil {
				fmt.Fprintf(w, "Foreground Color:\t%s\n", *miniapp.ActiveRelease.Appearance.ForegroundColor)
			}
			if miniapp.ActiveRelease.Appearance.NavBackgroundColor != nil {
				fmt.Fprintf(w, "Nav Background Color:\t%s\n", *miniapp.ActiveRelease.Appearance.NavBackgroundColor)
			}
			if miniapp.ActiveRelease.Appearance.NavTextColor != nil {
				fmt.Fprintf(w, "Nav Text Color:\t%s\n", *miniapp.ActiveRelease.Appearance.NavTextColor)
			}
			w.Flush()
			fmt.Println()
		}

		// Configuration
		if miniapp.ActiveRelease.Configuration != nil {
			fmt.Printf("%s\n\n", color.CyanString("⚙️  Configuration"))
			w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			if miniapp.ActiveRelease.Configuration.DefaultLanguage != nil {
				fmt.Fprintf(w, "Default Language:\t%s\n", *miniapp.ActiveRelease.Configuration.DefaultLanguage)
			}
			if miniapp.ActiveRelease.Configuration.PrivacyPolicyUrl != nil {
				fmt.Fprintf(w, "Privacy Policy URL:\t%s\n", *miniapp.ActiveRelease.Configuration.PrivacyPolicyUrl)
			}
			if miniapp.ActiveRelease.Configuration.TermsOfServiceUrl != nil {
				fmt.Fprintf(w, "Terms of Service URL:\t%s\n", *miniapp.ActiveRelease.Configuration.TermsOfServiceUrl)
			}
			w.Flush()
			fmt.Println()
		}
	}
}

func MiniappReleases(releases []birdsdk.MiniProgramRelease) {
	for _, release := range releases {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		// Basic Info
		fmt.Printf("%s\n\n", color.CyanString("📦 Release Details"))
		if release.Uuid != nil {
			fmt.Fprintf(w, "UUID:\t%s\n", *release.Uuid)
		}
		if release.CreatedAt != nil {
			fmt.Fprintf(w, "Created:\t%s\n", release.CreatedAt.Format(time.RFC3339))
		}
		if release.UpdatedAt != nil {
			fmt.Fprintf(w, "Updated:\t%s\n", release.UpdatedAt.Format(time.RFC3339))
		}
		if release.Status != nil {
			fmt.Fprintf(w, "Status:\t%s\n", *release.Status)
		}
		if release.Url != nil {
			fmt.Fprintf(w, "URL:\t%s\n", *release.Url)
		}
		w.Flush()
		fmt.Println()

		// App Info
		if release.AppInfo != nil {
			fmt.Printf("%s\n\n", color.CyanString("ℹ️  App Info"))
			w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			if release.AppInfo.Name != nil {
				fmt.Fprintf(w, "Name:\t%s\n", *release.AppInfo.Name)
			}
			if release.AppInfo.Version != nil {
				fmt.Fprintf(w, "Version:\t%s\n", *release.AppInfo.Version)
			}
			if release.AppInfo.Description != nil {
				fmt.Fprintf(w, "Description:\t%s\n", *release.AppInfo.Description)
			}
			if len(release.AppInfo.Tags) > 0 {
				fmt.Fprintf(w, "Tags:\t%v\n", release.AppInfo.Tags)
			}
			w.Flush()
			fmt.Println()
		}

		// Appearance
		if release.Appearance != nil {
			fmt.Printf("%s\n\n", color.CyanString("🎨 Appearance"))
			w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			if release.Appearance.BackgroundColor != nil {
				fmt.Fprintf(w, "Background Color:\t%s\n", *release.Appearance.BackgroundColor)
			}
			if release.Appearance.ForegroundColor != nil {
				fmt.Fprintf(w, "Foreground Color:\t%s\n", *release.Appearance.ForegroundColor)
			}
			if release.Appearance.NavBackgroundColor != nil {
				fmt.Fprintf(w, "Nav Background Color:\t%s\n", *release.Appearance.NavBackgroundColor)
			}
			if release.Appearance.NavTextColor != nil {
				fmt.Fprintf(w, "Nav Text Color:\t%s\n", *release.Appearance.NavTextColor)
			}
			w.Flush()
			fmt.Println()
		}

		// Configuration
		if release.Configuration != nil {
			fmt.Printf("%s\n\n", color.CyanString("⚙️  Configuration"))
			w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			if release.Configuration.DefaultLanguage != nil {
				fmt.Fprintf(w, "Default Language:\t%s\n", *release.Configuration.DefaultLanguage)
			}
			if release.Configuration.PrivacyPolicyUrl != nil {
				fmt.Fprintf(w, "Privacy Policy URL:\t%s\n", *release.Configuration.PrivacyPolicyUrl)
			}
			if release.Configuration.TermsOfServiceUrl != nil {
				fmt.Fprintf(w, "Terms of Service URL:\t%s\n", *release.Configuration.TermsOfServiceUrl)
			}
			w.Flush()
			fmt.Println()
		}

		fmt.Println(strings.Repeat("-", 80))
		fmt.Println()
	}
}

func MiniappRelease(release *birdsdk.MiniProgramRelease) {
	if release == nil {
		return
	}

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println()

	// Basic Info
	fmt.Printf("%s\n\n", color.CyanString("📱 Release Details"))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	if release.Uuid != nil {
		fmt.Fprintf(w, "UUID:\t%s\n", *release.Uuid)
	}
	if release.CreatedAt != nil {
		fmt.Fprintf(w, "Created:\t%s\n", release.CreatedAt.Format(time.RFC3339))
	}
	if release.UpdatedAt != nil {
		fmt.Fprintf(w, "Updated:\t%s\n", release.UpdatedAt.Format(time.RFC3339))
	}
	if release.Status != nil {
		fmt.Fprintf(w, "Status:\t%s\n", *release.Status)
	}
	if release.Url != nil {
		fmt.Fprintf(w, "URL:\t%s\n", *release.Url)
	}
	if release.AppIcon != nil {
		fmt.Fprintf(w, "App Icon:\t%s\n", *release.AppIcon)
	}
	w.Flush()
	fmt.Println()

	// App Info
	if release.AppInfo != nil {
		fmt.Printf("%s\n\n", color.CyanString("ℹ️  App Info"))
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		if release.AppInfo.Name != nil {
			fmt.Fprintf(w, "Name:\t%s\n", *release.AppInfo.Name)
		}
		if release.AppInfo.Version != nil {
			fmt.Fprintf(w, "Version:\t%s\n", *release.AppInfo.Version)
		}
		if release.AppInfo.Description != nil {
			fmt.Fprintf(w, "Description:\t%s\n", *release.AppInfo.Description)
		}
		if len(release.AppInfo.Tags) > 0 {
			fmt.Fprintf(w, "Tags:\t%v\n", release.AppInfo.Tags)
		}
		w.Flush()
		fmt.Println()
	}

	// Appearance
	if release.Appearance != nil {
		fmt.Printf("%s\n\n", color.CyanString("🎨 Appearance"))
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		if release.Appearance.BackgroundColor != nil {
			fmt.Fprintf(w, "Background Color:\t%s\n", *release.Appearance.BackgroundColor)
		}
		if release.Appearance.ForegroundColor != nil {
			fmt.Fprintf(w, "Foreground Color:\t%s\n", *release.Appearance.ForegroundColor)
		}
		if release.Appearance.NavBackgroundColor != nil {
			fmt.Fprintf(w, "Nav Background Color:\t%s\n", *release.Appearance.NavBackgroundColor)
		}
		if release.Appearance.NavTextColor != nil {
			fmt.Fprintf(w, "Nav Text Color:\t%s\n", *release.Appearance.NavTextColor)
		}
		w.Flush()
		fmt.Println()
	}

	// Configuration
	if release.Configuration != nil {
		fmt.Printf("%s\n\n", color.CyanString("⚙️  Configuration"))
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		if release.Configuration.DefaultLanguage != nil {
			fmt.Fprintf(w, "Default Language:\t%s\n", *release.Configuration.DefaultLanguage)
		}
		if release.Configuration.PrivacyPolicyUrl != nil {
			fmt.Fprintf(w, "Privacy Policy URL:\t%s\n", *release.Configuration.PrivacyPolicyUrl)
		}
		if release.Configuration.TermsOfServiceUrl != nil {
			fmt.Fprintf(w, "Terms of Service URL:\t%s\n", *release.Configuration.TermsOfServiceUrl)
		}
		w.Flush()
		fmt.Println()
	}

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println()

}
