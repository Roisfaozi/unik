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

func (jc JapaneseCalendar) Transform(t time.Time) (year int, month int, day int, era string) {
	m := int(t.Month())
	d := t.Day()

	// Check eras from newest to oldest
	if !t.Before(dateReiwa) {
		return t.Year() - 2018, m, d, "Reiwa"
	}
	if !t.Before(dateHeisei) {
		return t.Year() - 1988, m, d, "Heisei"
	}
	if !t.Before(dateShowa) {
		return t.Year() - 1925, m, d, "Showa"
	}
	if !t.Before(dateTaisho) {
		return t.Year() - 1911, m, d, "Taisho"
	}
	if !t.Before(dateMeiji) {
		return t.Year() - 1867, m, d, "Meiji"
	}

	// Fallback for pre-Meiji
	return t.Year(), m, d, "Seireki"
}
