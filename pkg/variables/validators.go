package variables

import (
	"errors"
)

// Integer allow only integer value
func Integer(val interface{}) error {
	if _, ok := val.(int); !ok {
		return errors.New("value should be an integer")
	}

	return nil
}
