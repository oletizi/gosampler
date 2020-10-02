package mixer

import (
	"osampler/audio"
	"osampler/transform"
)

// I suppose we could call this a multiplexer, but this is audio, not electronics
type mixer struct {
	inputs []audio.Buffer
	output audio.Buffer
}

func (m *mixer) CalculateBuffer() {
	for _, input := range m.inputs {
		for i, v := range input.Data() {
			m.output.Data()[i] += v
		}
	}
}

func New(inputs []audio.Buffer, output audio.Buffer) transform.Transform {
	return &mixer{inputs, output}
}
