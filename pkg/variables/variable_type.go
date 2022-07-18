package variables

import "errors"

type VariableType string

var (
	TextType        VariableType = "text"
	ConfirmType     VariableType = "confirm"
	MultiLineType   VariableType = "multi_line"
	SelectType      VariableType = "select"
	MultiSelectType VariableType = "multi_select"
)

func (vt VariableType) Validate() error {
	switch vt {
	case TextType, MultiLineType, SelectType, MultiSelectType, ConfirmType:
		return nil
	}

	return errors.New("invalid variable type " + string(vt))
}
