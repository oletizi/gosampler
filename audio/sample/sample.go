package sample

import (
	"log"

	"osampler/instrument"
	"osampler/midi"
)

type Sample interface{}

type s struct {
	filepath string
}

func New(filepath string) Sample {
	return &s{filepath}
}

// Sample-based Instrument
type inst struct {
	config instrument.Config
}

func NewInstrument(config instrument.Config) instrument.Instrument {
	return &inst{config}
}

func (i *inst) NoteOn(c midi.Context, note midi.Note, channel, velocity int) {
	log.Printf("NOTE ON: note: %v, channel: %v, velocity: %v, I should play these files: %v: ",
		note, channel, velocity, i.config.FilesFor(note))
}

func (i *inst) NoteOff(c midi.Context, note midi.Note, channel, velocity int) {
	log.Printf("NOTE OFF: I should stop playing the audio files!")
}
