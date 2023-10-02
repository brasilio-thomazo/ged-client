package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path     string
	Name     string
	Filename string
	Filesize int64
	File     *os.File
}

func NewFile(name, path string) (*File, error) {
	f := &File{Path: path, Name: name}

	if strings.IndexAny(path, "/") != 0 {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		f.Path = filepath.Join(dir, path)
	}
	f.Filename = filepath.Join(f.Path, name)
	finfo, _ := os.Stat(f.Filename)
	if finfo != nil {
		f.Filesize = finfo.Size()
	}
	return f, nil
}

func (f *File) Create() error {
	_, err := os.Stat(f.Path)
	if os.IsNotExist(err) {
		if err := os.Mkdir(f.Path, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(f.Filename)
	if err != nil {
		return err
	}

	f.File = file
	return nil
}

func (f *File) Open() error {
	_, err := os.Stat(f.Filename)
	if err != nil {
		return err
	}

	file, err := os.Open(f.Filename)
	if err != nil {
		return err
	}
	f.File = file
	return nil
}
