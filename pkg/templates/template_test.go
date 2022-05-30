package templates

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTemplate(t *testing.T) {
	tempPath, err := os.MkdirTemp("", "dago-test")
	require.NoError(t, err)

	t.Run("should create a new template", func(t *testing.T) {
		_, err = CreateTemplate("sample", tempPath)
		require.NoError(t, err)

		fsInfo, err := os.Stat(filepath.Join(tempPath, "sample"))
		require.NoError(t, err)

		require.True(t, fsInfo.IsDir())
	})

	t.Run("should fail if the folder exist ", func(t *testing.T) {
		_, err = CreateTemplate("sample", tempPath)
		require.Error(t, err)
	})
}
