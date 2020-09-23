package util

import (
	"log"
	"path"
	"runtime"
)

// Returns the path to the file containing the calling function
func MyPath() string {
	// runtime voodoo to find the directory of the current file.
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Can't find myself or you!")
	}
	return filename
}

// Returns the path to the directory containing the calling function
func MyDir() string {
	return path.Dir(MyPath())
}
