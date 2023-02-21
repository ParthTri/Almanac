package parser

import (
	"errors"
)

type Event struct {
	Name					string
	Time					[]string
	Description		string
	Tags					[]string
}

type Day struct {
	Date			string
	Events		[]Event
} 

type Almanac []Day;

// List of ascii identifiers: [, ], *, +, \n, \t
var Identifiers [6]byte = [6]byte{91, 93, 42, 43, 9, 10}

func checkIdentifier(target byte) bool {
	for _, id := range Identifiers {
		if target == id {
			return true
		}	
	}
	return false
}

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

// Check at most 13 bytes to see if a time stamp is included, this includes a range
// If we exlude whitespace i.e. 09:00-09:10, the hyphen should be at 5
// Colons should be at 2 and 8
func findTime(data []byte) ([]string, error) {
	var time []string
	
	// Removing all whitespace
	trimmed := []byte{}
	for _, val := range data {
		if val != 9 && val != 32{
			trimmed = append(trimmed, val)
		}	else if val == 10 {
			return time, errors.New("Next Set")
		} else {
			continue
		}
	}
	
	if len(trimmed) != 11 {
		return time, errors.New("Next Set")
	}
	
	if trimmed[5] == 45 && trimmed[2] == 58 && trimmed[8] == 58 { 
		time = append(time, string(trimmed[:5]))		
		time = append(time, string(trimmed[6:]))		
	} else {
		return time, errors.New("Next set")
	}

	return time, nil
}

