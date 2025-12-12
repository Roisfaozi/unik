package smart

import (
	"fmt"
	"strings"
	"time"
)

// Duration formats a time.Duration into a verbose string (e.g., "2 hours 30 minutes").
// It breaks down time into hours, minutes, and seconds.
// Zero units are omitted (e.g., "1 hour" instead of "1 hour 0 minutes").
func Duration(d time.Duration, lang string) string {
	seconds := int(d.Seconds())
	if seconds == 0 {
		return "0 " + GetPlural(lang, "sec", 0)
	}

	h := seconds / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60

	var parts []string

	// Hours
	if h > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", h, GetPlural(lang, "hour", h)))
	}

	// Minutes
	if m > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", m, GetPlural(lang, "min", m)))
	}

	// Seconds
	if s > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", s, GetPlural(lang, "sec", s)))
	}

	return strings.Join(parts, " ")
}
