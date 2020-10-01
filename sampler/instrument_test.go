package sampler

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler"
)

func TestNew(t *testing.T) {
	ass := require.New(t)
	var config osampler.Config = &mockConfig{}
	inst := New(config)
	ass.NotNil(inst)
}

// Mock Config
type mockConfig struct {
	samplesForCalls []osampler.Note
}

func (m *mockConfig) SamplesFor(note osampler.Note) []osampler.Sample {
	m.samplesForCalls = append(m.samplesForCalls, note)
	return []osampler.Sample{}
}
