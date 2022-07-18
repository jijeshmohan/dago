package templates

import (
	"fmt"
	"os"
	"path"
	"strings"
)

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
		if removeErr := removeTemplate(folderPath); removeErr != nil {
			fmt.Printf("unable to clean the directory %v\n", removeErr)
		}

		return Template{}, err
	}

	return t, nil
}

func removeTemplate(folderPath string) error {
	return os.RemoveAll(folderPath)
}

func sampleTemplate(name string) Template {
	return Template{
		Name: name,
	}
}
