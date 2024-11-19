package formatting

import (
	"fmt"
	"time"
)

// formatRelativeTime converts a time.Time to a human-readable relative time string.
func FormatRelativeTimeWithExpired(t time.Time) string {
	duration := time.Until(t)
	if duration < 0 {
		return "Expired"
	}

	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	if days > 0 {
		return fmt.Sprintf("%d day(s) from now", days)
	} else if hours > 0 {
		return fmt.Sprintf("%d hour(s) from now", hours)
	} else {
		return fmt.Sprintf("%d minute(s) from now", minutes)
	}
}

// FormatRelativeTime converts a time.Time to a human-readable relative time string, handling both past and future times.
func FormatRelativeTime(t time.Time) string {
	duration := time.Until(t)
	if duration < 0 {
		duration = -duration
		days := int(duration.Hours()) / 24
		hours := int(duration.Hours()) % 24
		minutes := int(duration.Minutes()) % 60

		if days > 0 {
			return fmt.Sprintf("%d day(s) ago", days)
		} else if hours > 0 {
			return fmt.Sprintf("%d hour(s) ago", hours)
		} else {
			return fmt.Sprintf("%d minute(s) ago", minutes)
		}
	}

	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	if days > 0 {
		return fmt.Sprintf("%d day(s) from now", days)
	} else if hours > 0 {
		return fmt.Sprintf("%d hour(s) from now", hours)
	} else {
		return fmt.Sprintf("%d minute(s) from now", minutes)
	}
}
