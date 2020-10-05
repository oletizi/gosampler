package instrument

import (
	"osampler/midi"
)

type Config interface {
	FilesFor(note midi.Note) []string
}
