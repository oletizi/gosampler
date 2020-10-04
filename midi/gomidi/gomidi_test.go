package gomidi

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/gomidi/midi/reader"

	"osampler/midi"
)

type noteArgs struct {
	context  midi.Context
	channel  int
	note     midi.Note
	velocity int
}

type mockInstrument struct {
	noteOnCalls  []*noteArgs
	noteOffCalls []*noteArgs
}

func (m *mockInstrument) NoteOn(c midi.Context, note midi.Note, channel, velocity int) {
	m.noteOnCalls = append(m.noteOnCalls, &noteArgs{c, channel, note, velocity})
}

func (m *mockInstrument) NoteOff(c midi.Context, note midi.Note, channel, velocity int) {
	m.noteOffCalls = append(m.noteOffCalls, &noteArgs{c, channel, note, velocity})
}

func TestNewInstrumentAdaptor(t *testing.T) {
	ass := require.New(t)
	i := &mockInstrument{}
	adaptor := NewInstrumentAdaptor(i)
	ass.NotNil(adaptor)

	// set the tempo
	tempo := 99.0
	var position = reader.Position{}
	adaptor.Tempo(position, tempo)

	channel := 0
	key := 60
	velocity := 100

	// test note on call
	ass.Equal(0, len(i.noteOnCalls))
	adaptor.NoteOn(&position, uint8(channel), uint8(key), uint8(velocity))
	ass.Equal(1, len(i.noteOnCalls))

	noteArgs := i.noteOnCalls[0]

	testNoteArgs(ass, tempo, noteArgs, key, channel, velocity)

	// test note off call
	ass.Equal(0, len(i.noteOffCalls))
	adaptor.NoteOff(&position, uint8(channel), uint8(key), uint8(velocity))
	ass.Equal(1, len(i.noteOffCalls))

	noteArgs = i.noteOffCalls[0]
	testNoteArgs(ass, tempo, noteArgs, key, channel, velocity)

}

func testNoteArgs(ass *require.Assertions, tempo float64, noteArgs *noteArgs, key int, channel int, velocity int) {
	ass.Equal(tempo, noteArgs.context.Tempo())
	ass.Equal(key, noteArgs.note.Value())
	ass.Equal(channel, noteArgs.channel)
	ass.Equal(velocity, noteArgs.velocity)
}
