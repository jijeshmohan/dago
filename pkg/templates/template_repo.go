package templates

import (
	"errors"
	"fmt"
	"io/fs"
	"strings"

	"github.com/jijeshmohan/dago/pkg/xfilesystem"
)

type Repository interface {
	GetTemplate(templateName string) (Template, error)
	GetAllTemplateNames() []string
}

type fileSystemRepository struct {
	fs        xfilesystem.FS
	templates map[string]Template
}

func NewFSRepository(filesystem xfilesystem.FS) (Repository, error) {
	templates, err := loadTemplates(filesystem)
	if err != nil {
		return nil, err
	}

	return &fileSystemRepository{
		templates: templates,
		fs:        filesystem,
	}, nil
}

func (r *fileSystemRepository) GetTemplate(templateName string) (Template, error) {
	t, ok := r.templates[strings.ToLower(templateName)]
	if !ok {
		return Template{}, errors.New("template not found")
	}

	return t, nil
}

func (r *fileSystemRepository) GetAllTemplateNames() []string {
	results := make([]string, 0, len(r.templates))
	for name := range r.templates {
		results = append(results, name)
	}

	return results
}

func loadTemplates(fsys xfilesystem.FS) (map[string]Template, error) {
	content, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, fmt.Errorf("unable to read template dir %v", err)
	}

	templatesMap := make(map[string]Template)
	for _, c := range content {
		if !c.IsDir() {
			continue
		}

		templateDir, err := fsys.Sub(c.Name())
		if err != nil {
			continue
		}

		template, err := loadTemplate(templateDir)
		if err != nil {
			fmt.Printf("Error while loading template %s: %v\n", c.Name(), err)
			continue
		}

		templatesMap[template.Name] = template
	}

	return templatesMap, nil
}
