package file

import (
	"io"
	"log"
	"os"

	"github.com/go-audio/aiff"
	goaudio "github.com/go-audio/audio"

	"osampler/audio"
	"osampler/transform"
)

//
// AIFF file reader
//

type aiffInput struct {
	transform.BaseTransform
	buffer   audio.Buffer
	in       *aiff.Decoder
	inBuffer *goaudio.IntBuffer
}

func NewAiffInput(buffer audio.Buffer, infile *os.File) transform.Transform {
	in := aiff.NewDecoder(infile)
	return &aiffInput{buffer: buffer, in: in, inBuffer: intBuffer(buffer.Context(), buffer.Size())}
}

func (a *aiffInput) CalculateBuffer() int {

	bytesRead, err := a.in.PCMBuffer(a.inBuffer)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		log.Printf("Error reading from input: %v", err)
	}
	if bytesRead != len(a.inBuffer.Data) {
		// we didn't read a full buffer, so trim the buffer
		a.inBuffer.Data = a.inBuffer.Data[:bytesRead]
	}
	for i := 0; i < len(a.inBuffer.Data); i++ {
		a.buffer.Data()[i] = float64(a.inBuffer.Data[i])
	}
	return bytesRead
}

//
// AIFF file writer
//

type aiffOutput struct {
	buffer    audio.Buffer
	outBuffer *goaudio.IntBuffer
	out       *aiff.Encoder
}

func (a *aiffOutput) CalculateBuffer() int {
	// convert and copy floating-point buffer data to the go-audio buffer
	for i := 0; i < a.buffer.Size(); i++ {
		a.outBuffer.Data[i] = int(a.buffer.Data()[i])
	}
	// write the go-audio buffer to the encoder.
	err := a.out.Write(a.outBuffer)
	if err != nil {
		log.Printf("Error writing: %v", err)
	}
	return a.buffer.Size()
}

func (a *aiffOutput) Close() error {
	return a.out.Close()
}

func NewAiffOutput(buffer audio.Buffer, outfile *os.File) transform.Transform {
	c := buffer.Context()
	out := aiff.NewEncoder(outfile, c.SampleRate(), c.BitDepth(), c.ChannelCount())
	return &aiffOutput{buffer: buffer, out: out, outBuffer: intBuffer(c, buffer.Size())}
}

func intBuffer(c audio.Context, size int) *goaudio.IntBuffer {
	return &goaudio.IntBuffer{
		Format: &goaudio.Format{
			NumChannels: c.ChannelCount(),
			SampleRate:  c.SampleRate(),
		},
		Data:           make([]int, size),
		SourceBitDepth: c.BitDepth(),
	}
}
