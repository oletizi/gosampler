package goaudio

import (
	"io"
	"os"
	"testing"

	"github.com/go-audio/aiff"

	"github.com/stretchr/testify/require"

	goaudio2 "github.com/go-audio/audio"

	"osampler/audio"
	"osampler/test"
)

func TestAiffSink(t *testing.T) {

	ass := require.New(t)
	outfile, err := test.TempFile("sink-test-*.aif")
	path := outfile.Name()
	ass.Nil(err)
	ass.NotNil(outfile)

	context := audio.NewContext(44100, 16, 1)

	sink := NewAiffSink(context, outfile)
	ass.NotNil(sink)

	infile, err := os.Open(test.ResolvePath("audio/sound.aif"))
	ass.Nil(err)
	ass.NotNil(infile)

	in := aiff.NewDecoder(infile)
	in.ReadInfo()

	bufferSize := 512

	inBuffer := &goaudio2.IntBuffer{
		Format:         in.Format(),
		Data:           make([]int, bufferSize),
		SourceBitDepth: int(in.BitDepth),
	}

	outBuffer := audio.NewBuffer(context, bufferSize)

	bytesRead := 0
	var doneReading bool
	for err == nil {
		bytesRead, err = in.PCMBuffer(inBuffer)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			ass.FailNow("Error reading from input: %v", err)
		}
		if bytesRead != len(inBuffer.Data) {
			// we didn't read a full buffer, so trim the buffer
			inBuffer.Data = inBuffer.Data[:bytesRead]
			doneReading = true
		}

		for i := 0; i < len(inBuffer.Data); i++ {
			outBuffer.Data()[i] = float64(inBuffer.Data[i])
		}

		if err = sink.Write(outBuffer); err != nil {
			ass.FailNow("Error writing to output: %v", err)
		}
		if doneReading {
			break
		}
	} // done reading/writing

	err = sink.Close()
	ass.Nil(err)

	err = outfile.Close()
	ass.Nil(err)

	written, err := os.Open(path)
	ass.Nil(err)
	ass.NotNil(written)

	actual := aiff.NewDecoder(written)
	ass.NotNil(actual)
	actual.ReadInfo()

	ass.Equal(in.BitDepth, actual.BitDepth)
	t.Logf("Copied file://%v", infile.Name())
	t.Logf("To file://%v", path)
}
