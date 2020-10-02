package sin

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
	"osampler/audio/goaudio"
	"osampler/test"
	"osampler/transform/gain"
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

	sinTransform := New(buffer, float64(frequency), float64(phase))
	ass.NotNil(sinTransform)

	gainFactor := 10000.0
	gainTransform := gain.New(buffer, gainFactor)

	outfile, err := test.TempFile("test-*.aif")
	ass.Nil(err)
	ass.NotNil(outfile)
	filename := outfile.Name()

	out := goaudio.NewAiffSink(context, outfile)
	iterations := 100

	for i := 0; i < iterations; i++ {
		sinTransform.CalculateBuffer()
		gainTransform.CalculateBuffer()
		err := out.Write(buffer)
		ass.Nil(err)
	}
	err = out.Close()
	ass.Nil(err)

	err = outfile.Close()
	ass.Nil(err)

	t.Logf("Wrote file://%v", filename)
}
