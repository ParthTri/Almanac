package parser

import (
	"testing"
)

const TestData = `2023-02-13
	09:00 - 09:10 Meditate +health
		* Meditate at beach
`


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

