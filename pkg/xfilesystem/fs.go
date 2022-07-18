package xfilesystem

import (
	"io/fs"
	"path/filepath"
)

type FS interface {
	fs.FS
	BasePath() string
	Sub(dir string) (FS, error)
}

type filesystem struct {
	basePath string
	fsys     fs.FS
}

func NewFileSystem(path string, fsys fs.FS) FS {
	return &filesystem{
		basePath: path,
		fsys:     fsys,
	}
}

func (f *filesystem) Open(name string) (fs.File, error) {
	return f.fsys.Open(name)
}

func (f *filesystem) BasePath() string {
	return f.basePath
}

func (f *filesystem) Sub(dir string) (FS, error) {
	newFs, err := fs.Sub(f.fsys, dir)
	if err != nil {
		return nil, err
	}

	basePath := filepath.Join(f.basePath, dir)
	return NewFileSystem(basePath, newFs), nil
}
