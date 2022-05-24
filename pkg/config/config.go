package config

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func NewConfig(configPath string) error {
	configPath = expandUserPath(configPath)

	config := Config{
		TemplatesPath: path.Join(configPath, "/templates"),
	}

	if isConfigFileExist(configPath) {
		return fmt.Errorf("config file already exist in the path: %s", configPath)
	}

	return config.save(configPath)
}

func Load(configPath string) (Config, error) {
	configPath = expandUserPath(configPath)

	if !isConfigFileExist(configPath) {
		return Config{}, fmt.Errorf("config file is not exist in the path: %s", configPath)
	}

	data, err := os.ReadFile(configFilePath(configPath))
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return config, nil
}

type Config struct {
	TemplatesPath string `yaml:"templates-path"`
}

func (c *Config) save(configPath string) error {
	d, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshalling config: %v", err)
	}

	return writeFile(configFilePath(configPath), d)
}

func writeFile(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating config file: %v", err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}

func configFilePath(configPath string) string {
	return path.Join(configPath, "dago.yaml")
}

func isConfigFileExist(configPath string) bool {
	_, err := os.Stat(configFilePath(configPath))
	return err == nil
}

func expandUserPath(folderPath string) string {
	if !strings.HasPrefix(folderPath, "~/") {
		return folderPath
	}

	usr, err := user.Current()
	if err != nil {
		panic("unable to get current user directory :" + err.Error())
	}

	dir := usr.HomeDir
	folderPath = filepath.Join(dir, folderPath[2:])
	return folderPath
}
