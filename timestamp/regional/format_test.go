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
			got := Format(fixedTime, tt.region, tt.lang)
			if got != tt.expected {
				t.Errorf("Format() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFormatID_EdgeCases(t *testing.T) {
	jan := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := Format(jan, RegionID, LangID); got != "01 Januari 2023" {
		t.Errorf("ID Jan failed: %v", got)
	}

	dec := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	if got := Format(dec, RegionID, LangID); got != "31 Desember 2023" {
		t.Errorf("ID Dec failed: %v", got)
	}
}
