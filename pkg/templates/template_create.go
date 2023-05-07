package templates

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jijeshmohan/dago/pkg/variables"
)

func CreateTemplate(name string, folderPath string) (Template, error) {
	folderPath = path.Join(folderPath, strings.ToLower(name))
	if isPathExist(folderPath) {
		return Template{}, fmt.Errorf("template already exist in the path: %s", folderPath)
	}

	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return Template{}, fmt.Errorf("failed to create template folder %v", err)
	}

	variables, err := getVariablesForTemplate()
	if err != nil {
		return Template{}, err
	}

	t := sampleTemplate(name)
	t.Variables = variables
	if err := t.save(folderPath); err != nil {
		if removeErr := removeTemplate(folderPath); removeErr != nil {
			fmt.Printf("unable to clean the directory %v\n", removeErr)
		}

		return Template{}, err
	}

	return t, nil
}

func getVariablesForTemplate() ([]variables.Variable, error) {
	variables := make([]variables.Variable, 0)
	for {
		variableType := &survey.Confirm{
			Message: "Do you want to add a variable?",
			Default: false,
			Help:    "If yes, a new variable will be added to the template with followup questions",
		}

		var ok bool
		if err := survey.AskOne(variableType, &ok); err != nil {
			return variables, err
		}

		if !ok {
			break
		}

		variable, err := getVariable()
		if err != nil {
			return variables, err
		}

		variables = append(variables, variable)
	}

	return variables, nil
}

func getVariable() (variables.Variable, error) {
	variableTypeQ := &survey.Select{
		Message: "Select the variable type",
		Options: []string{
			"text",
			"confirm",
			"multi_line",
			"select",
			"multi_select",
		},
		Default: "text",
	}

	var variableTypeStr string
	if err := survey.AskOne(variableTypeQ, &variableTypeStr); err != nil {
		return variables.Variable{}, err
	}

	variableNameQ := &survey.Input{
		Message: "Enter the variable name",
	}
	var variableName string
	if err := survey.AskOne(variableNameQ, &variableName, survey.WithValidator(survey.Required), survey.WithValidator(variableNameValidator)); err != nil {
		return variables.Variable{}, err
	}

	variableMessageQ := &survey.Input{
		Message: "Enter the variable message",
	}
	var variableMessage string
	if err := survey.AskOne(variableMessageQ, &variableMessage); err != nil {
		return variables.Variable{}, err
	}

	variableHelpQ := &survey.Input{
		Message: "Enter the variable help",
	}
	var variableHelp string
	if err := survey.AskOne(variableHelpQ, &variableHelp); err != nil {
		return variables.Variable{}, err
	}

	variable := variables.Variable{
		Name:    variableName,
		Type:    variables.VariableType(variableTypeStr),
		Message: variableMessage,
		Help:    variableHelp,
	}

	if variableTypeStr == "select" || variableTypeStr == "multi_select" {
		variableOptionsQ := &survey.Input{
			Message: "Enter the variable options (comma separated)",
		}
		var variableOptions string
		if err := survey.AskOne(variableOptionsQ, &variableOptions); err != nil {
			return variables.Variable{}, err
		}

		variableOptions = strings.TrimSpace(variableOptions)
		if variableOptions == "" {
			return variables.Variable{}, fmt.Errorf("variable options cannot be empty")
		}

		variableOptionsArr := strings.Split(variableOptions, ",")
		for i, option := range variableOptionsArr {
			variableOptionsArr[i] = strings.TrimSpace(option)
		}
		variable.Options = variableOptionsArr
	}

	return variable, nil
}

func removeTemplate(folderPath string) error {
	return os.RemoveAll(folderPath)
}

func sampleTemplate(name string) Template {
	return Template{
		Name: name,
	}
}

func variableNameValidator(val interface{}) error {
	if str, ok := val.(string); !ok || (strings.Contains(str, "-") || strings.Contains(str, ".")) {
		return errors.New("name should not contain . or -")
	}

	return nil
}
