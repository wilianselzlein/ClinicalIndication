package cleanupText

import (
	"regexp"
	"log"
)

type PunctuationWrapper struct {
	Text WrittenText
}

func (t *PunctuationWrapper) Render() string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	return re.ReplaceAllString(t.Text.Render(), " ")
}