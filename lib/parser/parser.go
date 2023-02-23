package parser

import (
	"errors"
	"strconv"
	"strings"
)


type TimeRepeat	struct {
	Duration			float64
	Unit					string
}

type Event struct {
	Name					string
	Time					[]string
	TimeRepeat		*TimeRepeat
	Description		string
	Tags					[]string
}

type Day struct {
	Date			string
	Events		[]*Event
	Date				[]string
} 

type Almanac []*Day;

// List of ascii identifiers: [, ], *, +, \n, \t
var Identifiers [6]byte = [6]byte{91, 93, 42, 43, 9, 10}
var TimeUnits		[3]byte = [3]byte{72, 77, 83}

func checkIdentifier(target byte) bool {
	for _, id := range Identifiers {
		if target == id {
			return true
		}	
	}
	return false
}

func checkTimeUnit(target byte) bool {
	for _, id := range TimeUnits {
		if target == id || target == id+32 {
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
		for _, val := range data {
			if (val >= 48 && val <= 57) || val == 45{
				date += string(val)
			} else if val == 32 || val == 42 {
				break
			}
		}
	} else {
		return errors.New("Next set")
	}

	day.Date = append(day.Date, date)
	return nil
}

// Check at most 13 bytes to see if a time stamp is included, this includes a range
// If we exlude whitespace i.e. 09:00-09:10, the hyphen should be at 5
// Colons should be at 2 and 8
// NOTE: This function will only work if the time values are passed in proper form
func (event *Event)setTime(data []byte) (error) {
	event.TimeRepeat = &TimeRepeat{}	

	// Removing all whitespace
	trimmed := []byte{}

	if data[0] == 9 {
		trimmed = data[1:]
	}

	time := []byte{}
	for i := 0; i < len(trimmed); i++ {
		if trimmed[i]	== 91 {
			subset := trimmed[i+1:]
			for j := 0; j < len(subset); j++ {
				if subset[j] == 93 {
					break
				} 
				time = append(time, subset[j])
			}
		} 
	}

	if len(time) != 0 {
		if time[5] == 45 && time[2] == 58 && time[8] == 58 { 
			event.Time = append(event.Time, string(time[:5]))
			event.Time = append(event.Time, string(time[6:11]))

			var durStart, durEnd int
			subset := time[12:]

			for i := 0; i < len(subset); i++ {
				if subset[i] >= 48 && subset[i] <= 57 {
					durStart = i
				} else if checkTimeUnit(subset[i]) {
					durEnd = i
				}
			}

			if durEnd != 0 {
				durr, err := strconv.ParseFloat(string(subset[durStart-2:durEnd]), 10)
				if err != nil {
					return err
				}
				event.TimeRepeat.Duration = durr
				event.TimeRepeat.Unit = string(subset[durEnd])
			}
		}
	} else if len(trimmed) > 0 {
		if trimmed[5] == 45 && trimmed[2] == 58 && trimmed[8] == 58 { 
			event.Time = append(event.Time, string(trimmed[:5]))
			event.Time = append(event.Time, string(trimmed[6:11]))
		}
	}

	if event.Time == nil {
		return errors.New("Next Set")
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
		} else if ( val == 32 || val == 10 || index+1 == len(data) ) && len(tmp) != 0 {
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

func ParseFile(data []byte) *Almanac {
	Almanac := &Almanac{}

	dataString := string(data)
	Lines := strings.Split(dataString, "\n")

	for index := 0; index < len(Lines); index++ {
		line := Lines[index]
		day := &Day{}
	
		byteLine := []byte(line)
		if len(byteLine) == 0 {
			continue	
		}
		err := day.setDate(byteLine)

		if err == nil {
			var x []string
			var subset []string

			if len(byteLine) > 10 {
				x = append(x, line[10:])
				subset = append(x, Lines[index+1])
			} else {
				subset = Lines[index+1:]
			}

			for i := 0; i < len(subset); i++ {
				byteLine := []byte(subset[i])
				if len(byteLine) == 0 {
					break
				}

				x := &Day{}
				if err := x.setDate(byteLine); err == nil { 
					break 
				}

				event := &Event{}
				err = event.setTime(byteLine) 	

				if err == nil && len(byteLine) > 11 {
					curr := string(subset[i])

					currSplit := strings.Split(curr, " ")
					if curr[0] == 32 {
						currSplit = currSplit[4:]
					} else {
						currSplit = currSplit[3:]
					}

					curr = strings.Join(currSplit, " ")
					byteCurr := []byte(curr)

					event.setEventName(byteCurr)

					event.setTags(byteCurr)	

					if next := subset[i+1]; strings.Contains(next, "*") {
						event.setDescription([]byte(next))
						i++
					} 
				}

				if event.Name != "" {
					day.Events = append(day.Events, event)
				}
			}
		} 

		if len(day.Date) != 0 {
			*Almanac = append(*Almanac, day)
		}
	}

	return Almanac
}
