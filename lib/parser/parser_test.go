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

func TestFindTime(t *testing.T) {
	want := []string{"09:00", "09:10"}
	data := []byte(TestData)
	
	var time []string
	var err error

	for index := 0; index < len(data) && index+13 <= len(data); index += 1{
		time, err = findTime(data[index:index+13])
		if err == nil {
			break
		}
	}

	if time[0] != want[0] && time[1] != want[1] { 
		t.Errorf("Wanted %v got %v", want, time)
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
