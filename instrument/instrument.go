package instrument

import "osampler/midi"

// Abstraction for an instrument that receives midi messages
type Instrument interface {
	NoteOn(c midi.Context, note midi.Note, channel, velocity int)
	NoteOff(c midi.Context, note midi.Note, channel, velocity int)
}
