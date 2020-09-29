package sfz

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"osampler"
	"osampler/note"
	"osampler/sample"
	"osampler/test"
)

func TestNewConfigFactory(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ass := require.New(t)
	config, err := New(test.ResolvePath("sfz/test.sfz"))
	ass.NotNil(config)
	ass.Nil(err)

	var expected []osampler.Sample
	expected = append(expected, sample.New(samplePath("sample/37.wav")))

	theNote, err := note.New(37)
	ass.Nil(err)

	// Test for key=
	samples := config.SamplesFor(theNote)
	ass.NotNil(samples)
	ass.Equal(expected, samples)

	// Test for key range
	expected[0] = sample.New(samplePath("sample/13-24.wav"))
	theNote, _ = note.New(20)
	samples = config.SamplesFor(theNote)
	ass.Equal(expected, samples)

	// Test for overlapping regions
	expected[0] = sample.New(samplePath("sample/25-30.wav"))
	expected = append(expected, sample.New(samplePath("sample/27-36.wav")))
	theNote, _ = note.New(29)
	samples = config.SamplesFor(theNote)
	ass.Equal(expected, samples)
}

func samplePath(path string) string {
	rv := test.ResolvePath(filepath.Join("sfz", path))
	return rv
}
