package main

import (
	"fmt"
	"time"
)

// ConvertDate : Converts given time to the given timezone
func ConvertDate(zone string, date time.Time) time.Time {
	location, err := time.LoadLocation(zone)
	if err != nil {
		fmt.Printf("Error converting date: %q\n", err)
	}
	return date.In(location)
}
