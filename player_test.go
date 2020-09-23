package osampler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/test"
)

func TestPlayerBasics(t *testing.T) {
	require := require.New(t)
	path := test.ResolvePath("audio/sound.mp3")
	require.FileExists(path)
}
