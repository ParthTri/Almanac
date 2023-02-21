package parser

import (
	"testing"
)

const TestData = `2023-02-13
	09:00 - 09:10 Meditate +health
		* Meditate at beach
`

var Outcome Almanac = Almanac{
	Day{
		Date: "2023-02-13",
		Events: []Event{
			Event{
				Name: "Meditate",
				Time: []string{"09:00", "09:10"},
				Description: "Meditate at beach",
				Tags: []string{"health"},
			},
		},
	},
}

func TestSetDate(t *testing.T) {
	want := Day{
		Date: "2023-02-13",
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

	if result.Date != want.Date {
		t.Errorf("Wanted %v got %v", want.Date, result.Date)
	} 
}

func TestSetTime(t *testing.T) {
	want := &Event{
				Time: []string{"09:00", "09:10"},
	}
	result := &Event{}
	data := []byte(TestData)
	
	for index := 0; index < len(data) && index+13 <= len(data); index += 1{
		err := result.setTime(data[index:index+13])
		if err == nil {
			break
		}
	}

	if result.Time[0] != want.Time[0] && result.Time[1] != want.Time[1] { 
		t.Errorf("Wanted %v got %v", want.Time, result.Time)
	} 
}

func TestFindEventName(t *testing.T) {
	want := "Meditate"
	data := []byte(TestData)

	var eventName string
	for index := 0; index < len(data) && index + 13 <= len(data); index++ {
		_, err := findTime(data[index:index+13])
		if err == nil {
			eventName, err = findEventName(data[index+13:])
			if err != nil {
				t.Error(err)
			} else {
				break
			}
		}
	}

	if eventName != want {
		t.Errorf("Event Name Not Found\n Wanted %v Got %v\n %v \n %v", want, eventName, []byte(want), []byte(eventName))
	}
}

// func TestFindTags(t *testing.T) {
