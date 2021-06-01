package cleanupText

import (
	"github.com/bbalet/stopwords"
)

type StopWordsWrapper struct {
	Text WrittenText
}

func (t *StopWordsWrapper) Render() string {
	stopwords.DontStripDigits()
	return stopwords.CleanString(t.Text.Render(), "pt", false)
}

