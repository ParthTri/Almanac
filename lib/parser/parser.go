package parser

import (
	"errors"
	"fmt"
)

type Event struct {
	Name				string
	Time				string
	Description string
	Tags				[]string
}

type Events struct {
	Date			string
	Events		[]Event
} 

type Almanac []Events;

// Check 10 bytes of data to match for date syntax YYYY-MM-DD
//	Hyphens = 4, 7
//	Rest numbers
func findDate(data []byte) (string, error) {
	var date string

	if data[4] == 45 && data[7] == 45 {
		for index, val := range data {
			if (val > 48 || val < 57) || index == 4 || index == 7{
				date += string(val)
			} else {
				return "", errors.New("Next set")
			}
		}
	} else {
		return "", errors.New("Next set")
	}

	return date, nil
}

