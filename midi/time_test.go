package midi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewContext(t *testing.T) {
	ass := require.New(t)
	tempo := 120.0
	ppq := 960.0
	c := NewContext(tempo, ppq)
	ass.NotNil(c)
	ass.Equal(ppq, c.PPQ())
	ass.Equal(tempo, c.Tempo())
	ass.Equal(60*1000/(tempo*ppq), c.MsPerTick())
}
