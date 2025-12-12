package main

import (
	"fmt"
	"time"

	"github.com/Roisfaozi/unik/timestamp"
	"github.com/Roisfaozi/unik/timestamp/regional"
)

func main() {
	now := time.Now()
	unixNow := now.Unix()
	
	// 5 hours ago
	fiveHoursAgo := now.Add(-5 * time.Hour).Unix()
	
	// 3 days ago
	threeDaysAgo := now.Add(-24 * 3 * time.Hour).Unix()
	
	// 6 months ago
	sixMonthsAgo := now.Add(-24 * 30 * 6 * time.Hour).Unix()

	fmt.Println("--- Timestamp Utility Demo ---")
	
	fmt.Println("\n1. Smart Formatting (Adaptive):")
	fmt.Printf("Now:          %s\n", timestamp.Smart(unixNow))
	fmt.Printf("5h Ago:       %s\n", timestamp.Smart(fiveHoursAgo)) 
	fmt.Printf("3 Days Ago:   %s\n", timestamp.Smart(threeDaysAgo))  
	fmt.Printf("6 Months Ago: %s\n", timestamp.Smart(sixMonthsAgo)) 

	fmt.Println("\n2. Social Formatting (Relative):")
	fmt.Printf("Standard (EN): %s\n", timestamp.Social(fiveHoursAgo))
	fmt.Printf("Short (EN):    %s\n", timestamp.SocialShort(fiveHoursAgo))
	fmt.Printf("Standard (ID): %s\n", timestamp.Social(fiveHoursAgo, timestamp.WithLanguage("id")))

	fmt.Println("\n3. Regional Formatting:")
	// Indonesia
	fmt.Printf("Indonesia (ID): %s\n", timestamp.Regional(unixNow, regional.RegionID, 
		timestamp.WithTimezone("Asia/Jakarta"),
		timestamp.WithLanguage("id")))
	
	// Thailand
	fmt.Printf("Thailand (TH):  %s (Buddhist Era)\n", timestamp.Regional(unixNow, regional.RegionTH,
		timestamp.WithTimezone("Asia/Bangkok")))
	
	// US
	fmt.Printf("USA (US):       %s\n", timestamp.Regional(unixNow, regional.RegionUS,
		timestamp.WithTimezone("America/New_York")))
		
	// Japan
	fmt.Printf("Japan (JP):     %s\n", timestamp.Regional(unixNow, regional.RegionJP,
		timestamp.WithTimezone("Asia/Tokyo")))
	
	// Japan (Native Era)
	fmt.Printf("Japan (Era):    %s\n", timestamp.Regional(unixNow, regional.RegionJP,
		timestamp.WithTimezone("Asia/Tokyo"),
		timestamp.WithCalendar(regional.JapaneseCalendar{})))

	fmt.Println("\n4. Formal/Legal:")
	fmt.Printf("Formal: %s\n", timestamp.Formal(unixNow, timestamp.WithTimezone("Asia/Jakarta")))

	fmt.Println("\n5. Duration:")
	fmt.Printf("Seconds (100s): %s\n", timestamp.Duration(100))
	fmt.Printf("Complex (8400s): %s\n", timestamp.Duration(8400))
}
