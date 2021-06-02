package cleanupText

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type NormalizeWrapper struct {
	Text WrittenText
}

var normalizer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func normalize(str string) (string, error) {
	s, _, err := transform.String(normalizer, str)
	if err != nil {
		return "", err
	}
	return s, err
}

func (t *NormalizeWrapper) Render() string {
	res, _ := normalize(t.Text.Render())
	return res
}

