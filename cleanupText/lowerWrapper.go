package cleanupText

import "strings"

type LowerWrapper struct {
	Text WrittenText
}

func (t *LowerWrapper) Render() string {
	return strings.ToLower(t.Text.Render())
}