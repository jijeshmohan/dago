package xfilesystem

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

type RwFS interface {
	fs.FS

	BasePath() string
	Sub(dir string) (RwFS, error)

	Mkdir(name string) error
	CreateFile(name string) (*os.File, error)
	CleanupNewFiles() error
}

type rwFilesystem struct {
	filesystem
	newFilesMap map[string]bool
}

func NewRWFileSystem(path string, fsys fs.FS) RwFS {
	return &rwFilesystem{
		filesystem: filesystem{
			basePath: path,
			fsys:     fsys,
		},
		newFilesMap: make(map[string]bool),
	}
}

func (f *rwFilesystem) Mkdir(name string) error {
	if found, _ := exists(filepath.Join(f.basePath, name)); found {
		return os.ErrExist
	}

	f.newFilesMap[filepath.Join(f.basePath, name)] = true

	return os.Mkdir(filepath.Join(f.basePath, name), 0755)
}

func (f *rwFilesystem) CreateFile(name string) (*os.File, error) {
	if found, _ := exists(filepath.Join(f.basePath, name)); found {
		return nil, os.ErrExist
	}

	f.newFilesMap[filepath.Join(f.basePath, name)] = false
	return os.Create(filepath.Join(f.basePath, name))
}

func (f *rwFilesystem) Sub(dir string) (RwFS, error) {
	newFs, err := f.filesystem.Sub(dir)
	if err != nil {
		return nil, err
	}

	return &rwFilesystem{
		filesystem: filesystem{
			basePath: newFs.BasePath(),
			fsys:     newFs,
		},
		newFilesMap: f.newFilesMap,
	}, nil
}

func (f *rwFilesystem) CleanupNewFiles() error {
	toDelete := make([]string, 0, len(f.newFilesMap))
	for path := range f.newFilesMap {
		toDelete = append(toDelete, path)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(toDelete)))

	for _, path := range toDelete {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
