package osampler

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasics(t *testing.T) {
	assert := assert.New(t)
	cwd, err := os.Getwd()
	log.Printf("working directory: cwd: %v, err: %v", cwd, err)

	path := filepath.Join(cwd, "../test/midi/jesu.mid")
	assert.NotNil(path)

	LoadFile(path)
}
