package xstring

import pluralize "github.com/gertd/go-pluralize"

var pluralizeClient *pluralize.Client

func init() {
	pluralizeClient = pluralize.NewClient()
}

func Plural(str string) string {
	return pluralizeClient.Plural(str)
}

func Singular(str string) string {
	return pluralizeClient.Singular(str)
}
