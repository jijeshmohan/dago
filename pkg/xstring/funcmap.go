package xstring

var templateMap = map[string]interface{}{
	"pluralize":   Plural,
	"singularize": Singular,
}

func FuncMap() map[string]interface{} {
	return templateMap
}
