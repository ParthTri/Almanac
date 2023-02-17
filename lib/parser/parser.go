package parser

import (
	"errors"
	"fmt"
)

type Event struct {
	Name				string
	Time				string
	Description string
	Tags				[]string
}

type Events struct {
	Date			string
	Events		[]Event
} 

type Almanac []Events;

