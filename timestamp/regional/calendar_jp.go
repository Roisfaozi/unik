package regional

import "time"

// JapaneseCalendar implements CalendarSystem for Japan's Gengō (Era) system.
// Supports modern eras: Meiji, Taisho, Showa, Heisei, and Reiwa.
type JapaneseCalendar struct{}

var (
	// Era Start Dates
	dateReiwa  = time.Date(2019, 5, 1, 0, 0, 0, 0, time.UTC)
	dateHeisei = time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC)
	dateShowa  = time.Date(1926, 12, 25, 0, 0, 0, 0, time.UTC)
	dateTaisho = time.Date(1912, 7, 30, 0, 0, 0, 0, time.UTC)
	dateMeiji  = time.Date(1868, 1, 25, 0, 0, 0, 0, time.UTC) // Approximate Gregorian start
)

func (jc JapaneseCalendar) Transform(t time.Time) (era string, year int) {
	// Check eras from newest to oldest
	if !t.Before(dateReiwa) {
		return "Reiwa", t.Year() - 2018
	}
	if !t.Before(dateHeisei) {
		return "Heisei", t.Year() - 1988
	}
	if !t.Before(dateShowa) {
		return "Showa", t.Year() - 1925
	}
	if !t.Before(dateTaisho) {
		return "Taisho", t.Year() - 1911
	}
	if !t.Before(dateMeiji) {
		return "Meiji", t.Year() - 1867
	}

	// Fallback for pre-Meiji (just return Gregorian as a safety or a specific "Pre-Meiji" era)
	return "Seireki", t.Year()
}
