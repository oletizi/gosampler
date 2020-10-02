package mixer

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)
	c := audio.NewContext(44100, 16, 2)
	a := audio.NewBuffer(c, 2)
	b := audio.NewBuffer(c, 2)
	output := audio.NewBuffer(c, 2)
	a.Data()[0] = .5
	b.Data()[0] = -.5

	a.Data()[0] = 1
	b.Data()[0] = 2.1

	inputs := []audio.Buffer{a, b}
	mix := New(inputs, output)
	ass.NotNil(mix)

	mix.CalculateBuffer()

	ass.Equal(a.Data()[0]+b.Data()[0], output.Data()[0])
	ass.Equal(a.Data()[1]+b.Data()[1], output.Data()[1])
}

func TestMixAudioFiles(t *testing.T) {
	//ass := require.New(t)

}
