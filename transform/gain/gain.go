package gain

import (
	"osampler/audio"
	"osampler/transform"
)

type gain struct {
	buffer audio.Buffer
	factor float64
}

func New(buffer audio.Buffer, factor float64) transform.Transform {
	return gain{buffer, factor}
}

func (g gain) CalculateBuffer() {
	for i, v := range g.buffer.Data() {
		g.buffer.Data()[i] = g.factor * v
	}
}

func (g gain) Buffer() audio.Buffer {
	return g.buffer
}
