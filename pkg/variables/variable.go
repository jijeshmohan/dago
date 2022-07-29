package variables

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jijeshmohan/dago/pkg/textrender"
)

type VariablesGetter interface {
	GetValues(vars Variables) (map[string]interface{}, error)
}

// Variable is a representation of input variables in the template
type Variable struct {
	Name        string       `json:"name,omitempty"        yaml:"name"`
	Message     string       `json:"message,omitempty"     yaml:"message"`
	Help        string       `json:"help,omitempty"        yaml:"help"`
	Default     interface{}  `json:"default,omitempty"     yaml:"default,omitempty"`
	Type        VariableType `json:"type,omitempty"        yaml:"type"`
	Options     []string     `json:"options,omitempty"     yaml:"options,omitempty"`
	Validators  []string     `json:"validators,omitempty"  yaml:"validators,omitempty"`
	Transformer string       `json:"transformer,omitempty" yaml:"transformer,omitempty"`
}

//Validate variable
func (v Variable) Validate() error {
	if v.Name == "" {
		return errors.New("name can't be empty")
	}

	if v.Name != strings.ToLower(v.Name) {
		return errors.New("name should be lowercase")
	}

	if strings.Contains(v.Name, "-") || strings.Contains(v.Name, ".") {
		return errors.New("name should not contain . or -")
	}

	if err := v.Type.Validate(); err != nil {
		return err
	}

	return nil
}

type Variables []Variable

func (vars Variables) Validate() error {
	varMap := make(map[string]struct{})
	for _, v := range vars {
		if _, ok := varMap[v.Name]; ok {
			return fmt.Errorf("duplicate variable names are not allowed: %s", v.Name)
		}

		if err := v.Validate(); err != nil {
			return err
		}

		varMap[v.Name] = struct{}{}
	}

	return nil
}

func (vars Variables) populateTemplateValues(data map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{}, len(data))
	for k, v := range data {
		str, ok := v.(string)
		if !ok {
			result[k] = v
			continue
		}

		renderValue, err := textrender.RenderString(str, data)
		if err != nil {
			return nil, err
		}
		result[k] = renderValue
	}

	return result, nil
}
