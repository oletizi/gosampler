package sampler

import (
	"osampler"
)

type midiSampler struct {
}

func New(config osampler.Config) osampler.Instrument {
	return &midiSampler{}
}

func (m *midiSampler) NoteOn(note *osampler.Note) {
	panic("implement me")
}

func (m *midiSampler) NoteOff(note *osampler.Note) {
	panic("implement me")
}
