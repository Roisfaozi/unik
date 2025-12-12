package regional

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		region   Region
		wantYear int
		wantErr  bool
	}{
		{"Parse US", "12/25/2023 03:30 PM", RegionUS, 2023, false},
		{"Parse EU", "25/12/2023 15:30", RegionEU, 2023, false},
		{"Parse CA", "2023-12-25", RegionCA, 2023, false},
		{"Parse ID", "25 Desember 2023", RegionID, 2023, false},
		{"Parse TH", "25/12/2566", RegionTH, 2023, false}, 
		{"Parse JP", "2023/12/25", RegionJP, 2023, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input, tt.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.Year() != tt.wantYear {
				t.Errorf("Parse() year = %v, want %v", got.Year(), tt.wantYear)
			}
		})
	}
}
