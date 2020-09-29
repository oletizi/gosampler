package sfz

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	note2 "osampler/note"
	"osampler/test"
)

func TestNewConfigFactory(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ass := require.New(t)
	config, err := New(test.ResolvePath("sfz/test.sfz"))
	ass.NotNil(config)
	ass.Nil(err)

	note, err := note2.New(21)
	ass.Nil(err)

	log.Printf("calling SamplesFor(%v)", note)
	samples := config.SamplesFor(note)
	ass.NotNil(samples)
	log.Printf("samples: %v", samples)
}
