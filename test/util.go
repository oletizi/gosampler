package test

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"osampler/util"
)

// Resolves a subpath to a full path in the test directory for easy test asset resolution.
func ResolvePath(subpath string) string {
	filename := util.MyPath()
	rv := filepath.Join(path.Dir(filename), subpath)
	return rv
}

// Wrapper around ioutil.TempFile to add a timestamp so temp files will sort by name properly
func TempFile(pattern string) (*os.File, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) + "-*"
	timestampedPattern := strings.ReplaceAll(pattern, "*", timestamp)
	return ioutil.TempFile("", timestampedPattern)
}
