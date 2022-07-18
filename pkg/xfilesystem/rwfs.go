package xfilesystem

import (
	"io/fs"
	"os"
	"path/filepath"
)

type RwFS interface {
	fs.FS

	BasePath() string
	Sub(dir string) (RwFS, error)

	Mkdir(name string) error
	CreateFile(name string) (*os.File, error)
}

type rwFilesystem struct {
	filesystem
}

func NewRWFileSystem(path string, fsys fs.FS) RwFS {
	return &rwFilesystem{
		filesystem: filesystem{
			basePath: path,
			fsys:     fsys,
		},
	}
}

func (f *rwFilesystem) Mkdir(name string) error {
	return os.Mkdir(filepath.Join(f.basePath, name), 0755)
}

func (f *rwFilesystem) CreateFile(name string) (*os.File, error) {
	return os.Create(filepath.Join(f.basePath, name))
}

func (f *rwFilesystem) Sub(dir string) (RwFS, error) {
	newFs, err := f.filesystem.Sub(dir)
	if err != nil {
		return nil, err
	}

	return NewRWFileSystem(newFs.BasePath(), newFs), nil
}
