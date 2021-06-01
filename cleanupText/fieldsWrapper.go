package cleanupText

import (
	"strings"
)

type FieldsWrapper struct {
	Text WrittenText
}

func (t *FieldsWrapper) Render() string {
	return strings.Join(strings.Fields(t.Text.Render()), " ")
}