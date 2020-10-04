package mixer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
	"osampler/test"
	"osampler/transform"
	"osampler/transform/aiff"
	"osampler/transform/multi"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)
	a := audio.NewBuffer(2)
	b := audio.NewBuffer(2)
	output := audio.NewBuffer(2)
	a.Data()[0] = .5
	b.Data()[0] = -.5

	a.Data()[0] = .1
	b.Data()[0] = .1

	inputs := []audio.Buffer{a, b}
	mix := New(inputs, output)
	ass.NotNil(mix)

	mix.CalculateBuffer()

	ass.Equal(a.Data()[0]+b.Data()[0], output.Data()[0])
	ass.Equal(a.Data()[1]+b.Data()[1], output.Data()[1])
}

func TestPassthrough(t *testing.T) {
	ass := require.New(t)
	bufferSize := 512

	// Setup inputs
	hh := open(ass, "audio/hh.aif")
	hhBuffer := audio.NewBuffer(bufferSize)
	hhIn, c, err := aiff.NewAiffInput(hhBuffer, hh)
	ass.Nil(err)

	// Setup outputs
	mixOutfile, err := test.TempFile("mixed-*.aif")
	copyOutfile, err := test.TempFile("copy-*.aif")

	mixOutBuffer := audio.NewBuffer(bufferSize)
	mixOut := aiff.NewAiffOutput(c, mixOutBuffer, mixOutfile)
	// wire input buffer directly the copy output buffer
	copyOut := aiff.NewAiffOutput(c, hhBuffer, copyOutfile)

	// wire up a mixer between input buffers and output buffer
	mixer := New([]audio.Buffer{hhBuffer}, mixOutBuffer)

	// pump buffers
	bytes := hhIn.CalculateBuffer()
	mixer.CalculateBuffer()
	mixOut.CalculateBuffer()
	copyOut.CalculateBuffer()

	for bytes > 0 {
		bytes = hhIn.CalculateBuffer()
		mixer.CalculateBuffer()
		mixOut.CalculateBuffer()
		copyOut.CalculateBuffer()
	}
	err = copyOut.Close()
	ass.Nil(err)

	err = mixOut.Close()
	ass.Nil(err)

	err = copyOutfile.Close()
	ass.Nil(err)

	err = mixOutfile.Close()
	ass.Nil(err)

	t.Logf("Copyied to file://%v", copyOutfile.Name())
	t.Logf("Mixed to file://%v", mixOutfile.Name())

}

func TestMixAudioFiles(t *testing.T) {
	ass := require.New(t)
	hh := open(ass, "audio/hh.aif")
	snare := open(ass, "audio/snare.aif")
	kick := open(ass, "audio/kick.aif")

	outfile, err := test.TempFile("mixed-*.aif")
	ass.Nil(err)
	ass.NotNil(outfile)

	bufferSize := 512
	// TODO: Handle transcoding on the fly?
	context := audio.NewCDContext()

	hhBuffer := audio.NewBuffer(bufferSize)
	hhIn, c, err := aiff.NewAiffInput(hhBuffer, hh)
	ass.Nil(err)
	ass.Equal(context, c)
	ass.NotNil(hhIn)

	snareBuffer := audio.NewBuffer(bufferSize)
	snareIn, c, err := aiff.NewAiffInput(snareBuffer, snare)
	ass.Nil(err)
	ass.Equal(context, c)
	ass.NotNil(snareIn)

	kickBuffer := audio.NewBuffer(bufferSize)
	kickIn, c, err := aiff.NewAiffInput(kickBuffer, kick)
	ass.Nil(err)
	ass.Equal(context, c)
	ass.NotNil(kickIn)

	outBuffer := audio.NewBuffer(bufferSize)
	//outGain := gain.New(outBuffer, .7)
	out := aiff.NewAiffOutput(context, outBuffer, outfile)

	mixer := New([]audio.Buffer{
		hhBuffer,
		snareBuffer,
		kickBuffer,
	}, outBuffer)

	inputs := multi.New([]transform.Transform{
		hhIn,
		snareIn,
		kickIn,
		mixer,
		out})
	bytes := inputs.CalculateBuffer()
	totalBytes := bytes
	for bytes > 0 {
		bytes = inputs.CalculateBuffer()
		totalBytes += bytes
		//fmt.Printf("bytes: %v, totalBytes: %v\n", bytes, totalBytes)
	}
	out.Close()
	err = outfile.Close()
	ass.Nil(err)

	t.Logf("Total bytes: %v", totalBytes)
	t.Logf("Wrote file://%v", outfile.Name())
}

func open(ass *require.Assertions, subpath string) *os.File {
	rv, err := os.Open(test.ResolvePath(subpath))
	ass.Nil(err)
	ass.NotNil(rv)
	return rv
}
