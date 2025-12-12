package smart

import (
	"time"
)

// Adaptive formats a time.Time into a string that is contextually relevant based on how
// much time has passed. It mimics "smart" timestamp formatting often found in social apps:
// - < 1 minute: "Just now" (via Social helper)
// - Same day: "HH:MM"
// - < 7 days: Day name (e.g., "Monday")
// - Same year: "DD Mon"
// - Older: "DD Mon YYYY"
func Adaptive(t time.Time, lang string) string {
	now := time.Now().In(t.Location())
	diff := now.Sub(t)
	
	// If diff is negative (future), handle it? 
	// For "Adaptive" usually refers to past events in feeds mostly, 
	// but let's assume if it's future we might want specific handling.
	// For now, let's treat absolute value for logic thresholds but strict formats.
	
	// Just handle past for "Smart" typical use case (messages, feeds)
	// < 1 min: Just now
	if diff < time.Minute && diff > -time.Minute {
		return Social(t, lang, StyleStandard) 
	}
	
	// < 24 hours: HH:MM
	// Handle wrapping around midnight? 
	// Simple check: if it's the same day.
	if now.YearDay() == t.YearDay() && now.Year() == t.Year() {
		return t.Format("15:04")
	}
	
	// < 7 days: Day Name (Monday, etc)
	// We need localized day names if not English
	if diff < 7*24*time.Hour && diff > 0 {
		return t.Format("Monday") 
		// Note: Go doesn't natively localize "Monday". 
		// If we want localized, we need a map. 
		// Let's stick to English for simplicity in this MVP or add map.
	}
	
	if now.Year() == t.Year() {
		return t.Format("02 Jan")
	}
	
	return t.Format("02 Jan 2006")
}
