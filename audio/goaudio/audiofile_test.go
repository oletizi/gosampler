package goaudio

import (
	"io"
	"io/ioutil"
	"math"
	"os"
	"testing"

	"github.com/go-audio/aiff"
	"github.com/go-audio/audio"
	"github.com/stretchr/testify/require"

	audio2 "osampler/audio"
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

	buf := &audio.IntBuffer{
		Format:         in.Format(),
		Data:           make([]int, bufSize),
		SourceBitDepth: int(in.BitDepth),
	}
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

func TestGenAndWrite(t *testing.T) {
	ass := require.New(t)
	outfile, err := test.TempFile("test-*.aif") //ioutil.TempFile("", "test-"+strconv.FormatInt(time.Now().Unix(), 10)+"-*.aif")
	ass.Nil(err)
	outfileName := outfile.Name()

	bufferSize := 512

	ctx := audio2.NewContext(44100, 16, 1)
	buf := audio2.NewBuffer(ctx, bufferSize)
	iterations := 2 ^ 256
	out := aiff.NewEncoder(outfile, ctx.SampleRate(), ctx.BitDepth(), ctx.ChannelCount())

	outBuf := audio.IntBuffer{
		Format:         audio.FormatMono44100,
		Data:           make([]int, bufferSize),
		SourceBitDepth: ctx.BitDepth(),
	}
	for b := 0; b < iterations; b++ {
		fillBuffer(buf)
		copyBuffer(buf, outBuf)
		err := out.Write(&outBuf)
		ass.Nil(err)
	}
	err = out.Close()
	ass.Nil(err)

	err = outfile.Close()
	ass.Nil(err)

	t.Logf("wrote file://%v", outfileName)
}

func copyBuffer(a audio2.Buffer, b audio.IntBuffer) audio.IntBuffer {
	for i := 0; i < a.Size(); i++ {
		b.Data[i] = int(a.Data()[i])
	}
	return b
}

func fillBuffer(buf audio2.Buffer) {
	var amplitude float64 = 2500
	var frequency float64 = 440

	var phase float64 = 0
	var deltaTime = 1 / float64((buf.Context()).SampleRate())
	var currentTime float64 = 0

	for sample := 0; sample < buf.Size(); sample++ {
		v := amplitude * math.Sin(2*math.Pi*frequency*currentTime+phase)
		buf.Data()[sample] = v
		currentTime += deltaTime
	}
}
