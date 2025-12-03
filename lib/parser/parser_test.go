package parser

import (
	"slices"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	expected := Event{
		Date:        "2025-12-03",
		Name:        "Test Event",
		Time:        []string{"14:00"},
		Description: "This is a test event",
		Tags:        []string{"work"},
	}

	//

	input := "2025-12-03 14:00 Test Event +work * This is a test event\n"
	reader := strings.NewReader(input)

	parser := NewParser(reader)

	output, err := parser.Parse()

	if err != nil {
		t.Errorf("An error occured %v", err.Error())
		return
	}

	if output.Date != expected.Date {
		t.Errorf("Dates do not match, got '%v' expected '%v'", output.Date, expected.Date)
	}

	if !slices.Equal(output.Time, expected.Time) {
		t.Errorf("Times do not match, got '%v' expected '%v'", output.Time, expected.Time)
	}

	if output.Name != expected.Name {
		t.Errorf("Names do not match, got '%v' expected '%v'", output.Name, expected.Name)
	}

	if output.Description != expected.Description {
		t.Errorf("Descriptions do not match, got '%v' expected '%v'", output.Description, output.Description)
	}

	if !slices.Equal(output.Tags, expected.Tags) {
		t.Errorf("Tags do not match, got '%v' expected '%v'", output.Tags, expected.Tags)
	}

	t.Logf("%v", output)
}
