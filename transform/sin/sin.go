package sin

import (
	"math"

	"osampler/audio"
	"osampler/transform"
)

type sin struct {
	amplitude float64
	frequency float64
	phase     float64

	ctx         audio.Context
	buf         audio.Buffer
	currentTime float64
	bufferTime  float64
}

func New(amplitude float64, frequency float64, phase float64, context audio.Context, buffer audio.Buffer) transform.Transform {
	return sin{amplitude: amplitude,
		frequency:   frequency,
		phase:       phase,
		ctx:         context,
		buf:         buffer,
		currentTime: 0,
		bufferTime:  1 / float64(buffer.Context().SampleRate()),
	}
}

func (s sin) CalculateBuffer(b audio.Buffer) {

	for sample := 0; sample < s.buf.Size(); sample++ {
		v := s.amplitude * math.Sin(2*math.Pi*s.frequency*s.currentTime+s.phase)
		s.buf.Data()[sample] = v
		s.currentTime += s.bufferTime
	}
}
