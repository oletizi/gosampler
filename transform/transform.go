package transform

import "osampler/audio"

type Transform interface {
	CalculateBuffer()
	Buffer() audio.Buffer
}
