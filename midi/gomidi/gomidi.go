package gomidi

import (
	"log"

	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"

	instrument2 "osampler/instrument"
	"osampler/midi"
)

// abstraction for all callbacks required to accept a midi stream with gomidi
type MidiSink interface {
	Header(h smf.Header)
	NoteOn(p *reader.Position, channel, key, velocity uint8)
	NoteOff(p *reader.Position, channel, key, velocity uint8)
	Tempo(p reader.Position, bpm float64)
}

// adaptor between gomidi API and osampler Instrument API
// implements MidiSink interface and calls appropriate Instrument functions in gomidi callbacks
type InstrumentAdaptor struct {
	context    *context
	instrument instrument2.Instrument
}

// Mutable midi context
type context struct {
	*midi.BaseContext
	tempo float64
}

func (c *context) Tempo() float64 {
	return c.tempo
}

func NewInstrumentAdaptor(i instrument2.Instrument) *InstrumentAdaptor {
	c := &context{}
	return &InstrumentAdaptor{c, i}
}

func (adaptor *InstrumentAdaptor) Header(h smf.Header) {
	log.Printf("Header: %v", h.String())
}

func (adaptor *InstrumentAdaptor) NoteOn(r *reader.Position, channel, key, velocity uint8) {
	note, err := midi.NewNote(int(key))
	if err != nil {
		log.Printf("NoteOn: Error creating midi note: %v", err)
	} else {
		adaptor.instrument.NoteOn(adaptor.context, note, int(channel), int(velocity))
	}
}

// XXX: more elegant with function pointer?
func (adaptor *InstrumentAdaptor) NoteOff(p *reader.Position, channel, key, velocity uint8) {
	note, err := midi.NewNote(int(key))
	if err != nil {
		log.Printf("NoteOff: Error creating midi note: %v", err)
	} else {
		adaptor.instrument.NoteOff(adaptor.context, note, int(channel), int(velocity))
	}
}

func (adaptor *InstrumentAdaptor) Tempo(p reader.Position, bpm float64) {
	adaptor.context.tempo = bpm
}
