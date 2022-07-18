package variables

import (
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/stretchr/testify/require"
)

func TestConsoleGetter_createQuestionsFromVariables(t *testing.T) {
	t.Run("should create survey questions from the variables", func(t *testing.T) {
		vars := Variables{
			Variable{
				Name:    "project_name",
				Message: "enter project name",
				Help:    "Please enter project name",
				Type:    TextType,
			},
			Variable{
				Name:    "project_name_with_default",
				Message: "enter project name",
				Type:    TextType,
				Default: "default_value",
			},
			Variable{
				Name:    "multi_line_text",
				Message: "enter multi line test",
				Help:    "Please enter multi line test",
				Type:    MultiLineType,
				Default: "default_value",
			},
			Variable{
				Name:    "select_name",
				Message: "select one of the option",
				Help:    "Please select ",
				Type:    SelectType,
				Options: []string{"option1", "option2", "option3"},
				Default: "option2",
			},
			Variable{
				Name:    "multi_select_name",
				Message: "select options",
				Help:    "Please select options",
				Type:    MultiSelectType,
				Options: []string{"option1", "option2", "option3"},
				Default: "option2",
			},
			Variable{
				Name:    "confirm-type",
				Message: "confirm",
				Help:    "Can you confirm?",
				Type:    ConfirmType,
				Default: true,
			},
		}

		expectedQuestions := []*survey.Question{
			{
				Name: vars[0].Name,
				Prompt: &survey.Input{
					Message: vars[0].Message,
					Help:    vars[0].Help,
				},
			},
			{
				Name: vars[1].Name,
				Prompt: &survey.Input{
					Message: vars[1].Message,
					Help:    vars[1].Help,
					Default: vars[1].Default.(string),
				},
			},
			{
				Name: vars[2].Name,
				Prompt: &survey.Multiline{
					Message: vars[2].Message,
					Help:    vars[2].Help,
					Default: vars[2].Default.(string),
				},
			},
			{
				Name: vars[3].Name,
				Prompt: &survey.Select{
					Message: vars[3].Message,
					Help:    vars[3].Help,
					Options: vars[3].Options,
					Default: vars[3].Default.(string),
				},
			},
			{
				Name: vars[4].Name,
				Prompt: &survey.MultiSelect{
					Message: vars[4].Message,
					Help:    vars[4].Help,
					Options: vars[4].Options,
					Default: vars[4].Default.(string),
				},
			},
			{
				Name: vars[5].Name,
				Prompt: &survey.Confirm{
					Message: vars[5].Message,
					Help:    vars[5].Help,
					Default: true,
				},
			},
		}

		questions, err := NewConsoleGetter().createQuestionsFromVariables(vars)

		require.NoError(t, err)
		require.Equal(t, len(vars), len(questions))
		require.EqualValues(t, expectedQuestions, questions)
	})
	t.Run("should return error for invalid variable", func(t *testing.T) {
		_, err := NewConsoleGetter().createQuestionsFromVariables(Variables{{
			Name:    "invalid_variable",
			Message: "enter project name",
		}})

		require.Error(t, err)
	})
	t.Run("Validators", func(t *testing.T) {
		t.Run("should enter validator", func(t *testing.T) {
			questions, err := NewConsoleGetter().createQuestionsFromVariables(Variables{{
				Name:       "project_name",
				Message:    "enter project name",
				Help:       "Please enter project name",
				Type:       TextType,
				Validators: []string{"required"},
			}})
			require.NoError(t, err)
			require.Equal(t, 1, len(questions))
			require.NotNil(t, questions[0].Validate)

			require.NotPanics(t, func() { survey.ComposeValidators(questions[0].Validate) })
		})
		t.Run("should allow multiple validators", func(t *testing.T) {
			questions, err := NewConsoleGetter().createQuestionsFromVariables(Variables{{
				Name:       "project_name",
				Message:    "enter project name",
				Help:       "Please enter project name",
				Type:       TextType,
				Validators: []string{"required", "integer"},
			}})
			require.NoError(t, err)
			require.Equal(t, 1, len(questions))

			require.NotPanics(t, func() { survey.ComposeValidators(questions[0].Validate) })
		})
		t.Run("return error for invalid validator", func(t *testing.T) {
			_, err := NewConsoleGetter().createQuestionsFromVariables(Variables{{
				Name:       "project_name",
				Message:    "enter project name",
				Help:       "Please enter project name",
				Type:       TextType,
				Validators: []string{"invalid_validator"},
			}})

			require.ErrorContains(t, err, "invalid validator")
		})
	})
	t.Run("transformer", func(t *testing.T) {
		t.Run("should return transformer ", func(t *testing.T) {
			questions, err := NewConsoleGetter().createQuestionsFromVariables(Variables{{
				Name:        "project_name",
				Message:     "enter project name",
				Help:        "Please enter project name",
				Type:        TextType,
				Transformer: "lower",
			}})
			require.NoError(t, err)
			require.Equal(t, 1, len(questions))
			require.NotNil(t, questions[0].Transform)
		})
		t.Run("should return error for unknown transformer", func(t *testing.T) {
			_, err := NewConsoleGetter().createQuestionsFromVariables(Variables{{
				Name:        "project_name",
				Message:     "enter project name",
				Help:        "Please enter project name",
				Type:        TextType,
				Transformer: "unknown",
			}})
			require.ErrorContains(t, err, "invalid transformer")
		})
	})
}
