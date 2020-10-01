package audio

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/go-audio/aiff"
	"github.com/go-audio/audio"
	"github.com/stretchr/testify/require"

	"osampler/test"
)

func TestReadWrite(t *testing.T) {
	const bufSize = 512
	ass := require.New(t)
	infile, err := os.Open(test.ResolvePath("audio/sound.aif"))

	ass.Nil(err)
	ass.NotNil(infile)
	in := aiff.NewDecoder(infile)
	in.ReadInfo()

	t.Logf("in.Size: %v", in.Size)
	t.Logf("sample rate: %v", in.SampleRate)
	t.Logf("bit depth: %v", in.BitDepth)
	t.Logf("sample frames: %v", in.NumSampleFrames)
	t.Logf("channels: %v", in.NumChans)
	t.Logf("channels * sample frames: %v", uint32(in.NumChans)*in.NumSampleFrames)
	t.Logf("pcm size: %v", in.PCMSize)

	var buf *audio.IntBuffer = &audio.IntBuffer{
		Format:         in.Format(),
		Data:           make([]int, bufSize),
		SourceBitDepth: int(in.BitDepth),
	}
	err = in.FwdToPCM()
	ass.Nil(err)

	outfile, err := ioutil.TempFile("", "test-*.aif")
	ass.Nil(err)
	outfileName := outfile.Name()

	out := aiff.NewEncoder(outfile, in.SampleRate, int(in.BitDepth), int(in.NumChans))

	totalBytesRead := 0
	bytesRead := 0
	var doneReading bool
	for err == nil {
		bytesRead, err = in.PCMBuffer(buf)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			ass.FailNow("Error reading from input: %v", err)
		}
		if bytesRead != len(buf.Data) {
			// we didn't read a full buffer, so trim the buffer
			buf.Data = buf.Data[:bytesRead]
			doneReading = true
		}
		if err = out.Write(buf); err != nil {
			ass.FailNow("Error writing to output: %v", err)
		}
		if doneReading {
			break
		}
	} // done reading/writing

	t.Logf("Read %v bytes", totalBytesRead)

	if err = out.Close(); err != nil {
		ass.FailNow("Failed to close the output stream: %v", err.Error())
	}

	if err = outfile.Close(); err != nil {
		ass.FailNow("Failed to close the output file: %v", err.Error())
	}

	t.Logf("wrote file://%v", outfileName)

	expected, err := os.Open(outfileName)
	ass.Nil(err)
	ass.NotNil(expected)

	d := aiff.NewDecoder(expected)
	ass.NotNil(d)
	d.ReadInfo()
	ass.Equal(in.SampleRate, d.SampleRate)
	ass.Equal(in.BitDepth, d.BitDepth)
	ass.Equal(in.NumChans, d.NumChans)
	ass.Equal(in.Encoding, d.Encoding)
}
