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
		return "0 " + GetTrans(lang, "sec")
	}

	h := seconds / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60

	var parts []string

	// Hours
	if h > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", h, GetTrans(lang, "hour")))
	}

	// Minutes
	if m > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", m, GetTrans(lang, "min")))
	}

	// Seconds
	// Show seconds if it's the only unit (< 1 min) OR if specifically requested?
	// The requirement is "3 seconds", "1m 40s", "2h 20m".
	// Usually if hours exist, seconds are often noise, but for "2h 20m" format implying
	// we show strict breakdown.
	// Let's allow seconds logic: if (h==0 && m==0) show seconds.
	// OR if user specifically wants full precision? 
	// Based on request "1 min 40 sec", seconds should be shown even if minutes exist.
	
	if s > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", s, GetTrans(lang, "sec")))
	}

	return strings.Join(parts, " ")
}
