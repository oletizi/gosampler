package mixer

import (
	"osampler/audio"
	"osampler/transform"
)

// I suppose we could call this a multiplexer, but this is audio, not electronics
type mixer struct {
	transform.BaseTransform
	inputs []audio.Buffer
	output audio.Buffer
}

func (m *mixer) CalculateBuffer() int {
	for _, input := range m.inputs {
		for i, v := range input.Data() {
			m.output.Data()[i] += v
		}
	}
	return m.output.Size()
}

func New(inputs []audio.Buffer, output audio.Buffer) transform.Transform {
	// TODO: Add a check to make sure the inputs are all the same size
	return &mixer{transform.BaseTransform{}, inputs, output}
}
