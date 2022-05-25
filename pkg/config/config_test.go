package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jijeshmohan/dago/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	tempPath, err := os.MkdirTemp("", "dago-test")
	require.NoError(t, err)

	t.Run("should create a new config file in the given path", func(t *testing.T) {
		err := config.Create(tempPath)
		require.NoError(t, err)

		_, err = os.Stat(filepath.Join(tempPath, "dago.yaml"))
		require.NoError(t, err)
	})
	t.Run("should return error if the file is already exist", func(t *testing.T) {
		err := config.Create(tempPath)
		require.Error(t, err)
	})
	t.Run("should work with user path", func(t *testing.T) {
		tempPath := "~/.config/dago-test"
		err := os.Mkdir(getUserHomeDirectory(t)+"/.config/dago-test", 0755)
		require.NoError(t, err)
		defer os.RemoveAll(getUserHomeDirectory(t) + "/.config/dago-test")

		err = config.Create(tempPath)
		require.NoError(t, err)

		_, err = os.Stat(filepath.Join(getUserHomeDirectory(t)+"/.config/dago-test", "dago.yaml"))
		require.NoError(t, err)
	})
	t.Run("should create directory if it does not exist", func(t *testing.T) {
		tempPath := "/tmp/.config/dago-test-temp"
		defer os.RemoveAll(tempPath)

		err = config.Create(tempPath)
		require.NoError(t, err)

		_, err = os.Stat(filepath.Join(tempPath, "dago.yaml"))
		require.NoError(t, err)
	})
}

func TestLoad(t *testing.T) {
	tempPath, err := os.MkdirTemp("", "dago-test")
	require.NoError(t, err)

	err = config.Create(tempPath)
	require.NoError(t, err)

	t.Run("should be able to load config from path", func(t *testing.T) {
		config, err := config.Load(tempPath)
		require.NoError(t, err)
		require.Equal(t, filepath.Join(tempPath, "templates"), config.TemplatesPath)
	})
	t.Run("should return error if file does not exists", func(t *testing.T) {
		tempPath, err := os.MkdirTemp("", "dago-test-another")
		require.NoError(t, err)

		_, err = config.Load(tempPath)
		require.Error(t, err)
	})

}

func getUserHomeDirectory(t *testing.T) string {
	userDir, err := os.UserHomeDir()
	require.NoError(t, err)

	return userDir
}
