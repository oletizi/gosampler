package mixer

import (
	"log"

	"osampler/audio"
	"osampler/transform"
)

// I suppose we could call this a multiplexer, but this is audio, not electronics
type mixer struct {
	transform.BaseTransform
	inputs []audio.Buffer
	output audio.Buffer
}

func New(inputs []audio.Buffer, output audio.Buffer) transform.Transform {
	return &mixer{transform.BaseTransform{}, inputs, output}
}

func (m *mixer) CalculateBuffer() int {
	// find  the size of the largest input buffer
	maxBytes := 0
	for _, v := range m.inputs {
		if v.Size() > maxBytes {
			maxBytes = v.Size()
		}
	}

	// resize the output buffer to match the largest input buffer
	if maxBytes != m.output.Size() {
		m.output.Resize(maxBytes)
		log.Printf("Resized output buffer: %v", m.output.Size())
	}
	// For each input buffer...
	for buffersIndex, input := range m.inputs {
		// process each buffer
		for samplesIndex, newValue := range input.Data() {
			var currentValue float64
			// if this is the first buffer, initialized the output buffer to zero (otherwise we're mixing garbage from the
			// previous output buffer in with the current output buffer
			if buffersIndex == 0 {
				currentValue = 0
			} else {
				currentValue = m.output.Data()[samplesIndex]
			}
			sum := currentValue + newValue
			if sum > 1 {
				// hard clip
				sum = 1
			}
			m.output.Data()[samplesIndex] = sum
		}
	}
	return m.output.Size()
}
