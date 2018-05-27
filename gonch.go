package gonch

import (
	"os"
)

// TempFile a temporary-file in temporary-dir
type TempFile struct {
	Name       string
	Path       string
	Content    []byte
	Permission os.FileMode
}

// TempDir a temporary-dir
type TempDir struct {
	TempFiles  map[string]*TempFile
	Path       string
	Permission os.FileMode
}
