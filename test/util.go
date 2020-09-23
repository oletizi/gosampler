package test

import (
	"log"
	"path"
	"path/filepath"

	"osampler/util"
)

// Resolves a subpath to a full path in the test directory for easy test asset resolution.
func ResolvePath(subpath string) string {
	filename := util.MyPath()
	rv := filepath.Join(path.Dir(filename), subpath)
	log.Printf("Resolved path: %v", rv)
	return rv
}
