package mock

import (
	"github.com/faiface/beep"

	"osampler/audio"
)

type mockOut struct {
	PlayCalls int
}

func NewMockOut() audio.Out {
	return &mockOut{}
}

func (m *mockOut) Play(device beep.StreamSeekCloser) {
	m.PlayCalls++
}
