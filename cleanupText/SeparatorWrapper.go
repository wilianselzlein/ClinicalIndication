package cleanupText

import (
	"regexp"
	"strings"
)

type SeparatorWrapper struct {
	Text WrittenText
}

func (t *SeparatorWrapper) Render() string {
	var result []string
	re := regexp.MustCompile(`(\d+)`)
	for _, word := range strings.Fields(t.Text.Render()) {
		result = append(result, re.ReplaceAllString(word, " $1 "))

	}
	return strings.Join(result, " ")
}