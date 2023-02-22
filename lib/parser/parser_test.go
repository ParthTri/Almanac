package parser

import (
	"testing"
)

const TestData = `2023-02-13
	09:00 - 09:10 Meditate +health
		* Meditate at beach
`

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

