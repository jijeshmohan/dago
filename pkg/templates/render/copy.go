package render

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/jijeshmohan/dago/pkg/textrender"
	"github.com/jijeshmohan/dago/pkg/xfilesystem"
)

var ignoreFiles = []string{".git", "node_modules", "dago-template.yaml"}

func shouldIgnore(baseName string) bool {
	for _, f := range ignoreFiles {
		if f == baseName {
			return true
		}
	}

	return false
}

func recursiveCopy(srcFS xfilesystem.FS, destFS xfilesystem.RwFS, data map[string]interface{}) error {
	content, err := fs.ReadDir(srcFS, ".")
	if err != nil {
		return fmt.Errorf("unable to read dir %s: %v", srcFS.BasePath(), err)
	}

	for _, file := range content {
		baseName := file.Name()
		if shouldIgnore(baseName) {
			continue
		}

		renderName, err := textrender.RenderString(baseName, data)
		if err != nil {
			return err
		}

		if strings.TrimSpace(renderName) == "" {
			continue
		}

		if !file.IsDir() {
			if err := textrender.RenderFile(srcFS, baseName, destFS, renderName, data); err != nil {
				return err
			}

			continue
		}

		if err := destFS.Mkdir(renderName); err != nil {
			return err
		}

		subSrc, err := srcFS.Sub(baseName)
		if err != nil {
			return err
		}

		subDest, err := destFS.Sub(renderName)
		if err != nil {
			return err
		}

		if err := recursiveCopy(subSrc, subDest, data); err != nil {
			return err
		}
	}

	return nil
}
