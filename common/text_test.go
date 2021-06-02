package common

import "testing"

func TestFormatedTextSimple(t *testing.T) {
	text := "ACHÉ 1 12 123 oe esq a aa aaa + de pára onde ônd inter \n [CLORIDRATO] (acido) {äcida}, internationalization ."
	want := "ache 1 12 123 olho esquerdo esquerdo aa aaa ond inter cloridrato acido acida internationalization"
	res := FormatText(text)

	if  res != want {
		t.Errorf("FormatText(%s) = %s; want %s", text, res, want)
	}
}
