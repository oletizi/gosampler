package transform

import "osampler/audio"

type Transform interface {
	Buffer() audio.Buffer
	CalculateBuffer() int
	Close() error
}

type BaseTransform struct{}

func (n BaseTransform) Buffer() audio.Buffer { return nil }
func (n BaseTransform) CalculateBuffer() int { return 0 }
func (n BaseTransform) Close() error         { return nil }
