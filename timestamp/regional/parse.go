package regional

import (
	"fmt"
	"strings"
	"time"
)

// Parse parses a date string according to the specified region's format.
// It returns a time.Time object and an error if parsing fails.
//
// Example:
//
//	timeUS, err := Parse("10/25/2023 03:30 PM", RegionUS)
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(timeUS) // 2023-10-25 15:30:00 +0000 UTC
//
//	timeEU, err := Parse("25/10/2023 15:30", RegionEU)
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(timeEU) // 2023-10-25 15:30:00 +0000 UTC
//
//	timeCA, err := Parse("2023-10-25", RegionCA)
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(timeCA) // 2023-10-25 00:00:00 +0000 UTC
func Parse(dateStr string, region Region) (time.Time, error) {
	var layout string

	switch region {
	case RegionUS:
		layout = "01/02/2006 03:04 PM"
	case RegionEU:
		layout = "02/01/2006 15:04"
	case RegionCA:
		layout = "2006-01-02"
	case RegionID:

		return parseID(dateStr)
	case RegionTH:
		return parseTH(dateStr)
	case RegionVN, RegionMY, RegionSG:
		layout = "02/01/2006"
	case RegionPH:
		layout = "01/02/2006"
	case RegionJP:
		layout = "2006/01/02"
	case RegionKR:
		layout = "2006.01.02"
	case RegionCN:
		layout = "2006-01-02"
	case RegionISO:
		layout = "2006-01-02 15:04:05"
	default:
		return time.Time{}, fmt.Errorf("unsupported region for parsing: %s", region)
	}

	return time.Parse(layout, dateStr)
}

// parseID parses an Indonesian date string (e.g., "01 Januari 2006") into a time.Time object.
// It handles Indonesian month names by converting them to English before parsing.
//
// Example:
//
//	timeID, err := parseID("25 April 2023")
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(timeID) // 2023-04-25 00:00:00 +0000 UTC
func parseID(dateStr string) (time.Time, error) {
	replacer := strings.NewReplacer(
		"Januari", "January",
		"Februari", "February",
		"Maret", "March",
		"April", "April",
		"Mei", "May",
		"Juni", "June",
		"Juli", "July",
		"Agustus", "August",
		"September", "September",
		"Oktober", "October",
		"November", "November",
		"Desember", "December",
	)
	normalized := replacer.Replace(dateStr)
	return time.Parse("02 January 2006", normalized)
}

// parseTH parses a Thai date string (e.g., "25/10/2023") into a time.Time object.
// It handles Thai date format by converting it to a standard format before parsing.
//
// Example:
//
//	timeTH, err := parseTH("25/10/2023")
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(timeTH) // 2023-10-25 00:00:00 +0000 UTC
func parseTH(dateStr string) (time.Time, error) {
	
    parts := strings.Split(dateStr, "/")
    if len(parts) != 3 {
        return time.Time{}, fmt.Errorf("invalid format for TH")
    }
    
    day := parts[0]
    month := parts[1]
    yearBE := parts[2]
    
    var yBE int
    if _, err := fmt.Sscanf(yearBE, "%d", &yBE); err != nil {
        return time.Time{}, fmt.Errorf("invalid year BE: %w", err)
    }
    yAD := yBE - 543
    
    isoStr := fmt.Sprintf("%04d-%s-%s", yAD, month, day)
    return time.Parse("2006-01-02", isoStr)
}
