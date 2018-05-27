package gonch

import (
	"io"
	"os"
)

// Operator operations for TempDir
type Operator interface {
	AddDir(name, path string, perm os.FileMode) error
	AddFile(name, path string, content []byte, perm os.FileMode) error
	AddFiles(files []*TempFile) error
	File(name string) (*TempFile, error)
	Dir() string
	io.Closer
}
