package sfz

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/midi"
	"osampler/test"
)

func TestNewConfigFactory(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ass := require.New(t)
	config, err := New(test.ResolvePath("sfz/test.sfz"))
	ass.NotNil(config)
	ass.Nil(err)

	var expected []string
	expected = append(expected, samplePath("37.wav"))

	theNote, err := midi.NewNote(37)
	ass.Nil(err)

	// Test for key
	files := config.FilesFor(theNote)
	ass.NotNil(files)
	ass.Equal(expected, files)

	// Test for key range
	expected[0] = samplePath("13-24.wav")
	theNote, _ = midi.NewNote(20)
	files = config.FilesFor(theNote)
	ass.Equal(expected, files)

	// Test for overlapping regions
	expected[0] = samplePath("25-30.wav")
	expected = append(expected, samplePath("27-36.wav"))
	theNote, _ = midi.NewNote(29)
	files = config.FilesFor(theNote)
	ass.Equal(expected, files)
}

func samplePath(path string) string {
	rv := test.ResolvePath(filepath.Join("sfz", "sample", path))
	return rv
}
