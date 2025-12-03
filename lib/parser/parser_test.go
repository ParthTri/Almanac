package parser

import (
	"slices"
	"strings"
	"testing"
)

func TestParserSingleEntry(t *testing.T) {
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

func TestParseMultipleEntries(t *testing.T) {
	expected := []Event{
		{
			Date:        "2025-12-03",
			Name:        "Morning Meeting",
			Time:        []string{"09:30"},
			Description: "Daily team sync",
			Tags:        []string{"work"},
		},
		{
			Date:        "2025-12-03",
			Name:        "Gym Session",
			Time:        []string{"18:00"},
			Description: "Evening workout",
			Tags:        []string{"health", "exercise"},
		},
	}

	input :=
		"2025-12-03 09:30 Morning Meeting +work * Daily team sync\n" +
			"2025-12-03 18:00 Gym Session +health +exercise * Evening workout\n"

	reader := strings.NewReader(input)
	parser := NewParser(reader)

	output, err := parser.ParseAll() // assuming multi-entry method; rename if needed
	if err != nil {
		t.Fatalf("ParseAll returned error: %v", err)
	}

	if len(output) != len(expected) {
		t.Fatalf("Expected %d events, got %d", len(expected), len(output))
	}

	for i := range expected {
		ev := output[i]
		exp := expected[i]

		if ev.Date != exp.Date {
			t.Errorf("Event %d: Date mismatch, got '%v' expected '%v'", i, ev.Date, exp.Date)
		}
		if !slices.Equal(ev.Time, exp.Time) {
			t.Errorf("Event %d: Time mismatch, got '%v' expected '%v'", i, ev.Time, exp.Time)
		}
		if ev.Name != exp.Name {
			t.Errorf("Event %d: Name mismatch, got '%v' expected '%v'", i, ev.Name, exp.Name)
		}
		if ev.Description != exp.Description {
			t.Errorf("Event %d: Description mismatch, got '%v' expected '%v'", i, ev.Description, exp.Description)
		}
		if !slices.Equal(ev.Tags, exp.Tags) {
			t.Errorf("Event %d: Tags mismatch, got '%v' expected '%v'", i, ev.Tags, exp.Tags)
		}

		t.Log(ev)
	}
}
