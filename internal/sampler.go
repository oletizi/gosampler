package osampler

import (
	"gitlab.com/gomidi/midi/reader"
	"log"
)

type printer struct{}

func (pr printer) noteOn(p *reader.Position, channel, key, vel uint8) {

}

func LoadFile(filename string) {
	//var p printer
	//
	//rd := reader.New(reader.NoLogger(),
	//	reader.NoteOn(p.noteOn),
	//	)
	//err = rd.ReadSMFFile(rd, f)
	log.Printf("LoadFile: %v", filename)
}
