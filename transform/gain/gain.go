package gain

import (
	"osampler/audio"
	"osampler/transform"
)

type gain struct {
	transform.BaseTransform

	buffer audio.Buffer
	factor float64
}

func New(buffer audio.Buffer, factor float64) transform.Transform {
	return &gain{transform.BaseTransform{}, buffer, factor}
}

func (g *gain) CalculateBuffer() int {
	for i, v := range g.buffer.Data() {
		g.buffer.Data()[i] = g.factor * v
	}
	return g.buffer.Size()
}
