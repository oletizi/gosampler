package audio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)

	buf := NewBuffer(2)
	ass.NotNil(buf)

	buf.Data()[1] = 1
	assertZeroOne(ass, buf)

	ass.Equal(2, buf.Size())

	buf.Resize(3)
	ass.Equal(3, buf.Size())

	buf.Resize(0)
	ass.Equal(0, buf.Size())
}

// make sure data survives intact when passing by value
func assertZeroOne(ass *require.Assertions, b Buffer) {
	ass.Equal(float64(0), b.Data()[0])
	ass.Equal(float64(1), b.Data()[1])
}
