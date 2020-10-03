package sin

import (
	"math"

	"osampler/audio"
	"osampler/transform"
)

type sin struct {
	transform.BaseTransform
	frequency float64
	phase     float64

	ctx         audio.Context
	buf         audio.Buffer
	currentTime float64
	bufferTime  float64
}

func New(context audio.Context, buffer audio.Buffer, frequency float64, phase float64) transform.Transform {
	return &sin{
		buf:         buffer,
		frequency:   frequency,
		phase:       phase,
		ctx:         context,
		currentTime: 0,
		bufferTime:  1 / float64(context.SampleRate()),
	}
}

func (s *sin) CalculateBuffer() int {
	for sample := 0; sample < s.buf.Size(); sample++ {
		v := math.Sin(2*math.Pi*s.frequency*s.currentTime + s.phase)
		s.buf.Data()[sample] = v
		s.currentTime += s.bufferTime
	}
	return s.buf.Size()
}
