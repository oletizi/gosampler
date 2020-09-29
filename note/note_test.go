package note

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	ass := require.New(t)
	var value int
	var name string
	var err error

	for i := 0; i < 128; i++ {
		note, err := New(i)
		ass.NotNil(note)
		ass.Nil(err)
		ass.Equal(i, note.Value())
	}

	// below lower bound
	value = -1
	note, err := New(value)
	ass.Nil(note)
	ass.NotNil(err)

	// at lower bound
	value = 0
	note, err = New(value)
	ass.NotNil(note)
	ass.Nil(err)
	ass.Equal(value, note.Value())

	// at upper bound
	value = 127
	note, err = New(value)
	ass.NotNil(note)
	ass.Nil(err)
	ass.Equal("G9", note.Name())
	ass.Equal(value, note.Value())

	// above upper bound
	value = 128
	note, err = New(value)
	ass.Nil(note)
	ass.NotNil(err)

	// Test A0
	value = 21
	name = "A0"
	note, err = New(value)
	ass.Nil(err)
	ass.NotNil(note)
	ass.Equal(value, note.Value())
	ass.Equal(name, note.Name())

	// Test middle-C
	value = 60
	name = "C4"
	note, err = New(value)
	ass.NotNil(note)
	ass.Nil(err)
	ass.Equal(value, note.Value())
	ass.Equal(name, note.Name())

	// Test A440
	value = 69
	name = "A4"
	note, err = New(value)
	ass.Nil(err)
	ass.NotNil(note)
	ass.Equal(value, note.Value())
	ass.Equal(name, note.Name())
}

func TestEquals(t *testing.T) {
	ass := require.New(t)
	a, _ := New(1)
	b, _ := New(1)
	ass.Equal(a, b)

	b, _ = New(2)
	ass.NotEqual(a, b)
}
