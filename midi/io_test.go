package midi

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gomidi/midi/reader"

	"osampler/test"
)

func noteOn(p *reader.Position, channel, key, velocity uint8) {
	//log.Printf("note ON : position: %v; channel: %v; key: %v; velocity: %v", p, channel, key, velocity)
}

func noteOff(p *reader.Position, channel, key, velocity uint8) {
	//log.Printf("note OFF: position: %v; channel: %v; key: %v; velocity: %v", p, channel, key, velocity)
}

func TestMidiFileIO(t *testing.T) {
	ass := require.New(t)
	// to disable logging, pass mid.NoLogger() as option
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.NoteOn(noteOn),
		reader.NoteOff(noteOff),
	)
	err := reader.ReadSMFFile(rd, test.ResolvePath("midi/jesu.mid"))
	ass.Nil(err)
}
