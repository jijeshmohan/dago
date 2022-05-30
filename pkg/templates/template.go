package templates

import (
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

const TemplateFile = `dago-template.yaml`

func CreateTemplate(name string, folderPath string) (Template, error) {
	folderPath = path.Join(folderPath, strings.ToLower(name))
	if isPathExist(folderPath) {
		return Template{}, fmt.Errorf("template already exist in the path: %s", folderPath)
	}

	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return Template{}, fmt.Errorf("failed to create template folder %v", err)
	}

	t := sampleTemplate(name)
	if err := t.save(folderPath); err != nil {
		return Template{}, err
	}

	return t, nil
}

type Template struct {
	Name      string   `yaml:"name,omitempty"`
	Variables []string `yaml:"variables,omitempty"`
}

func (t Template) save(folderPath string) error {
	data, err := yaml.Marshal(&t)
	if err != nil {
		return fmt.Errorf("invalid yaml data %v", err)
	}

	return os.WriteFile(path.Join(folderPath, TemplateFile), data, 0644)
}

func sampleTemplate(name string) Template {
	return Template{
		Name: name,
	}
}

func isPathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
