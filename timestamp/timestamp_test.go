package timestamp_test

import (
	"testing"
	"time"

	"github.com/Roisfaozi/unik/timestamp"
	"github.com/Roisfaozi/unik/timestamp/regional"
)

func TestSmart(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name     string
		unix     int64
		expected string
	}{
		{"Current", now.Unix(), "Just now"}, 
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := timestamp.Smart(tt.unix)
			if got != "just now" && got != "Just now" { 
				t.Errorf("Smart() = %v, want 'just now'", got)
			}
		})
	}
}

func TestRegional(t *testing.T) {
	fixedTime := time.Date(2023, 12, 25, 15, 30, 0, 0, time.UTC)
	unix := fixedTime.Unix()

	tests := []struct {
		region   regional.Region
		opts     []timestamp.Option
		expected string
	}{
		{regional.RegionUS, []timestamp.Option{timestamp.WithTimezone("UTC")}, "12/25/2023 03:30 PM"},
		{regional.RegionEU, []timestamp.Option{timestamp.WithTimezone("UTC")}, "25/12/2023 15:30"},
		{regional.RegionJP, []timestamp.Option{timestamp.WithTimezone("UTC")}, "2023/12/25"},
	}

	for _, tt := range tests {
		t.Run(string(tt.region), func(t *testing.T) {
			got := timestamp.Regional(unix, tt.region, tt.opts...)
			if got != tt.expected {
				t.Errorf("Regional(%s) = %v, want %v", tt.region, got, tt.expected)
			}
		})
	}
}

func TestSocial(t *testing.T) {
	now := time.Now()
	fiveMinsAgo := now.Add(-5 * time.Minute).Unix()
	
	got := timestamp.Social(fiveMinsAgo) 
	if got != "5 minutes ago" {
		t.Errorf("Social() = %v, want '5 minutes ago'", got)
	}
}
