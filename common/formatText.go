package common

import (
	decorator "ClinicalIndication/cleanupText"
)

func FormatText(s string) string {
	var wrapper decorator.WrittenText

	wrapper = &decorator.TextWrapper{Text: s}
	wrapper = &decorator.LowerWrapper{Text: wrapper}
	wrapper = &decorator.NormalizeWrapper{Text: wrapper}
	wrapper = &decorator.PunctuationWrapper{Text: wrapper}
	wrapper = &decorator.SeparatorWrapper{Text: wrapper}
	wrapper = &decorator.AbbreviationWrapper{Text: wrapper}
	wrapper = &decorator.StopWordsWrapper{Text: wrapper}
	//wrapper = &decorator.SmallWrapper{Text: wrapper}
	wrapper = &decorator.FieldsWrapper{Text: wrapper}
	return wrapper.Render()
}