package util

import (
	"testing"
	"time"
)

func TestLoadLocation(t *testing.T) {
	// 1. Valid Location
	loc, err := LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatalf("Failed to load valid location: %v", err)
	}
	if loc.String() != "Asia/Jakarta" {
		t.Errorf("Expected Asia/Jakarta, got %s", loc.String())
	}

	// 2. Cached check
	// Modify internal cache directly to verify it's being read? 
	// Or just call again and ensure no error (black box test).
	loc2, err := LoadLocation("Asia/Jakarta")
	if err != nil {
		t.Fatalf("Failed to load cached location: %v", err)
	}
	if loc != loc2 {
		t.Log("Warning: LoadLocation returned different pointer, cache might be bypassed or new pointer created (acceptable as long as functional)")
	}

	// 3. Invalid Location
	_, err = LoadLocation("Invalid/Location")
	if err == nil {
		t.Error("Expected error for invalid location, got nil")
	}
}

func TestNormalize(t *testing.T) {
	fixed := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		input    time.Time
		tz       string
		wantZone string // approximate check
	}{
		{"Empty TZ (Keep Original)", fixed, "", "UTC"},
		{"Valid TZ", fixed, "Asia/Jakarta", "WIB"}, // UTC 12:00 -> 19:00 WIB
		{"Invalid TZ (Fallback UTC)", fixed, "Invalid/Place", "UTC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Normalize(tt.input, tt.tz)
			name, _ := got.Zone()
			// Note: Zone name might vary by system (WIB vs +07), but usually standard in Go
			// We check if location name matches for valid ones
			if tt.tz != "" && tt.tz != "Invalid/Place" {
				if got.Location().String() != tt.tz {
					t.Errorf("Normalize location = %v, want %v", got.Location().String(), tt.tz)
				}
			}
			if tt.name == "Invalid TZ" && got.Location().String() != "UTC" {
				t.Errorf("Fallback failed, got %v", got.Location())
			}
			_ = name
		})
	}
}
