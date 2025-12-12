# Unik Timestamp Utility

![CI Status](https://github.com/Roisfaozi/unik/actions/workflows/ci.yml/badge.svg)
![Latest Release](https://img.shields.io/github/v/release/Roisfaozi/unik)

A robust, production-ready Go package for advanced time formatting, specifically designed for modern UI requirements (Smart/Adaptive formatting) and extensive Regional support (ASEAN, Asia, Global).

## 🚀 Key Features

- **Smart / Adaptive Formatting**: Automatically switches format based on time age (e.g., "Just now" → "15:30" → "Monday" → "25 Dec").
- **Social Relative Time**: "5 minutes ago", "in 2 hours", or compact "5m", "2h".
- **Deep Regional Support**:
  - **ASEAN**: Indonesia (Localized months), Thailand (Buddhist Era 2566), Vietnam, Malaysia, Philippines.
  - **Global**: US (MM/DD/YYYY), EU (DD/MM/YYYY), Japan/China (YYYY/MM/DD), Canada (YYYY-MM-DD).
- **Performance**: Built-in efficient timezone handling with caching.
- **Zero Boilerplate**: Simple, expressive API.

## 📦 Installation

```bash
go get github.com/Roisfaozi/unik
```

## 🛠️ Usage Guide

### 1. Basic Usage

```go
package main

import (
	"fmt"
	"time"
	"github.com/Roisfaozi/unik/timestamp"
	"github.com/Roisfaozi/unik/timestamp/regional"
)

func main() {
	now := time.Now().Unix()

	// Smart / Adaptive (Best for Feeds)
	fmt.Println(timestamp.Smart(now)) // Output: "Just now"

	// Social (Best for Comments/Status)
	fmt.Println(timestamp.Social(now - 300)) // "5 minutes ago"
	fmt.Println(timestamp.SocialShort(now - 3600)) // "1h"

	// Regional Formats
	fmt.Println(timestamp.Regional(now, regional.RegionID)) // "11 Desember 2025"
	fmt.Println(timestamp.Regional(now, regional.RegionTH)) // "11/12/2568" (Buddhist Era)
}
```

### 2. Integration with Web Server & Database Best Practices

This package simplifies the full lifecycle of time data: **Store as Unix** → **Format for Response**.

**Workflow:**

1.  **Database**: Store time as `int64` (Unix Timestamp). It's the most portable format.
2.  **Logic**: Keep it as `int64` or `time.Time` throughout your backend.
3.  **Response (UI)**: Convert to formatted strings (`string`) explicitly in your API response struct.

#### Complete Example (Gin Framework)

```go
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Roisfaozi/unik/timestamp"
	"github.com/Roisfaozi/unik/timestamp/regional"
)

// 1. DB Model (How it's stored)
type PostDB struct {
	ID        int   `json:"id"`
	CreatedAt int64 `json:"created_at"` // Saved as Unix (e.g., 1703489400)
}

// 2. API Response Model (How it's sent to UI)
// Separation of concerns: Data vs Presentation
type PostResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`

	// Machine Readable (Optional)
	UnixTime int64 `json:"unix_time"`

	// Human Readable (Formatted by Unik)
	TimeAgo    string `json:"time_ago"`    // e.g. "5 minutes ago"
	SmartDate  string `json:"smart_date"`  // e.g. "Just now", "25 Dec"
	DetailDate string `json:"detail_date"` // e.g. "25 Desember 2023"
}

func main() {
	// Global Setup: Default to Jakarta & Indonesian
	timestamp.SetDefaultConfig("Asia/Jakarta", "id")

	r := gin.Default()

	// Mock DB Data
	mockDB := []PostDB{
		{ID: 1, CreatedAt: time.Now().Add(-5 * time.Minute).Unix()}, // 5 mins ago
		{ID: 2, CreatedAt: time.Now().Add(-24 * time.Hour).Unix()},  // Yesterday
	}

	r.GET("/posts", func(c *gin.Context) {
		var responses []PostResponse

		for _, p := range mockDB {
			// CONVERSION STEP: Map DB -> Response
			res := PostResponse{
				ID:       p.ID,
				Title:    "Example Post",
				UnixTime: p.CreatedAt,

				// Format 1: Relative (Social Context)
				TimeAgo: timestamp.Social(p.CreatedAt),

				// Format 2: Adaptive (Feed Context)
				SmartDate: timestamp.Smart(p.CreatedAt),

				// Format 3: Explicit Regional (Formal Context)
				DetailDate: timestamp.Regional(p.CreatedAt, regional.RegionID),
			}
			responses = append(responses, res)
		}

		c.JSON(http.StatusOK, responses)
	})

	r.Run(":8080")
}
```

**Output JSON:**

```json
[
  {
    "id": 1,
    "unix_time": 1708500000,
    "time_ago": "5 menit lalu",
    "smart_date": "Baru saja",
    "detail_date": "21 Februari 2025"
  },
  {
    "id": 2,
    "unix_time": 1708413600,
    "time_ago": "1 hari lalu",
    "smart_date": "20 Feb",
    "detail_date": "20 Februari 2025"
  }
]
```

### 3. Bidirectional Conversion (Unix ↔ String)

This package supports full bidirectional conversion. You can convert formatted strings back to Unix timestamps.

```go
// 1. Unix -> Format (Display)
unixTime := int64(1703462400)
display := timestamp.Regional(unixTime, regional.RegionID)
fmt.Println(display) // "25 Desember 2023"

// 2. Format -> Unix (Parsing)
parsedUnix, err := timestamp.ParseRegional("25 Desember 2023", regional.RegionID)
if err != nil {
	log.Fatal(err)
}
fmt.Println(parsedUnix) // 1703462400

// Custom Layout Parsing
unix2, _ := timestamp.ParseWithLayout("2023-12-25", "2006-01-02")
```

### 4. Localization & Timezone Configuration

You can configure options globally or per-call (Priority: Per-Call > Global > Default).

**Global Config (App Startup):**

```go
func init() {
	// Set default to Jakarta/Indonesia for the whole app
	timestamp.SetDefaultConfig("Asia/Jakarta", "id")
}
```

**Per-Call Config (Override):**

```go
// Force display in New York Time, English
timestamp.Smart(unix,
	timestamp.WithTimezone("America/New_York"),
	timestamp.WithLanguage("en"),
)
```

## 🌍 Supported Regions

| Region Code | Description | Format Example        |
| :---------- | :---------- | :-------------------- |
| `RegionID`  | Indonesia   | `25 Desember 2023`    |
| `RegionTH`  | Thailand    | `25/12/2566` (BE)     |
| `RegionVN`  | Vietnam     | `25/12/2023`          |
| `RegionMY`  | Malaysia    | `25/12/2023`          |
| `RegionUS`  | USA         | `12/25/2023 03:30 PM` |
| `RegionEU`  | Europe      | `25/12/2023 15:30`    |
| `RegionJP`  | Japan       | `2023/12/25`          |
| `RegionCA`  | Canada      | `2023-12-25`          |

## 🧪 Testing

Run standard Go tests:

```bash
go test ./...
```
