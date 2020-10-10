package gomidi

import (
	"log"

	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/smf"

	"osampler/instrument"
	"osampler/midi"
)

// abstraction for all callbacks required to accept a midi stream with gomidi
//type MidiSink interface {
//	header(h smf.header)
//	noteOn(p *reader.Position, channel, key, velocity uint8)
//	noteOff(p *reader.Position, channel, key, velocity uint8)
//	tempo(p reader.Position, bpm float64)
//}

// adaptor between gomidi API and osampler Instrument API
// implements MidiSink interface and calls appropriate Instrument functions in gomidi callbacks
type InstrumentAdaptor struct {
	in         *reader.Reader
	context    *context
	channel    int
	instrument instrument.Instrument
}

// Mutable midi context
type context struct {
	*midi.BaseContext
	tempo float64
}

func (c *context) Tempo() float64 {
	return c.tempo
}

func NewInstrumentAdapter(channel int, i instrument.Instrument) *InstrumentAdaptor {
	c := &context{}
	adapter := &InstrumentAdaptor{nil, c, channel, i}
	var r = reader.New(
		reader.NoLogger(),
		reader.SMFHeader(adapter.header),
		reader.NoteOn(adapter.noteOn),
		reader.NoteOff(adapter.noteOff),
		reader.TempoBPM(adapter.tempo),
	)
	adapter.in = r
	return adapter
}

func (adaptor *InstrumentAdaptor) header(h smf.Header) {
	log.Printf("header: %v", h.String())
}

func (adaptor *InstrumentAdaptor) noteOn(r *reader.Position, channel, key, velocity uint8) {
	if adaptor.channel == int(channel) {
		note, err := midi.NewNote(int(key))
		if err != nil {
			log.Printf("noteOn: Error creating midi note: %v", err)
		} else {
			adaptor.instrument.NoteOn(adaptor.context, note, int(channel), int(velocity))
		}
	}
}

// XXX: more elegant with function pointer?
func (adaptor *InstrumentAdaptor) noteOff(p *reader.Position, channel, key, velocity uint8) {
	note, err := midi.NewNote(int(key))
	if err != nil {
		log.Printf("noteOff: Error creating midi note: %v", err)
	} else {
		adaptor.instrument.NoteOff(adaptor.context, note, int(channel), int(velocity))
	}
}

func (adaptor *InstrumentAdaptor) tempo(p reader.Position, bpm float64) {
	adaptor.context.tempo = bpm
}

func (adaptor *InstrumentAdaptor) ConsumeFile(filepath string) {
	err := reader.ReadSMFFile(adaptor.in, filepath)
	if err != nil {
		log.Printf("Error reaading midi file: %v", err)
	}
}
