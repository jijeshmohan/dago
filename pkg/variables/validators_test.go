package variables

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInteger(t *testing.T) {
	t.Run("should be valid for integer value", func(t *testing.T) {
		var x interface{} = 1
		err := Integer(x)
		require.NoError(t, err)
	})
	t.Run("should return error for float", func(t *testing.T) {
		var x interface{} = 1.3
		err := Integer(x)
		require.Error(t, err)
	})
	t.Run("should return error for string", func(t *testing.T) {
		var x interface{} = "sample"
		err := Integer(x)
		require.Error(t, err)
	})
}
