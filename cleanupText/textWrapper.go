package cleanupText

type TextWrapper struct {
	Text string
}

func (t *TextWrapper) Render() string {
	return t.Text
}