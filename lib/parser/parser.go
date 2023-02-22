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
	Events		[]*Event
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
func (day *Day)setDate(data []byte) (error) {
	var date string

	if data[4] == 45 && data[7] == 45 {
		for index, val := range data {
			if (val > 48 || val < 57) || index == 4 || index == 7{
				date += string(val)
			} else {
				return errors.New("Next set")
			}
		}
	} else {
		return errors.New("Next set")
	}

	day.Date = date
	return nil
}

// Check at most 13 bytes to see if a time stamp is included, this includes a range
// If we exlude whitespace i.e. 09:00-09:10, the hyphen should be at 5
// Colons should be at 2 and 8
// NOTE: This function will only work if the time values are passed in proper form
func (event *Event)setTime(data []byte) (error) {
	
	// Removing all whitespace
	trimmed := []byte{}
	for _, val := range data {
		if val != 9 && val != 32{
			trimmed = append(trimmed, val)
		}	else if val == 10 {
			return errors.New("Next Set")
		} else {
			continue
		}
	}
	
	if len(trimmed) != 11 {
		return errors.New("Next Set")
	}
	
	if trimmed[5] == 45 && trimmed[2] == 58 && trimmed[8] == 58 { 
		event.Time = append(event.Time, string(trimmed[:5]))
		event.Time = append(event.Time, string(trimmed[6:]))
	} else {
		return errors.New("Next set")
	}

	return nil
}

// Return the detected event name of a given set of bytes after finding the time
// The function keeps considers parts of an event name until it reaches an identifier
func (event *Event)setEventName(data []byte) (error) {
	name := []byte{}

	if data[0] == 32 {
		data = data[1:]
	}

	for _, val := range data {
		if !checkIdentifier(val) {
			name = append(name, val)
		} else {
			break
		}
	}

	if len(name) == 0 {
		return errors.New("No event name found")
	}

	if name[len(name)-1] == 32 {
		name = name[:len(name)-1]
	}

	event.Name = string(name)	
	return nil
}

// Find all the tags associated with a given date and set them as Event.Tags
func (event *Event)setTags(data []byte) (error) {
	tmp := []byte{}

	for index, val := range data {
		if val == 43 {
			compact := data[index:]
			for x := 0; x < len(compact); x++ {
				if compact[x] != 32 && compact[x] != 10{
					tmp = append(tmp, compact[x])
				} else {
					break
				}
			}
		} else if val == 32 && len(tmp) != 0 {
			event.Tags = append(event.Tags, string(tmp))	
			tmp = []byte{}
		} 	
	}
	
	return nil
}

// Find the description and set it as Event.Description
func (event *Event)setDescription(data []byte) (error)  {
	for index, val := range data {
		if val == 42 {
			compact := data[index+2:]
			for i := 0; i < len(compact); i++ {
				if compact[i] != 10 {
					event.Description += string(compact[i])
				} else {
					break
				}
			}
		} else if val == 10 && event.Description != "" {
			break
		}
	}

	return nil	
}


