package regional

import "time"

// CalendarSystem defines the interface for converting a Gregorian time
// into a specific cultural calendar era and year.
type CalendarSystem interface {
	// Transform converts a Gregorian time into the specific calendar's date components.
	// Returns: year, month, day, and era name (optional).
	Transform(t time.Time) (year int, month int, day int, era string)
}
