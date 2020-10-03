package multi

import (
	"osampler/transform"
)

type multi struct {
	transform.BaseTransform
	set []transform.Transform
}

func (m *multi) CalculateBuffer() int {
	bytesProcessed := 0
	for _, v := range m.set {
		bytes := v.CalculateBuffer()
		if bytes > bytesProcessed {
			bytesProcessed = bytes
		}
	}
	return bytesProcessed
}

func New(set []transform.Transform) transform.Transform {
	return &multi{set: set}
}
