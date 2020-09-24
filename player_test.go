package osampler

import (
	"testing"

	"github.com/faiface/beep"
	"github.com/stretchr/testify/require"

	"osampler/test"
)

type mockOut struct {
	PlayCalls int
}

func TestPlayFile(t *testing.T) {
	out := &mockOut{}
	require := require.New(t)
	path := test.ResolvePath("audio/sound.mp3")
	require.FileExists(path)

	err := PlayFile(out, path)
	require.Nil(err)
	require.Equal(1, out.PlayCalls)
}

func (m *mockOut) Play(device beep.StreamSeekCloser) {
	m.PlayCalls++
}
func (m *mockOut) PlayCall(device beep.StreamCloser) {

}
