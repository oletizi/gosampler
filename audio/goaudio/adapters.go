package goaudio

import (
	"io"

	"github.com/go-audio/aiff"
	audio2 "github.com/go-audio/audio"

	"osampler/audio"
)

type sink struct {
	buf *audio2.IntBuffer
	out *aiff.Encoder
}

func (s sink) Close() error {
	return s.out.Close()
}

func (s sink) Write(b audio.Buffer) error {
	// initialize the go-audio buffer if it's nil or the wrong size
	if s.buf.Data == nil || len(s.buf.Data) != b.Size() {
		s.buf.Data = make([]int, b.Size())
	}
	// convert and copy floating-point buffer data to the go-audio buffer
	for i := 0; i < b.Size(); i++ {
		s.buf.Data[i] = int(b.Data()[i])
	}
	// write the go-audio buffer to the encoder.
	return s.out.Write(s.buf)
}

func NewAiffSink(context audio.Context, w io.WriteSeeker) audio.Sink {
	buf := audio2.IntBuffer{
		Format: &audio2.Format{
			NumChannels: context.ChannelCount(),
			SampleRate:  context.SampleRate(),
		},
		Data:           nil,
		SourceBitDepth: context.BitDepth(),
	}
	encoder := aiff.NewEncoder(w, context.SampleRate(), context.BitDepth(), context.ChannelCount())
	return sink{&buf, encoder}
}
