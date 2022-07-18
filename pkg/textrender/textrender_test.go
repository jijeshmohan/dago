package textrender

import (
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"github.com/jijeshmohan/dago/pkg/xfilesystem"
	"github.com/stretchr/testify/require"
)

func TestRenderString(t *testing.T) {
	t.Run("should ignore render if it is plain text", func(t *testing.T) {
		text := "this is very simple text"
		result, err := RenderString(text, nil)

		require.NoError(t, err)
		require.Equal(t, text, result)
	})
	t.Run("should fail if the template is not correct", func(t *testing.T) {
		text := "this is sample {{ "
		_, err := RenderString(text, nil)

		require.Error(t, err)
	})
	t.Run("should render content", func(t *testing.T) {
		text := "this is sample {{.name}}"
		result, err := RenderString(text, map[string]interface{}{
			"name": "Blog",
		})

		require.NoError(t, err)
		require.Equal(t, "this is sample Blog", result)
	})
}

func TestConditionalTemplate(t *testing.T) {
	text := `{{if eq .db "pg"}}yes{{end}}`
	result, err := RenderString(text, map[string]interface{}{
		"db": "pg",
	})

	require.NoError(t, err)
	require.Equal(t, "yes", result)

	text = `{{if .db}}yes{{end}}`
	result, err = RenderString(text, nil)

	require.NoError(t, err)
	require.Equal(t, "", result)

	text = `{{if .db}}yes{{end}}`
	result, err = RenderString(text, map[string]interface{}{
		"db": "",
	})

	require.NoError(t, err)
	require.Equal(t, "", result)
}

func TestRenderFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test-temp")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	destFS := xfilesystem.NewRWFileSystem(tempDir, os.DirFS(tempDir))

	fsTest := xfilesystem.NewFileSystem("/tmp", fstest.MapFS{
		"sample.go": {Data: []byte(`package {{.name}}\n// {{.author}}`)},
	})

	t.Run("should render file with template content", func(t *testing.T) {
		err := RenderFile(fsTest, "sample.go", destFS, "result.go", map[string]interface{}{
			"name":   "sample",
			"author": "user1",
		})
		require.NoError(t, err)

		content, err := fs.ReadFile(destFS, "result.go")
		require.NoError(t, err)
		require.Equal(t, []byte(`package sample\n// user1`), content)
	})
	t.Run("should render with no value tag if the data is not preset", func(t *testing.T) {
		err := RenderFile(fsTest, "sample.go", destFS, "result_no.go", map[string]interface{}{})
		require.NoError(t, err)

		content, err := fs.ReadFile(destFS, "result_no.go")
		require.NoError(t, err)
		require.Equal(t, `package <no value>\n// <no value>`, string(content))
	})
}
