package sin

import (
	"testing"

	"github.com/go-audio/aiff"
	"github.com/stretchr/testify/require"

	"osampler/audio"
	"osampler/test"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)
	sampleRate := 44100
	bitDepth := 16
	bufferSize := 512
	channelCount := 1
	context := audio.NewContext(sampleRate, bitDepth, channelCount)
	buffer := audio.NewBuffer(context, bufferSize)
	frequency := 440
	phase := 0

	transform := New(buffer, float64(frequency), float64(phase))
	ass.NotNil(transform)

	outfile, err := test.TempFile("test-*.aif")
	ass.Nil(err)
	ass.NotNil(outfile)
	filename := outfile.Name()

	out := aiff.NewEncoder(outfile, sampleRate, bitDepth, channelCount)

	iterations := 100

	for i := 0; i < iterations; i++ {
		transform.CalculateBuffer()

	}

	t.Logf("Wrote file://%v", filename)
}
