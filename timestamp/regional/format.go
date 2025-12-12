package regional

import (
	"fmt"
	"time"
)

var monthsID = []string{
	"", "Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

// Format formats a time.Time object based on the specified region and language.
// It provides different date and time formats for various regions.
//
// Example:
//   t := time.Date(2023, time.November, 15, 10, 30, 0, 0, time.UTC)
//   fmt.Println(Format(t, RegionUS, ""))     // Output: "11/15/2023 10:30 AM"
//   fmt.Println(Format(t, RegionEU, ""))     // Output: "15/11/2023 10:30"
//   fmt.Println(Format(t, RegionID, LangID)) // Output: "15 November 2023"
//   fmt.Println(Format(t, RegionTH, ""))     // Output: "15/11/2566" (assuming 2023 + 543 = 2566 BE)
func Format(t time.Time, region Region, lang string) string {
	switch region {
	case RegionUS:
		return t.Format("01/02/2006 03:04 PM")
	case RegionEU:
		return t.Format("02/01/2006 15:04")
	case RegionCA:
		return t.Format("2006-01-02")
	case RegionID:
		if lang == LangID {
			return formatID(t)
		}

		return formatID(t) 
	case RegionTH:
		return formatTH(t)
	case RegionVN:
		return t.Format("02/01/2006")
	case RegionMY, RegionSG:
		return t.Format("02/01/2006")
	case RegionPH:
		return t.Format("01/02/2006")
	case RegionJP:
		return t.Format("2006/01/02")
	case RegionKR:
		return t.Format("2006.01.02")
	case RegionCN:
		return t.Format("2006-01-02")
	case RegionISO:
		return t.Format("2006-01-02 15:04:05")
	default:
		return t.Format(time.RFC3339)
	}
}

func formatID(t time.Time) string {
	d := t.Day()
	m := t.Month()
	y := t.Year()
	
	mIdx := int(m)
	if mIdx < 1 || mIdx > 12 {
		return t.Format("02 Jan 2006")
	}

	return fmt.Sprintf("%02d %s %d", d, monthsID[mIdx], y)
}

// formatTH formats a Thai date string (e.g., "25/10/2023") into a time.Time object.
// It handles Thai date format by converting it to a standard format before parsing.
//
// Example:
//
//	timeTH, err := parseTH("25/10/2023")
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(timeTH) // 2023-10-25 00:00:00 +0000 UTC
func formatTH(t time.Time) string {
	be := t.Year() + 543
	return fmt.Sprintf("%02d/%02d/%d", t.Day(), t.Month(), be)
}
