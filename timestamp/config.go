package timestamp

import (
	"sync"
	"time"
)

type Config struct {
	DefaultTimezone string
	Language        string
}

var (
	defaultConfig = Config{
		DefaultTimezone: "UTC",
		Language:        "en",
	}
	configLock sync.RWMutex
)

type Option func(*Config)



// WithTimezone sets the default timezone for the operation.
//
// Example:
//
//	// Create a configuration with a specific timezone.
//	cfg := resolveConfig(WithTimezone("America/New_York"))
//	fmt.Println(cfg.DefaultTimezone) // Output: America/New_York
func WithTimezone(tz string) Option {
	return func(c *Config) {
		c.DefaultTimezone = tz
	}
}

// WithLanguage sets the default language for the operation.
//
// Example:
//
//	// Create a configuration with a specific language.
//	cfg := resolveConfig(WithLanguage("id"))
//	fmt.Println(cfg.Language) // Output: id
func WithLanguage(lang string) Option {
	return func(c *Config) {
		c.Language = lang
	}
}

// resolveConfig resolves the final configuration by applying a series of options
// to a copy of the package's default configuration.
//
// Example:
//
//	// Resolve a configuration with a custom timezone and language.
//	cfg := resolveConfig(WithTimezone("Europe/Berlin"), WithLanguage("de"))
//	fmt.Println(cfg.DefaultTimezone) // Output: Europe/Berlin
//	fmt.Println(cfg.Language)        // Output: de
func resolveConfig(opts ...Option) Config {
	configLock.RLock()
	cfg := defaultConfig
	configLock.RUnlock()
	
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

// SetDefaultConfig updates the package-level default configuration.
//
// Example:
//
//	// Update the default configuration with a custom timezone and language.
//	SetDefaultConfig("Europe/Berlin", "de")
func SetDefaultConfig(tz, lang string) {
	configLock.Lock()
	defer configLock.Unlock()
	defaultConfig.DefaultTimezone = tz
	defaultConfig.Language = lang
}

// SetDefaultTimezone updates only the default timezone.
//
// Example:
//
//	// Update the default timezone to "Europe/Berlin".
//	SetDefaultTimezone("Europe/Berlin")
func SetDefaultTimezone(tz string) {
	configLock.Lock()
	defer configLock.Unlock()
	defaultConfig.DefaultTimezone = tz
}

// SetDefaultLanguage updates only the default language.
//
// Example:
//
//	// Update the default language to "id".
//	SetDefaultLanguage("id")
func SetDefaultLanguage(lang string) {
	configLock.Lock()
	defer configLock.Unlock()
	defaultConfig.Language = lang
}

// UnixToTime converts a Unix timestamp (seconds since January 1, 1970 UTC) to a time.Time object.
//
// Example:
//
//	// Convert Unix timestamp 0 (epoch) to time.Time.
//	t := UnixToTime(0)
//	fmt.Println(t.UTC()) // Output: 1970-01-01 00:00:00 +0000 UTC

func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
