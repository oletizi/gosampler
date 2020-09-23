package osampler

import (
	"log"

	"gitlab.com/gomidi/midi/reader"
)

type printer struct{}

func (pr printer) noteOn(p *reader.Position, channel, key, vel uint8) {
	log.Printf("note ON: pos: %d, chan: %d, key: %v, vel: %d ", p.AbsoluteTicks, channel, key, vel)
}

func LoadFile(filename string) error {
	log.Printf("LoadFile: %v", filename)

	var p printer

	rd := reader.New(reader.NoLogger(),
		reader.NoteOn(p.noteOn),
	)

	return reader.ReadSMFFile(rd, filename)
}
