package textrender

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

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
	tempDir, err := ioutil.TempDir("", "test-temp")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	tempFile := generateFile(t, "sample.go", []byte(`package {{.name}}\n// {{.author}}`))

	t.Run("should render file with template content", func(t *testing.T) {
		err := RenderFile(tempFile, path.Join(tempDir, "result.go"), map[string]interface{}{
			"name":   "sample",
			"author": "user1",
		})
		require.NoError(t, err)

		content, err := ioutil.ReadFile(path.Join(tempDir, "result.go"))
		require.NoError(t, err)
		require.Equal(t, []byte(`package sample\n// user1`), content)
	})
	t.Run("should render with no value tag if the data is not preset", func(t *testing.T) {
		err := RenderFile(tempFile, path.Join(tempDir, "result_no.go"), map[string]interface{}{})
		require.NoError(t, err)

		content, err := ioutil.ReadFile(path.Join(tempDir, "result_no.go"))
		require.NoError(t, err)
		require.Equal(t, `package <no value>\n// <no value>`, string(content))
	})
}

func generateFile(t *testing.T, pattern string, content []byte) string {
	f, err := ioutil.TempFile("", pattern)
	require.NoError(t, err)
	defer f.Close()

	_, err = f.Write(content)
	require.NoError(t, err)

	return f.Name()
}
