package cleanupText

import (
	"strings"
)

var abbreviations = map[string]string{
	"oe": "olho esquerdo", 
	"esq": "esquerdo", 
	"od": "olho direito", 
	"dir": "direito",
	" ": "", 
}

type AbbreviationWrapper struct {
	Text WrittenText
}

func (t *AbbreviationWrapper) Render() string {
	var result string = t.Text.Render()
	for k, v := range abbreviations {
		result = strings.Replace(result, " " + k + " ", " " + v + " ", -1)
	}
	return result
}