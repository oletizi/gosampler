package aiff

import (
	"fmt"
	"io"
	"log"
	"os"

	aiffv "github.com/go-audio/aiff"
	audiov "github.com/go-audio/audio"

	"osampler/audio"
	"osampler/transform"
)

//
// AIFF file reader
//

type aiffInput struct {
	transform.BaseTransform
	buffer   audio.Buffer
	context  audio.Context
	in       *aiffv.Decoder
	inBuffer *audiov.IntBuffer
}

func contextFromFile(in *aiffv.Decoder) (audio.Context, error) {
	in.ReadInfo()
	var err error
	var bitDepth audio.BitDepth
	switch in.BitDepth {
	case 16:
		bitDepth = audio.NewBitDepth16()
		break
	default:
		//err = errors.New("unsupported bit depth: " + strconv.FormatUint(uint64(in.BitDepth), 10))
		panic("Unsupported bit depth")
	}
	return audio.NewContext(in.SampleRate, bitDepth, int(in.NumChans)), err
}

func NewAiffInput(buffer audio.Buffer, infile *os.File) (transform.Transform, audio.Context, error) {
	in := aiffv.NewDecoder(infile)
	context, err := contextFromFile(in)
	if err != nil {
		return nil, nil, err
	}
	return &aiffInput{buffer: buffer, context: context, in: in, inBuffer: newIntBuffer(context, buffer.Size())}, context, err
}

func (a *aiffInput) CalculateBuffer() int {
	bytesRead, err := a.in.PCMBuffer(a.inBuffer)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		log.Printf("Error reading from input: %v", err)
	}
	if bytesRead != len(a.inBuffer.Data) {
		// we didn't read a full buffer, so trim the buffer
		a.inBuffer.Data = a.inBuffer.Data[:bytesRead]
		a.buffer.Resize(bytesRead)
		fmt.Printf("========> bytes read less than buffer: %v; resized buffer: %v", bytesRead, a.buffer.Size())
	}
	for i := 0; i < len(a.inBuffer.Data); i++ {
		// XXX: Make divisor dependent on bit depth
		intValue := a.inBuffer.Data[i]
		floatValue := float64(intValue) / a.context.BitDepth().MaxInt()
		//if i%1000 == 0 {
		//	fmt.Printf("read: intValue: %v; floatValue: %v\n", intValue, floatValue)
		//}
		a.buffer.Data()[i] = floatValue
	}
	return bytesRead
}

//
// AIFF file writer
//

type aiffOutput struct {
	transform.BaseTransform
	buffer    audio.Buffer
	context   audio.Context
	outBuffer *audiov.IntBuffer
	out       *aiffv.Encoder
}

func NewAiffOutput(context audio.Context, buffer audio.Buffer, outfile *os.File) transform.Transform {
	out := aiffv.NewEncoder(outfile, context.SampleRate(), context.BitDepth().Depth(), context.ChannelCount())
	return &aiffOutput{buffer: buffer, context: context, out: out, outBuffer: newIntBuffer(context, buffer.Size())}
}

func (a *aiffOutput) CalculateBuffer() int {
	// convert and copy floating-point buffer data to the go-audio buffer
	for i := 0; i < a.buffer.Size(); i++ {
		floatValue := a.buffer.Data()[i]
		intValue := int(floatValue * a.context.BitDepth().MaxInt())

		//if i%1000 == 0 {
		//	fmt.Printf("floatValue: %v; int value: %v\n", floatValue, intValue)
		//}
		a.outBuffer.Data[i] = intValue
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

func newIntBuffer(c audio.Context, size int) *audiov.IntBuffer {
	return &audiov.IntBuffer{
		Format: &audiov.Format{
			NumChannels: c.ChannelCount(),
			SampleRate:  c.SampleRate(),
		},
		Data:           make([]int, size),
		SourceBitDepth: c.BitDepth().Depth(),
	}
}
