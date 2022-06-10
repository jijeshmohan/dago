package textrender

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/jijeshmohan/dago/pkg/xstring"
)

func RenderString(content string, data map[string]interface{}) (string, error) {
	if !strings.Contains(content, "{{") {
		return content, nil
	}

	t, err := createTemplate(content)
	if err != nil {
		return "", fmt.Errorf("unable to parse template %v", err)
	}

	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("render error :%v", err)
	}

	return buf.String(), nil
}

func RenderFile(srcPath, destPath string, data map[string]interface{}) error {
	templateContent, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	t, err := createTemplate(string(templateContent))
	if err != nil {
		return fmt.Errorf("unable to parse template file %s : %v", srcPath, err)
	}

	f, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("unable to create dest file %s: %v", destPath, err)
	}

	if err := t.Execute(f, data); err != nil {
		return fmt.Errorf("render error for %s: %v", srcPath, err)
	}

	return nil
}

func createTemplate(content string) (*template.Template, error) {
	return template.New("").Funcs(sprig.TxtFuncMap()).Funcs(xstring.FuncMap()).Parse(content)
}
