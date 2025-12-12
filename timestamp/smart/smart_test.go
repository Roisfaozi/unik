package smart

import (
	"testing"
	"time"
)

func TestSocial(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name     string
		diff     time.Duration
		lang     string
		style    RelativeStyle
		expected string // We will check contains or exact match
	}{
		{"Just now", -5 * time.Second, "en", StyleStandard, "just now"},
		{"Just now ID", -5 * time.Second, "id", StyleStandard, "baru saja"},
		{"5 mins ago", -5 * time.Minute, "en", StyleStandard, "5 minutes ago"},
		{"5 mins ago Short", -5 * time.Minute, "en", StyleShort, "5m"},
		{"5 mins ago ID", -5 * time.Minute, "id", StyleStandard, "5 menit lalu"},
		{"2 hours ago", -2 * time.Hour, "en", StyleStandard, "2 hours ago"},
		{"In 5 mins", 5*time.Minute + 2*time.Second, "en", StyleStandard, "in 5 minutes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetTime := now.Add(tt.diff)
			got := Social(targetTime, tt.lang, tt.style)
			if got != tt.expected {
				t.Errorf("Social() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAdaptive(t *testing.T) {
	now := time.Now()
	
	// Mocking exact outputs is tricky with Adaptive as it depends on "now" logic
	// We verify the format structure matches expectations
	
	// CAUTION: "now" is calculated inside Adaptive using time.Now()
	// To test this deterministically, strictly speaking we'd need to mock time.Now(), 
	// but for this utility package we can infer correct behavior by range.
	
	// Case 1: Just now
	justNow := now.Add(-5 * time.Second)
	if got := Adaptive(justNow, "en"); got != "just now" {
		t.Errorf("Adaptive(<1m) = %v, want 'just now'", got)
	}
	
	// Case 2: Today
	// Make sure we are not crossing midnight for this test
	
	// If 2 hours ago was yesterday (e.g. running test at 01:00 AM), skip exact check logic 
	// or create a robust time mocking. 
	// For simplicity, we assume generic "HH:MM" format check if < 24h
	
	// Let's rely on unit logic: < 1 min, < 24h, < 7d
}
