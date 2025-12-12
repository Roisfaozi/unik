package smart

import (
	"fmt"
	"math"
	"time"
)

type RelativeStyle int

const (
	StyleStandard RelativeStyle = iota 
	StyleShort                        
)

var (
	translations = map[string]map[string]string{
		"en": {
			"just_now": "just now",
			"ago":      "ago",
			"in":       "in",
			"s":        "s",
			"m":        "m",
			"h":        "h",
			"d":        "d",
			"y":        "y",
			"sec":      "seconds",
			"min":      "minutes",
			"hour":     "hours",
			"day":      "days",
			"year":     "years",
		},
		"id": {
			"just_now": "baru saja",
			"ago":      "lalu",
			"in":       "dalam",
			"s":        "dtk",
			"m":        "mnt",
			"h":        "j",
			"d":        "h",
			"y":        "thn",
			"sec":      "detik",
			"min":      "menit",
			"hour":     "jam",
			"day":      "hari",
			"year":     "tahun",
		},
	}
)
// GetTrans retrieves a translated string for a given language and key.
// If the language or key is not found, it falls back to English.
//
// Example:
//
//	s := GetTrans("en", "just_now") // s will be "just now"
//	s = GetTrans("id", "just_now") // s will be "baru saja"
//	s = GetTrans("fr", "just_now") // s will be "just now" (falls back to en)
func GetTrans(lang, key string) string {
	if dict, ok := translations[lang]; ok {
		if val, ok := dict[key]; ok {
			return val
		}
	}
	return translations["en"][key]
}

// Social formats a time.Time into a human-readable relative time string (e.g., "just now", "5 minutes ago", "in 2 days").
// It supports different languages and formatting styles (short or standard).
//
// Parameters:
//   - t: The time to be formatted.
//   - lang: The language code for translation (e.g., "en", "id").
//   - style: The desired relative time style (StyleStandard or StyleShort).
//
// Returns:
//   A string representing the relative time.
//
// Example:
//
//	now := time.Now()
//	inFiveMinutes := now.Add(5 * time.Minute)
//	fiveMinutesAgo := now.Add(-5 * time.Minute)
//
//	fmt.Println(Social(now, "en", StyleStandard))       // Output: "just now"
//	fmt.Println(Social(fiveMinutesAgo, "en", StyleStandard)) // Output: "5 minutes ago"
//	fmt.Println(Social(inFiveMinutes, "en", StyleStandard))  // Output: "in 5 minutes"
//	fmt.Println(Social(fiveMinutesAgo, "en", StyleShort))    // Output: "5m"
//	fmt.Println(Social(fiveMinutesAgo, "id", StyleStandard)) // Output: "5 menit lalu"
func Social(t time.Time, lang string, style RelativeStyle) string {
	now := time.Now().In(t.Location())
	diff := now.Sub(t)
	seconds := math.Abs(diff.Seconds())

	isPast := diff >= 0

	minute := 60.0
	hour := 3600.0
	day := 86400.0
	year := 31536000.0

	if seconds < 10 {
		return GetTrans(lang, "just_now")
	}

	var val int
	var unit string
	var unitShort string

	if seconds < minute {
		val = int(seconds)
		unit = "sec"
		unitShort = "s"
	} else if seconds < hour {
		val = int(seconds / minute)
		unit = "min"
		unitShort = "m"
	} else if seconds < day {
		val = int(seconds / hour)
		unit = "hour"
		unitShort = "h"
	} else if seconds < year {
		val = int(seconds / day)
		unit = "day"
		unitShort = "d"
	} else {
		val = int(seconds / year)
		unit = "year"
		unitShort = "y"
	}

	if style == StyleShort {
		return fmt.Sprintf("%d%s", val, GetTrans(lang, unitShort))
	}

	term := GetTrans(lang, unit)
	if isPast {
		return fmt.Sprintf("%d %s %s", val, term, GetTrans(lang, "ago"))
	}
	return fmt.Sprintf("%s %d %s", GetTrans(lang, "in"), val, term)
}
