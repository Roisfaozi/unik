package regional

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	// Fixed date: 2023-12-25 15:30:00 UTC
	fixedTime := time.Date(2023, 12, 25, 15, 30, 0, 0, time.UTC)

	tests := []struct {
		name     string
		region   Region
		lang     string
		expected string
	}{
		{"US Format", RegionUS, LangEN, "12/25/2023 03:30 PM"},
		{"EU Format", RegionEU, LangEN, "25/12/2023 15:30"},
		{"ID Format Standard", RegionID, LangID, "25 Desember 2023"},
		{"TH Format Buddhist Era", RegionTH, LangTH, "25/12/2566"}, 
		{"JP Format", RegionJP, LangEN, "2023/12/25"},
		{"CA Format", RegionCA, LangEN, "2023-12-25"},
		{"ISO Format", RegionISO, LangEN, "2023-12-25 15:30:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Format(fixedTime, tt.region, tt.lang, nil)
			if got != tt.expected {
				t.Errorf("Format() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormat_JPEra(t *testing.T) {
	// 2024-05-01 is Reiwa 6
	tm := time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
	// Passing a JapaneseCalendar instance
	got := Format(tm, RegionJP, LangEN, JapaneseCalendar{})
	expected := "Reiwa 6/05/01"
	if got != expected {
		t.Errorf("Format JP Era = %v, want %v", got, expected)
	}
}

func TestFormat_HijriEra(t *testing.T) {
	// July 19, 2023 is approx 1 Muharram 1445
	tm := time.Date(2023, 7, 19, 0, 0, 0, 0, time.UTC)
	
	// We use RegionJP just to trigger the calendar logic in Format function for now
	// In the future, we might have specific formatting for RegionAR or generic
	got := Format(tm, RegionJP, LangEN, HijriCalendar{})

	// Expected: AH 1445/01/01
	expected := "AH 1445/01/01"
	if got != expected {
		t.Errorf("Format Hijri Era = %q, want %q", got, expected)
	}
}

func TestFormatID_EdgeCases(t *testing.T) {
	jan := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := Format(jan, RegionID, LangID, nil); got != "01 Januari 2023" {
		t.Errorf("ID Jan failed: %v", got)
	}

	dec := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	if got := Format(dec, RegionID, LangID, nil); got != "31 Desember 2023" {
		t.Errorf("ID Dec failed: %v", got)
	}
}
