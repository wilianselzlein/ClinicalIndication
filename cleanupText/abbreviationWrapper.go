package cleanupText

import (
	//"strings"
	"regexp"
)

var abbreviations = map[string]string{
	"oe": "olho esquerdo", 
	"esq": "esquerdo", 
	"od": "olho direito", 
	"dir": "direito",
}

type AbbreviationWrapper struct {
	Text WrittenText
}

func (t *AbbreviationWrapper) Render() string {
	var result string = t.Text.Render()
	for k, v := range abbreviations {
		v = " " + v + " "
		var re = regexp.MustCompile(" " + k + " ") 
		result = re.ReplaceAllString(result, v)
		re = regexp.MustCompile("^" + k + " ")
		result = re.ReplaceAllString(result, v)
		re = regexp.MustCompile("^" + k + "$")
		result = re.ReplaceAllString(result, v)
		re = regexp.MustCompile(" " + k + "$")
		result = re.ReplaceAllString(result, v)
	}

	return result
}