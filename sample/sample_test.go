package sample

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	ass := require.New(t)
	filename := "the filename"
	s := New(filename)
	ass.NotNil(s)
	ass.Equal(filename, s.Filename())
}
