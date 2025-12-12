package regional

import "time"

// HijriCalendar implements CalendarSystem for the Islamic (Hijri) calendar.
// This implementation uses a standard arithmetic (tabular) approximation.
// Note: Actual Islamic dates depend on visual moon sighting (Rukyat) and may vary by 1-2 days.
type HijriCalendar struct{}

func (hc HijriCalendar) Transform(t time.Time) (year int, month int, day int, era string) {
	// 1. Convert Gregorian to Julian Day (JD) using standard integer math
	// Algorithm: Fliegel and Van Flandern (1968)
	gYear := t.Year()
	gMonth := int(t.Month())
	gDay := t.Day()

	a := (14 - gMonth) / 12
	y := gYear + 4800 - a
	m := gMonth + 12*a - 3
	
	// JD at noon
	jd := gDay + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045

	// 2. Convert JD to Hijri
	// Epoch: July 16, 622 AD (Julian Calendar) = JD 1948440
	// Epoch: July 16, 622 AD (Julian Calendar) = JD 1948440
	// Standard Tabular Islamic Calendar (Kuwaiti Algorithm)
	days := jd - 1948440

	// 30-year cycle: 10631 days
	cycle := days / 10631
	days = days % 10631

	hYear := 1 + (int(cycle) * 30)

	// Find year within cycle
	for {
		daysInYear := 354
		if isHijriLeap(hYear) {
			daysInYear = 355
		}
		if days < daysInYear {
			break
		}
		days -= daysInYear
		hYear++
	}

	// Find month
	hMonth := 1
	for {
		daysInMonth := 29
		// Odd months have 30 days, Even have 29
		// Except 12th month in leap year has 30
		if hMonth%2 != 0 {
			daysInMonth = 30
		} else {
			// Month 12 in leap year
			if hMonth == 12 && isHijriLeap(hYear) {
				daysInMonth = 30
			}
		}

		if days < daysInMonth {
			break
		}
		days -= daysInMonth
		hMonth++
	}

	hDay := days + 1 // days is 0-indexed remainder

	return hYear, hMonth, hDay, "AH"
}

// isHijriLeap checks if year is leap in the 30-year tabular cycle (Kuwaiti / Type II)
// Leap years: 2, 5, 7, 10, 13, 16, 18, 21, 24, 26, 29
func isHijriLeap(year int) bool {
	mod := year % 30
	switch mod {
	case 2, 5, 7, 10, 13, 16, 18, 21, 24, 26, 29:
		return true
	}
	// Note: Year 30 is NOT leap in Type II? 
	// Standard Type II (Kuwaiti): 2, 5, 7, 10, 13, 16, 18, 21, 24, 26, 29
	// Some variations include 15 instead of 16. We stick to Kuwaiti.
	return false
}
