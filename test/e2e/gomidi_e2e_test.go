package e2e

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio/sample"
	"osampler/midi/gomidi"
	"osampler/sfz"
	"osampler/test"
)

func TestE2E(t *testing.T) {
	ass := require.New(t)
	midifile := test.ResolvePath("midi/jesu.mid")
	sfzFile := test.ResolvePath("sfz/piano/piano.sfz")
	config, err := sfz.New(sfzFile)
	ass.Nil(err)
	instrument := sample.NewInstrument(config)
	channel := 0
	adaptor := gomidi.NewInstrumentAdapter(channel, instrument)
	adaptor.ConsumeFile(midifile)
	ass.NotNil(adaptor)
}
