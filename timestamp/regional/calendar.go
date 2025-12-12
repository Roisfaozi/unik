package regional

import "time"

// CalendarSystem defines the interface for converting a Gregorian time
// into a specific cultural calendar era and year.
type CalendarSystem interface {
	// Transform returns the era name and the year within that era for a given time.
	// Example for Japanese Calendar (2024): era="Reiwa", year=6
	Transform(t time.Time) (era string, year int)
}
