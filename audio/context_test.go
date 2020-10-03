package audio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvenienceFunctions(t *testing.T) {
	c := NewCDContext()
	assertContext(t, c, c.SampleRate(), c.BitDepth().Depth(), c.ChannelCount())
}

func TestNew(t *testing.T) {
	sampleRate := 48000
	bitDepth := 16
	channelCount := 5
	c := NewContext(sampleRate, NewBitDepth16(), channelCount)
	assertContext(t, c, sampleRate, bitDepth, channelCount)
}

func assertContext(t *testing.T, c Context, rate int, depth int, count int) {
	ass := require.New(t)
	ass.NotNil(c)
	ass.Equal(rate, c.SampleRate())
	ass.NotNil(c.BitDepth())
	ass.Equal(depth, c.BitDepth().Depth())
	ass.Equal(count, c.ChannelCount())
}
