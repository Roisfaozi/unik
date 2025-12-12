package timestamp

import (
	"time"

	"github.com/Roisfaozi/unik/timestamp/regional"
	"github.com/Roisfaozi/unik/timestamp/smart"
	"github.com/Roisfaozi/unik/timestamp/util"
)

// Smart converts a Unix timestamp to a human-readable string using adaptive formatting.
// It automatically selects the most appropriate format (e.g., "just now", "2 hours ago", "Dec 25, 2023").
//
// Example:
//
//	now := time.Now()
//	fmt.Println(Smart(now.Unix())) // Output: "just now" (or similar)
//
//	oneHourAgo := now.Add(-time.Hour)
//	fmt.Println(Smart(oneHourAgo.Unix())) // Output: "1 hour ago" (or similar)
func Smart(unix int64, opts ...Option) string {
	cfg := resolveConfig(opts...)
	t := util.Normalize(UnixToTime(unix), cfg.DefaultTimezone)
	return smart.Adaptive(t, cfg.Language)
}

// Social returns a relative time string (e.g., "2 hours ago", "in 5 minutes")
// Ideal for status updates and notifications.
//
// Example:
//
//	now := time.Now()
//	fmt.Println(Social(now.Unix())) // Output: "just now" (or similar)
//
//	oneHourAgo := now.Add(-time.Hour)
//	fmt.Println(Social(oneHourAgo.Unix())) // Output: "1 hour ago" (or similar)
func Social(unix int64, opts ...Option) string {
	cfg := resolveConfig(opts...)
	t := util.Normalize(UnixToTime(unix), cfg.DefaultTimezone)
	return smart.Social(t, cfg.Language, smart.StyleStandard)
}

// SocialShort returns a compact relative time string (e.g., "2h", "5m")
// Ideal for mobile interfaces with limited space.
//
// Example:
//
//	now := time.Now()
//	fmt.Println(SocialShort(now.Unix())) // Output: "just now" (or similar)
//
//	oneHourAgo := now.Add(-time.Hour)
//	fmt.Println(SocialShort(oneHourAgo.Unix())) // Output: "1 hour ago" (or similar)
func SocialShort(unix int64, opts ...Option) string {
	cfg := resolveConfig(opts...)
	t := util.Normalize(UnixToTime(unix), cfg.DefaultTimezone)
	return smart.Social(t, cfg.Language, smart.StyleShort)
}

// Regional formats a Unix timestamp into a localized date and time string based on a specified region.
//
// Example:
//
//	unixTimestamp := time.Date(2023, time.January, 1, 15, 30, 0, 0, time.UTC).Unix() // 1672587000
//	regionalDate := Regional(unixTimestamp, regional.US)
//	fmt.Println(regionalDate) // Output: 1/1/2023, 3:30 PM
func Regional(unix int64, region regional.Region, opts ...Option) string {
	cfg := resolveConfig(opts...)
	t := util.Normalize(UnixToTime(unix), cfg.DefaultTimezone)
	return regional.Format(t, region, cfg.Language)
}

// FullDateTime formats a Unix timestamp into a verbose date and time string.
// The format used is "Monday, January 2, 2006 15:04:05 MST".
//
// Example:
//
//	unixTimestamp := time.Date(2023, time.January, 1, 15, 30, 0, 0, time.UTC).Unix() // 1672587000
//	fullDateTime := FullDateTime(unixTimestamp)
//	fmt.Println(fullDateTime) // Output: Sunday, January 1, 2023 15:30:00 UTC
func FullDateTime(unix int64, opts ...Option) string {
	cfg := resolveConfig(opts...)
	t := util.Normalize(UnixToTime(unix), cfg.DefaultTimezone)
	return t.Format("Monday, January 2, 2006 15:04:05 MST")
}

func Formal(unix int64, opts ...Option) string {
	cfg := resolveConfig(opts...)
	t := util.Normalize(UnixToTime(unix), cfg.DefaultTimezone)
	return t.Format("2006-01-02 15:04:05 -0700 MST")
}

// Custom formats a Unix timestamp into a string using a custom Go time layout.
//
// Example:
//
//	unixTimestamp := time.Date(2023, time.January, 1, 15, 30, 0, 0, time.UTC).Unix() // 1672587000
//	customDate := Custom(unixTimestamp, "02 Jan 2006 15:04 MST")
//	fmt.Println(customDate) // Output: 01 Jan 2023 15:30 UTC

func Custom(unix int64, layout string, opts ...Option) string {
	cfg := resolveConfig(opts...)
	t := util.Normalize(UnixToTime(unix), cfg.DefaultTimezone)
	return t.Format(layout)
}

// ParseRegional parses a date string according to a specific region's format
// and returns its Unix timestamp.
//
// Example:
//
//	unixTimestamp, err := ParseRegional("25/12/2023", regional.ID)
//	if err != nil {
//		log.Fatalf("Error parsing date: %v", err)
//	}
//	fmt.Printf("Unix timestamp: %d\n", unixTimestamp)
//
// This would parse "25/12/2023" as December 25, 2023, based on Indonesia's (ID)
// typical date format.
func ParseRegional(dateStr string, region regional.Region) (int64, error) {
	t, err := regional.Parse(dateStr, region)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}


// ParseWithLayout parses a date string according to a specific layout
// and returns its Unix timestamp.
//
// Example:
//
//	unixTimestamp, err := ParseWithLayout("2023-12-25 15:30:00", "2006-01-02 15:04:05")
//	if err != nil {
//		log.Fatalf("Error parsing date: %v", err)
//	}
//	fmt.Printf("Unix timestamp: %d\n", unixTimestamp)
func ParseWithLayout(dateStr, layout string) (int64, error) {
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// Duration formats a duration in seconds into a detailed human-readable string.
// Examples: "3 seconds", "1 minute 40 seconds", "2 hours 20 minutes".
func Duration(seconds int64, opts ...Option) string {
	cfg := resolveConfig(opts...)
	d := time.Duration(seconds) * time.Second
	return smart.Duration(d, cfg.Language)
}
