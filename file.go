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

func newTempFile(n, p string, c []byte, perm os.FileMode) *TempFile {
	return &TempFile{
		Name:       n,
		Path:       p,
		Content:    c,
		Permission: perm,
	}
}
