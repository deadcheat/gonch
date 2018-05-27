package gonch

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewError(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
		t.Error("should be panic")
	}()
	d := New("test", "test")
	defer d.Close()
	d = New("test", "test")
	defer d.Close()
}

func TestFileError(t *testing.T) {
	d := New("", "")
	defer d.Close()

	if _, err := d.File("not exists"); err == nil {
		t.Error("should return error")
	}
}

func TestAddFile(t *testing.T) {
	d := New("", "")
	defer d.Close()

	// simple test
	testNameNormal := "testNormal"
	testPathNormal := "test.txt"
	testContentNormal := []byte("sample")
	testPermNormal := os.FileMode(0666)
	if err := d.AddFile(testNameNormal, testPathNormal, testContentNormal, testPermNormal); err != nil {
		t.Error("unexpected to occur error:", err)
	}
	f, err := d.File(testNameNormal)
	if err != nil {
		t.Error("unexpected to occur error:", err)
	}
	if f == nil ||
		f.Name != testNameNormal ||
		f.Path != filepath.Join(d.Dir(), testPathNormal) ||
		!reflect.DeepEqual(f.Content, testContentNormal) ||
		f.Permission != testPermNormal {
		t.Errorf("added file is not expected one, actual:%#+v, expected:%#+v", f, &TempFile{
			Name:       testNameNormal,
			Path:       filepath.Join(d.Dir(), testPathNormal),
			Content:    testContentNormal,
			Permission: testPermNormal,
		})
	}
}

func TestAddFileWithSubDir(t *testing.T) {
	d := New("", "")
	defer d.Close()

	// when path includes subdir
	testNameAddWithDir := "testAddWithDir"
	testPathAddWithDir := "/testdir/test.txt"
	testContentAddWithDir := []byte("sample")
	testPermAddWithDir := os.FileMode(0666)
	if err := d.AddFile(testNameAddWithDir, testPathAddWithDir, testContentAddWithDir, testPermAddWithDir); err != nil {
		t.Error("unexpected to occur error:", err)
	}
	f, err := d.File(testNameAddWithDir)
	if err != nil {
		t.Error("unexpected to occur error:", err)
	}
	if f == nil ||
		f.Name != testNameAddWithDir ||
		f.Path != filepath.Join(d.Dir(), testPathAddWithDir) ||
		!reflect.DeepEqual(f.Content, testContentAddWithDir) ||
		f.Permission != testPermAddWithDir {
		t.Errorf("added file is not expected one, actual:%#+v, expected:%#+v", f, &TempFile{
			Name:       testNameAddWithDir,
			Path:       filepath.Join(d.Dir(), testPathAddWithDir),
			Content:    testContentAddWithDir,
			Permission: testPermAddWithDir,
		})
	}
}

func TestAddFileErrorWhenWriteInvalidDir(t *testing.T) {
	d := New("", "")
	defer d.Close()

	duplicatedDir := "/duplicate"
	err := d.AddDir("duplicatedDir", duplicatedDir, 0444)
	if err != nil {
		panic(err)
	}
	if err = d.AddFile("invaliderr", filepath.Join(duplicatedDir, "invalid"), nil, 0666); err == nil {
		t.Errorf("should return error when add to invalid dir")
	}
}

func TestAddFileErrorWhenWriteDirToInvalidDir(t *testing.T) {
	d := New("", "")
	defer d.Close()

	duplicatedDir := "/duplicate"
	err := d.AddDir("duplicatedDir", duplicatedDir, 0444)
	if err != nil {
		panic(err)
	}
	if err = d.AddFile("invaliderr", filepath.Join(duplicatedDir, "/invaliddir/invalid"), nil, 0666); err == nil {
		t.Errorf("should return error when add to invalid dir")
	}
}

func TestAddFileErrorWhenParentIsNotDir(t *testing.T) {
	d := New("", "")
	defer d.Close()

	duplicatedDir := "/duplicate/invalidfile"
	err := d.AddFile("duplicatedDir", duplicatedDir, []byte("test"), 0666)
	if err != nil {
		panic(err)
	}
	if err = d.AddFile("invaliderr", filepath.Join(duplicatedDir, "/invalid"), nil, 0666); err == nil {
		t.Errorf("should return error when add to invalid dir")
	}
	if err != ErrInvalidStatus {
		t.Error("error should be ErrInvalidStatus but ", err)
	}
}

func TestAddDirErrorWhenInvalidPermission(t *testing.T) {

	d := New("", "")
	defer d.Close()

	duplicatedDir := "/duplicate/invalid"
	err := d.AddDir("duplicatedDir", duplicatedDir, 0444)
	if err == nil {
		t.Error("should return error")
	}
}

func TestAddFiles(t *testing.T) {
	d := New("", "")
	defer d.Close()

	err := d.AddFiles([]*TempFile{
		&TempFile{
			Name:       "testfile",
			Path:       "testdir/testfile.txt",
			Content:    []byte("sample"),
			Permission: os.ModePerm,
		},
	})
	if err != nil {
		t.Error("should not return error ", err)
	}
}

func TestAddFilesError(t *testing.T) {
	d := New("", "")
	defer d.Close()

	if err := d.AddDir("invalidDir", "testdir", 0444); err != nil {
		panic(err)
	}

	err := d.AddFiles([]*TempFile{
		&TempFile{
			Name:       "testfile",
			Path:       "testdir/testfile.txt",
			Content:    []byte("sample"),
			Permission: os.ModePerm,
		},
	})
	if err == nil {
		t.Error("should return error")
	}
}
