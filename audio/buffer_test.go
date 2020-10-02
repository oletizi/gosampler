package audio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)

	c := NewContext(441000, 16, 2)
	buf := NewBuffer(c, 2)
	ass.NotNil(buf)

	buf.Data()[1] = 1
	assertZeroOne(ass, buf)

	ass.Equal(2, buf.Size())
}

func assertZeroOne(ass *require.Assertions, b Buffer) {
	ass.Equal(float64(0), b.Data()[0])
	ass.Equal(float64(1), b.Data()[1])
}
