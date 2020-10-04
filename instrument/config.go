package instrument

import (
	"osampler/midi"
	"osampler/sample"
)

type Config interface {
	SamplesFor(note midi.Note) []sample.Sample
}
