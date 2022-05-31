package xstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuncMap(t *testing.T) {
	strmap := FuncMap()

	pluralFn := strmap["pluralize"].(func(string) string)
	singularFn := strmap["singularize"].(func(string) string)

	t.Run("should return plural form", func(t *testing.T) {
		s := pluralFn("cat")
		require.Equal(t, "cats", s)
	})

	t.Run("shoudl return singular form", func(t *testing.T) {
		s := singularFn("cats")
		require.Equal(t, "cat", s)
	})
}
