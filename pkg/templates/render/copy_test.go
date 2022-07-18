package render

import (
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"github.com/jijeshmohan/dago/pkg/xfilesystem"
	"github.com/stretchr/testify/require"
)

func Test_recursiveCopy(t *testing.T) {
	t.Run("should copy files are folder", func(t *testing.T) {
		fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
			"sample.go":  {Data: []byte(`package main`)},
			"dir1/file1": {Data: []byte(`file1 content`)},
			"dir2/file2": {Data: []byte(`file2 content`)},
		})

		tempDir, err := os.MkdirTemp("", "test-temp")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)
		destFS := xfilesystem.NewRWFileSystem(tempDir, os.DirFS(tempDir))

		err = recursiveCopy(fsTest, destFS, map[string]interface{}{})
		require.NoError(t, err)

		resultFS := os.DirFS(tempDir)
		contents, err := fs.ReadDir(resultFS, ".")
		require.NoError(t, err)
		require.Equal(t, 3, len(contents))

		for _, c := range contents {
			require.Contains(t, []string{"dir1", "dir2", "sample.go"}, c.Name())
		}
	})
	t.Run("should recursively copy files and folders", func(t *testing.T) {
		fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
			"dir1/file1":      {Data: []byte(`file1 content`)},
			"dir1/dir2/file2": {Data: []byte(`file2 content`)},
		})

		tempDir, err := os.MkdirTemp("", "test-temp")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)
		destFS := xfilesystem.NewRWFileSystem(tempDir, os.DirFS(tempDir))

		err = recursiveCopy(fsTest, destFS, map[string]interface{}{})
		require.NoError(t, err)

		resultFS := os.DirFS(tempDir)
		contents, err := fs.ReadDir(resultFS, ".")
		require.NoError(t, err)
		require.Equal(t, 1, len(contents))

		require.Equal(t, "dir1", contents[0].Name())
		content, err := fs.ReadFile(resultFS, "dir1/file1")
		require.NoError(t, err)
		require.Equal(t, "file1 content", string(content))

		content, err = fs.ReadFile(resultFS, "dir1/dir2/file2")
		require.NoError(t, err)
		require.Equal(t, "file2 content", string(content))
	})
	t.Run("should ignore files from the ignore list", func(t *testing.T) {
		fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
			"dir1/file1":              {Data: []byte(`file1 content`)},
			"node_modules/sample.txt": {Data: []byte(`file2 content`)},
			"dago-template.yaml":      {Data: []byte(`file2 content`)},
		})

		tempDir, err := os.MkdirTemp("", "test-temp")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)
		destFS := xfilesystem.NewRWFileSystem(tempDir, os.DirFS(tempDir))

		err = recursiveCopy(fsTest, destFS, map[string]interface{}{})
		require.NoError(t, err)

		resultFS := os.DirFS(tempDir)
		contents, err := fs.ReadDir(resultFS, ".")
		require.NoError(t, err)
		require.Equal(t, 1, len(contents))

		require.Equal(t, "dir1", contents[0].Name())
		content, err := fs.ReadFile(resultFS, "dir1/file1")
		require.NoError(t, err)
		require.Equal(t, "file1 content", string(content))
	})
	t.Run("file and folder name can be template", func(t *testing.T) {
		fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
			"{{.project}}/file1":  {Data: []byte(`file1 content`)},
			"{{.file_name}}.yaml": {Data: []byte(`config content`)},
		})

		tempDir, err := os.MkdirTemp("", "test-temp")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)
		destFS := xfilesystem.NewRWFileSystem(tempDir, os.DirFS(tempDir))

		err = recursiveCopy(fsTest, destFS, map[string]interface{}{
			"project":   "server",
			"file_name": "config",
		})
		require.NoError(t, err)

		resultFS := os.DirFS(tempDir)
		contents, err := fs.ReadDir(resultFS, ".")
		require.NoError(t, err)
		require.Equal(t, 2, len(contents))

		content, err := fs.ReadFile(resultFS, "server/file1")
		require.NoError(t, err)
		require.Equal(t, "file1 content", string(content))

		content, err = fs.ReadFile(resultFS, "config.yaml")
		require.NoError(t, err)
		require.Equal(t, "config content", string(content))
	})
}
