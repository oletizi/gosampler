package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
	"osampler/test"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)
	// TODO: Figure out a nice way to establish the context from an audio file
	// TODO: Figure out a nice way to transparently transform the file to fit the current audio context
	c := audio.NewContext(44100, 16, 2)
	buffer := audio.NewBuffer(c, 512)

	infile, err := os.Open(test.ResolvePath("audio/sound.aif"))
	ass.Nil(err)
	ass.NotNil(infile)

	outfile, err := test.TempFile("out-*.aif")

	input := NewAiffInput(buffer, infile)
	ass.NotNil(input)

	output := NewAiffOutput(buffer, outfile)
	ass.NotNil(output)

	for input.CalculateBuffer() > 0 {
		output.CalculateBuffer()
	}
	err = output.Close()
	ass.Nil(err)

	err = outfile.Close()
	ass.Nil(err)

	t.Logf("Write file://%v", outfile.Name())
}
