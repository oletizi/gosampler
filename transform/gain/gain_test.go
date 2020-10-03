package gain

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)
	b := audio.NewBuffer(2)
	data := b.Data()
	data[0] = 1
	data[1] = .5
	factor := float64(1000)
	gain := New(b, factor)
	ass.NotNil(gain)
	ass.Equal(1.0, b.Data()[0])
	ass.Equal(.5, b.Data()[1])

	gain.CalculateBuffer()

	ass.Equal(1*factor, b.Data()[0])
	ass.Equal(.5*factor, b.Data()[1])
}
