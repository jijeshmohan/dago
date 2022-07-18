package templates

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/jijeshmohan/dago/pkg/tasks"
	"github.com/jijeshmohan/dago/pkg/variables"
	"github.com/jijeshmohan/dago/pkg/xfilesystem"
	"gopkg.in/yaml.v3"
)

const TemplateFile = `dago-template.yaml`

type Template struct {
	Name      string              `yaml:"name,omitempty"`
	Variables variables.Variables `yaml:"variables,omitempty"`
	Tasks     []tasks.Task        `yaml:"tasks,omitempty"`

	fs xfilesystem.FS `yaml:"-"`
}

func (t Template) Filesystem() xfilesystem.FS {
	return t.fs
}

func (t Template) save(folderPath string) error {
	data, err := yaml.Marshal(&t)
	if err != nil {
		return fmt.Errorf("invalid yaml data %w", err)
	}

	return os.WriteFile(path.Join(folderPath, TemplateFile), data, 0644)
}

func loadTemplate(fsys xfilesystem.FS) (Template, error) {
	data, err := fs.ReadFile(fsys, TemplateFile)
	if err != nil {
		return Template{}, fmt.Errorf("unable to find template in the path %w", err)
	}

	var t Template
	if err := yaml.Unmarshal([]byte(data), &t); err != nil {
		return Template{}, fmt.Errorf("invalid template yml: %w", err)
	}

	for _, v := range t.Variables {
		if err := v.Validate(); err != nil {
			return Template{}, err
		}
	}

	t.fs = fsys

	return t, nil
}

func isPathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
