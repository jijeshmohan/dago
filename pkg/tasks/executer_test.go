package tasks

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jijeshmohan/dago/pkg/xlogger"
	"github.com/stretchr/testify/require"
)

func TestTaskExecuter_ExecuteTasks(t *testing.T) {
	taskExecuter := NewTaskExecuter(xlogger.NewEmptyLogger())

	tempDir, err := os.MkdirTemp("", "test-temp")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	t.Run("should execute all the tasks", func(t *testing.T) {
		tasks := []Task{
			{
				Command:     "mkdir",
				Arguments:   []string{"sample"},
				Path:        ".",
				IgnoreError: false,
			},
		}
		err := taskExecuter.ExecuteTasks(tasks, map[string]interface{}{}, tempDir)
		require.NoError(t, err)

		found, err := exists(filepath.Join(tempDir, "sample"))
		require.NoError(t, err)

		require.True(t, found)
	})
	t.Run("should render arguments", func(t *testing.T) {
		tasks := []Task{
			{
				Command:     "mkdir",
				Arguments:   []string{"{{.name}}"},
				Path:        ".",
				IgnoreError: false,
			},
		}
		err := taskExecuter.ExecuteTasks(tasks, map[string]interface{}{
			"name": "test",
		}, tempDir)
		require.NoError(t, err)

		found, err := exists(filepath.Join(tempDir, "test"))
		require.NoError(t, err)

		require.True(t, found)
	})
	t.Run("should not proceed if there is an error in the first command", func(t *testing.T) {
		tasks := []Task{
			{
				Command:     "mkdir",
				Arguments:   []string{"-xyz", "another"},
				Path:        ".",
				IgnoreError: false,
			},
		}
		err := taskExecuter.ExecuteTasks(tasks, map[string]interface{}{}, tempDir)
		require.Error(t, err)

		found, err := exists(filepath.Join(tempDir, "another"))
		require.NoError(t, err)

		require.False(t, found)
	})
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
