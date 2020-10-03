package aiff

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
	"osampler/test"
)

func TestBasics(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	ass := require.New(t)

	infile, err := os.Open(test.ResolvePath("audio/sound.aif"))
	ass.Nil(err)
	ass.NotNil(infile)

	buffer := audio.NewBuffer(512)
	outfile, err := test.TempFile("out-*.aif")

	input, c, err := NewAiffInput(buffer, infile)
	ass.Nil(err)
	ass.NotNil(input)

	bitDepth := c.BitDepth()
	ass.NotNil(bitDepth)
	ass.Equal(16, bitDepth.Depth())

	output := NewAiffOutput(c, buffer, outfile)
	ass.NotNil(output)
	bytesRead := input.CalculateBuffer()
	for bytesRead > 0 {
		bytesWritten := output.CalculateBuffer()
		ass.Equal(bytesRead, bytesWritten)
		bytesRead = input.CalculateBuffer()
	}
	err = output.Close()
	ass.Nil(err)

	err = outfile.Close()
	ass.Nil(err)

	// test that the aiff reader set the buffer size to the amount of bytes read
	ass.Equal(0, buffer.Size())

	t.Logf("Wrote file://%v", outfile.Name())
}
