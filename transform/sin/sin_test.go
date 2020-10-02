package sin

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
	"osampler/audio/goaudio"
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

	out := goaudio.NewAiffSink(context, outfile)
	iterations := 100

	for i := 0; i < iterations; i++ {
		transform.CalculateBuffer()
		b := transform.Buffer()
		// add gain manually for now
		for sample := 0; sample < b.Size(); sample++ {
			b.Data()[sample] = b.Data()[sample] * 1000
		}
		err := out.Write(b)
		ass.Nil(err)
	}
	err = out.Close()
	ass.Nil(err)

	err = outfile.Close()
	ass.Nil(err)

	t.Logf("Wrote file://%v", filename)
}
