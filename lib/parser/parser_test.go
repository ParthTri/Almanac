package parser

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"strings"
)

const TestData = `2023-02-13
	[09:00-09:10 1.5h] Meditate +health
		* Meditate at beach
`

func TestSetDate(t *testing.T) {
	want := Day{
		Date: []string{"2023-02-13"},
	}
	result := &Day{}
	
	data := []byte(TestData)
	
	var err error
	for index := range data {
		if index + 10 > len(data) {
			break
		}
		err = result.setDate(data[index:index+10])
		if err == nil {
			break
		}
	}

	if result.Date[0] != want.Date[0] {
		t.Errorf("Wanted %v got %v", want.Date, result.Date)
	} 
}

func TestSetTime(t *testing.T) {
	want := &Event{
		Time: []string{"09:00", "09:10"},
		TimeRepeat: &TimeRepeat{
			Duration: 1.5,
			Unit:			"h",
		},
	}
	result := &Event{}
	data := []byte(TestData)
	
	dataString := string(data)	
	dataSlice := strings.Split(dataString, "\n")

	for index := 0; index < len(dataSlice); index++{
		byteLine := []byte(dataSlice[index])
		err := result.setTime(byteLine)

		if err == nil {
			break
		} 	
	}

	if result.Time[0] != want.Time[0] && result.Time[1] != want.Time[1] { 
		t.Errorf("Wanted %v got %v", want.Time, result.Time)
	} 

	if result.TimeRepeat.Duration != want.TimeRepeat.Duration {
		t.Errorf("Wanted %v got %v", want.TimeRepeat.Duration, result.TimeRepeat.Duration)
	}


	if result.TimeRepeat.Unit != want.TimeRepeat.Unit {
		t.Errorf("Wanted %v got %v", want.TimeRepeat.Unit, result.TimeRepeat.Unit)
	}
}

func TestSetEventName(t *testing.T) {
	want := &Event{
		Name: "Meditate",
	}
	result := &Event{}
	data := []byte(TestData)

	for index := 0; index < len(data) && index + 13 <= len(data); index++ {
		err := result.setTime(data[index:index+13])
		if err == nil {
			err = result.setEventName(data[index+13:])
			if err != nil {
				t.Error(err)
			} else {
				break
			}
		}
	}

	if result.Name != want.Name {
		t.Errorf("Event Name Not Found\n Wanted %v Got %v\n %v \n %v", want.Name, result.Name, []byte(want.Name), []byte(result.Name))
	}
}

func TestSetTags(t *testing.T) {
	want := &Event{
		Tags: []string{"+health"},
	}
	result := &Event{}

	data := []byte(TestData)

	for index := 0; index < len(data) && index+13 <= len(data); index++ {
		err := result.setTime(data[index:index+13])
		if err == nil {
			result.setEventName(data[index+13:])
			err := result.setTags(data[index+13:])

			if err == nil {
				break
			}
		}
	}

	if result.Tags[0] != want.Tags[0] {
		t.Errorf("Tags Not Found.\n\t Wanted %v, Got %v", want.Tags[0], result.Tags[0])
	}
}

func TestSetDescription(t *testing.T) {
	want := &Event{
		Description: "Meditate at beach",
	}	
	result := &Event{}
	data := []byte(TestData)

	for index := 0; index < len(data) && index+13 <= len(data); index++ {
		err := result.setTime(data[index:index+13])	
		if err == nil {
			err = result.setDescription(data[index+13:])
			if err == nil {
				break	
			}
		}
	}
	if result.Description != want.Description {
		t.Errorf("Description Not Found.\n\t Wanted \"%v\", Got \"%v\"", want.Description, result.Description)
	}
}

func compareStructs(first Almanac, second Almanac) error {
	if len(first) != len(second) {
		return errors.New(fmt.Sprintf("Number of Dates do not match\n\t Wanted %v, Got %v", len(first), len(second)))
	}

	for i := 0; i < len(first); i++ {
		if first[i].Date[0] != second[i].Date[0] {
			return errors.New(fmt.Sprintf("Dates do not match\n\tWanted %v, Got %v", first[i].Date, second[i].Date))
		}

		if len(first[i].Events) != len(second[i].Events) {
			return errors.New(fmt.Sprintf("Number of events do not match\n\tWanted %v, Got %v", len(first[i].Events), len(second[i].Events)))
		}

		for j := 0; j < len(first[i].Events); j++ {
			if first[i].Events[j].Time[0] != second[i].Events[j].Time[0] || first[i].Events[j].Time[1] != second[i].Events[j].Time[1]{
				return errors.New(fmt.Sprintf("Event Times Do Not Match\n\tWanted %v, Got %v", first[i].Events[j].Time, second[i].Events[j].Time))
			}
			
			if first[i].Events[j].Name != second[i].Events[j].Name {
				return errors.New(fmt.Sprintf("Event names do not match\n\tWanted %v, Got %v", first[i].Events[j].Name, second[i].Events[j].Name))
			}

			if first[i].Events[j].Description != second[i].Events[j].Description {
				return errors.New(fmt.Sprintf("Event descriptions do not match\n\tWanted %v, Got %v", first[i].Events[j].Description, second[i].Events[j].Description))
			}

			for index, tag := range first[i].Events[j].Tags {
				if tag != second[i].Events[j].Tags[index] {
					return errors.New(fmt.Sprintf("Tags do not match\n\tWanted %v, Got %v", first[i].Events[j].Tags, second[i].Events[j].Tags))
				}
			}
		}
	}
	return nil
}

func TestParseFile(t *testing.T) {
	want := &Almanac{
		&Day{
			Date: []string{"2023-02-13"},
			Events: []*Event{
				&Event{
					Time: []string{"09:00", "09:10"},
					Name: "Meditate",
				},
				&Event{
					Time: []string{"17:00", "17:30"},
					Name: "Accounting Meeting",
					Tags: []string{"+work"},
					Description: "Talk about balance sheet",
				},
			},
		},
		&Day{
			Date: []string{"2023-02-14"},
			Events: []*Event{
				&Event{
					Time: []string{"14:00", "16:00"},
					Name: "Computer Science lecture",
					Tags: []string{"+school"},
					Description: "Make sure to create sketch notes",
				},
			},
		},
	}

	data, err := os.ReadFile("../../example.almanac")
	if err != nil {
		t.Error(err)
	}
	result := ParseFile(data)

	if err := compareStructs(*want, *result); err != nil{
		t.Error(err.Error())
	}
}
