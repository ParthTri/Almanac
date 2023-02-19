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

func TestFindDate(t *testing.T) {
	want := "2023-02-13"	
	data := []byte(TestData)
	
	var date string
	var err error
	for index := range data {
		if index + 10 > len(data) {
			break
		}
		date, err = findDate(data[index:index+10])
		if err == nil {
			break
		}
	}

	if date != want {
		t.Errorf("Wanted %v got %v", want, date)
	} 
}

func TestFindTime(t *testing.T) {
	want := []string{"09:00", "09:10"}
	data := []byte(TestData)
	
	var time []string
	var err error

	for index := range data {
		if index + 10 > len(data) {
			break
		}
		time, err = findTime(data[index:index+13])
		if err == nil {
			break
		}
	}

	if time[0] == want[0] && time[1] == want[1] { 
		t.Errorf("Wanted %v got %v", want, time)
	} 
}

func TestFindEventName(t *testing.T) {
	want := "Meditate"
	data := []byte(TestData)

	var eventName string
	for index := range data {
		if index + 10 > len(data) {
			break
		}
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
		t.Errorf("Event Name Not Found\n Wanted %v Got %v", want, eventName)
	}
}

