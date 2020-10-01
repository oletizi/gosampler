package transform

import "osampler/audio"

type Transform interface {
	CalculateBuffer(b audio.Buffer)
}
