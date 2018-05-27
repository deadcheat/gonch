package gonch

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// TempDir a temporary-dir
type TempDir struct {
	TempFiles map[string]*TempFile
	Path      string
}

// AddDir generate dir in generated temporary-dir
func (t *TempDir) AddDir(name, path string, perm os.FileMode) error {
	p := filepath.Join(t.Path, path)
	if err := os.MkdirAll(p, perm); err != nil {
		return err
	}
	t.TempFiles[name] = newTempFile(name, p, nil, perm)
	return nil
}

// AddFile add file to generated temporary-dir
func (t *TempDir) AddFile(name, path string, content []byte, perm os.FileMode) error {
	p := filepath.Join(t.Path, path)
	dp := filepath.Dir(p)
	f, err := os.Stat(dp)
	if err != nil {
		// file does not exists
		if err := os.MkdirAll(dp, os.ModePerm); err != nil {
			return err
		}
	}
	if f != nil && !f.IsDir() {
		return ErrInvalidStatus
	}
	if err := ioutil.WriteFile(p, content, perm); err != nil {
		return err
	}
	t.TempFiles[name] = newTempFile(name, p, content, perm)
	return nil
}

// AddFiles add specified files
func (t *TempDir) AddFiles(files []*TempFile) error {
	for i := range files {
		f := files[i]
		// may need to implement more for changing behavior when file duplicates
		if err := t.AddFile(f.Name, f.Path, f.Content, f.Permission); err != nil {
			return err
		}
		t.TempFiles[f.Name] = f
	}
	return nil
}

// File returns created temporary-file
func (t *TempDir) File(name string) (*TempFile, error) {
	f, ok := t.TempFiles[name]
	if !ok {
		return nil, ErrNotExists
	}
	return f, nil
}

// Dir returns their Path
func (t *TempDir) Dir() string {
	return t.Path
}

// Close implemetation of io.Closer
func (t *TempDir) Close() error {
	return os.RemoveAll(t.Path)
}

// New returns TempDirSerivice implements
func New(dir, prefix string) Operator {
	dir, err := ioutil.TempDir(dir, prefix)
	if err != nil {
		panic(err)
	}
	fm := make(map[string]*TempFile)
	return &TempDir{
		TempFiles: fm,
		Path:      dir,
	}
}
