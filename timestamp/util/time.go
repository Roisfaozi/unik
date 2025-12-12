package util

import (
	"sync"
	"time"
)

var (
	locationCache = make(map[string]*time.Location)
	mutex         sync.RWMutex
)

// LoadLocation is a cached wrapper around time.LoadLocation
func LoadLocation(name string) (*time.Location, error) {
	mutex.RLock()
	if loc, ok := locationCache[name]; ok {
		mutex.RUnlock()
		return loc, nil
	}
	mutex.RUnlock()

	mutex.Lock()
	defer mutex.Unlock()

	// Double check
	if loc, ok := locationCache[name]; ok {
		return loc, nil
	}

	loc, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}

	locationCache[name] = loc
	return loc, nil
}

// Normalize ensures the time is in the correct location based on timezone name.
// If timezone is empty, it returns the time as is.
// If timezone is invalid, it defaults to UTC (safety fallback).
func Normalize(t time.Time, timezone string) time.Time {
	if timezone == "" {
		return t
	}
	loc, err := LoadLocation(timezone)
	if err != nil {
		return t.UTC()
	}
	return t.In(loc)
}
