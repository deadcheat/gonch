# gonch

[![Build Status](https://travis-ci.org/deadcheat/gonch.svg?branch=master)](https://travis-ci.org/deadcheat/gonch) [![Coverage Status](https://coveralls.io/repos/github/deadcheat/gonch/badge.svg?branch=master&service=github)](https://coveralls.io/github/deadcheat/gonch?branch=master&service=github) [![GoDoc](https://godoc.org/github.com/deadcheat/gonch?status.svg)](https://godoc.org/github.com/deadcheat/gonch)

## What is this ?
This is a library for creating temporary file in temporary directory and removing all of temporary directory.
I don't want to write same codes many time. that motivated me.

## How to use it

Firs of all, use `go get` to download
```
go get -u github.com/deadcheat/gonch
```

### Create a temporary directory

```
d := gonch.New("", "")
defer d.Close()
```

Arguments for gonch.New are same as `ioutil.TempDir`'s.

### add file to temporary directory

```
err := d.AddFile("name", "path from temp-dir root", []byte("hello, world"), os.ModePerm)
```

Its arguments are below

- name: it's just a name, identify file you want to create itself
- path: file path from your temporary dir, it's okay whether relative path or absolute path, 'cause they'll be joined
- content: byte slice, a body of data
- permission: file permission.


### Load file using its name

```
f, err := d.File("file name")
```

### Add multiple files

```
err := d.AddFiles([]*TempFile{
		&TempFile{
			Name:       "testfile",
			Path:       "testdir/testfile.txt",
			Content:    []byte("sample"),
			Permission: os.ModePerm,
		},
		&TempFile{
			Name:       "testfile2",
			Path:       "testdir/testfile2.txt",
			Content:    []byte("sample2"),
			Permission: os.ModePerm,
		},
	})
```
