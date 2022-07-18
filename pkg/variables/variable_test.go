package variables

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVariable_Validate(t *testing.T) {
	v := Variable{
		Name:    "name",
		Message: "message",
		Help:    "sample help",
		Default: "default",
		Type:    "text",
	}
	t.Run("should be valid", func(t *testing.T) {
		err := v.Validate()
		require.NoError(t, err)
	})
	t.Run("not valid without a name", func(t *testing.T) {
		vv := v
		vv.Name = ""
		err := vv.Validate()
		require.EqualError(t, err, "name can't be empty")
	})
	t.Run("name should be valid", func(t *testing.T) {
		vv := v
		vv.Name = "name.sample"
		err := vv.Validate()
		require.EqualError(t, err, "name should not contain . or -")

		vv.Name = "-namesample"
		err = vv.Validate()
		require.EqualError(t, err, "name should not contain . or -")
	})
	t.Run("name should be lowercase", func(t *testing.T) {
		vv := v
		vv.Name = "Name"
		err := vv.Validate()
		require.EqualError(t, err, "name should be lowercase")
	})
	t.Run("type should be valid", func(t *testing.T) {
		err := v.Validate()
		require.NoError(t, err)

		v.Type = "invalid"
		err = v.Validate()
		require.EqualError(t, err, "invalid variable type invalid")
	})
}

func TestVariables_Validate(t *testing.T) {
	v := Variables{
		{
			Name:    "name",
			Message: "message",
			Help:    "sample help",
			Default: "default",
			Type:    "text",
		},
	}
	t.Run("should be valid", func(t *testing.T) {
		err := v.Validate()
		require.NoError(t, err)
	})
	t.Run("duplicate variable names are not allowed", func(t *testing.T) {
		v = append(v, Variable{
			Name:    "name",
			Message: "message",
			Help:    "sample help",
			Default: "default",
			Type:    "text",
		})
		err := v.Validate()
		require.EqualError(t, err, "duplicate variable names are not allowed: name")
	})
	t.Run("should be valid variable", func(t *testing.T) {
		v = append(v[:1], Variable{
			Name:    "",
			Message: "message",
			Help:    "sample help",
			Default: "default",
			Type:    "text",
		})
		err := v.Validate()
		require.EqualError(t, err, "name can't be empty")
	})
}
