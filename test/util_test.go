package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvePath(t *testing.T) {
	ass := assert.New(t)
	resolved := ResolvePath("midi/jesu.mid")
	ass.FileExists(resolved)
}
