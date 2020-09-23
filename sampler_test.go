package osampler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/test"
)

func TestSamplerBasics(t *testing.T) {
	require := require.New(t)
	path := test.ResolvePath("midi/jesu.mid")
	require.FileExistsf(path, "Test file missing: %v", path)
	err := LoadFile(path)
	require.Nil(err)
}
