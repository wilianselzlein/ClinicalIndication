package cleanupText

import "strings"

type SmallWrapper struct {
	Text WrittenText
}

func (t *SmallWrapper) Render() string {
	var ret []string
	for _, word := range strings.Fields(t.Text.Render()) {
		if len(word) > 1 { 
			ret = append(ret, word)
		}
	}
	return strings.Join(ret, " ")
}