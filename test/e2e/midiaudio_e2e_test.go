package e2e

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"

	"osampler/test"
)

type sink struct {
}

func (s *sink) noteOn(p *reader.Position, channel, key, velocity uint8) {
	//log.Printf("note ON : position: %v; channel: %v; key: %v; velocity: %v", p, channel, key, velocity)
}

func (s *sink) noteOff(p *reader.Position, channel, key, velocity uint8) {
	//log.Printf("note OFF: position: %v; channel: %v; key: %v; velocity: %v", p, channel, key, velocity)
}

func (s *sink) header(h smf.Header) {
	log.Printf("header: %v", h.String())
}

func (s *sink) tempo(p reader.Position, bpm float64) {
	log.Printf("Tempo: position: %v; bpm: %v", p, bpm)
}

func TestMidiInToAudio(t *testing.T) {
	ass := require.New(t)
	// to disable logging, pass mid.NoLogger() as option
	s := &sink{}
	rd := reader.New(reader.NoLogger(),
		// set the functions for the messages you are interested in
		reader.SMFHeader(s.header),
		reader.NoteOn(s.noteOn),
		reader.NoteOff(s.noteOff),
		reader.TempoBPM(s.tempo),
	)

	err := reader.ReadSMFFile(rd, test.ResolvePath("midi/jesu.mid"))
	ass.Nil(err)
}
