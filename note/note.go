package note

import (
	"errors"
	"strconv"

	"osampler"
)

type note struct {
	value int
	name  string
}

func (n note) Value() int {
	return n.value
}

func (n note) Name() string {
	return n.name
}

func New(value int) (osampler.Note, error) {
	name, err := valueToName(value)
	if err != nil {
		return nil, err
	} else {
		return note{name: name, value: value}, nil
	}
}

func valueToName(value int) (string, error) {
	var err error
	var name string
	offset := 12
	if value < 0 || value > 127 {
		err = errors.New("Out of bounds: " + strconv.Itoa(value))
	} else if value < offset {
		name = ""
	} else {
		name = []string{"C", "C#/Db", "D", "D#/Eb", "E", "F", "F#/Gb", "G", "G#/Ab", "A", "A#/Bb", "B"}[(value-offset)%12] + strconv.Itoa((value-offset)/12)
	}
	return name, err
}
