package variables

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jijeshmohan/dago/pkg/xstring"
)

type ConsoleGetter struct {
	validatorMap   map[string]survey.Validator
	transformerMap map[string]survey.Transformer
}

func NewConsoleGetter() ConsoleGetter {
	return ConsoleGetter{
		validatorMap: map[string]survey.Validator{
			"required": survey.Required,
			"integer":  Integer,
		},
		transformerMap: map[string]survey.Transformer{
			"lower":    survey.ToLower,
			"upper":    survey.Title,
			"plural":   plural,
			"singular": singular,
		},
	}
}

func (c ConsoleGetter) GetValues(vars Variables) (answers map[string]interface{}, err error) {
	questions, err := c.createQuestionsFromVariables(vars)
	if err != nil {
		return nil, err
	}

	answers = make(map[string]interface{})
	if err := survey.Ask(questions, &answers); err != nil {
		return nil, err
	}

	answers, err = vars.populateTemplateValues(answers)
	if err != nil {
		return nil, fmt.Errorf("unable to populate template values %w", err)
	}

	return answers, nil
}

func (c ConsoleGetter) createQuestionsFromVariables(vars Variables) ([]*survey.Question, error) {
	questions := make([]*survey.Question, len(vars))
	for i, v := range vars {
		prompt, err := createPrompt(v)
		if err != nil {
			return nil, fmt.Errorf("unable to construct question for variable %s: %w", v.Name, err)
		}

		validators, err := c.getValidators(v.Validators)
		if err != nil {
			return nil, fmt.Errorf("invalid validator for %s: %w", v.Name, err)
		}

		transformer, err := c.getTransformer(v.Transformer)
		if err != nil {
			return nil, fmt.Errorf("invalid transformer for %s: %w", v.Name, err)
		}

		questions[i] = &survey.Question{
			Name:      v.Name,
			Prompt:    prompt,
			Validate:  validators,
			Transform: transformer,
		}
	}

	return questions, nil
}

func createPrompt(v Variable) (survey.Prompt, error) {
	switch v.Type {
	case TextType:
		return &survey.Input{
			Message: v.Message,
			Default: getDefaultString(v.Default),
			Help:    v.Help,
		}, nil

	case MultiLineType:
		return &survey.Multiline{
			Message: v.Message,
			Default: getDefaultString(v.Default),
			Help:    v.Help,
		}, nil

	case SelectType:
		return &survey.Select{
			Message: v.Message,
			Options: v.Options,
			Default: v.Default,
			Help:    v.Help,
		}, nil

	case MultiSelectType:
		return &survey.MultiSelect{
			Message: v.Message,
			Options: v.Options,
			Default: v.Default,
			Help:    v.Help,
		}, nil

	case ConfirmType:
		return &survey.Confirm{
			Message: v.Message,
			Help:    v.Help,
			Default: getDefaultBool(v.Default),
		}, nil
	}

	return nil, errors.New("prompt: unsupported variable type")
}

func getDefaultString(defaultValue interface{}) string {
	value, ok := defaultValue.(string)
	if !ok {
		return ""
	}

	return value
}

func getDefaultBool(defaultValue interface{}) bool {
	value, ok := defaultValue.(bool)
	if !ok {
		return false
	}

	return value
}

func (c ConsoleGetter) getValidators(validatorStr []string) (survey.Validator, error) {
	if validatorStr == nil {
		return nil, nil
	}

	if len(validatorStr) == 1 {
		validatorFunc, ok := c.validatorMap[validatorStr[0]]
		if !ok {
			return nil, fmt.Errorf("invalid validator %s", validatorStr[0])
		}

		return validatorFunc, nil
	}

	validators := make([]survey.Validator, len(validatorStr))
	for _, validator := range validatorStr {
		validatorFunc, ok := c.validatorMap[validator]
		if !ok {
			return nil, fmt.Errorf("invalid validator %s", validatorStr[0])
		}

		validators = append(validators, validatorFunc)
	}

	return survey.ComposeValidators(validators...), nil
}

func (c ConsoleGetter) getTransformer(transformerStr string) (survey.Transformer, error) {
	if transformerStr == "" {
		return nil, nil
	}

	transformerFunc, ok := c.transformerMap[transformerStr]
	if !ok {
		return nil, fmt.Errorf("invalid transformer %s", transformerStr)
	}

	return transformerFunc, nil
}

func plural(ans interface{}) interface{} {
	return xstring.Plural(ans.(string))
}

func singular(ans interface{}) interface{} {
	return xstring.Singular(ans.(string))
}
