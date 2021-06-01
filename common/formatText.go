package common

import (
	decorator "ClinicalIndication/cleanupText"
)

func FormatText(s string) string {
	wrapperText := &decorator.TextWrapper{Text: s}
	lowerText := &decorator.LowerWrapper{Text: wrapperText}
	punctuationText := &decorator.PunctuationWrapper{Text: lowerText}
	separatorText := &decorator.SeparatorWrapper{Text: punctuationText}
	stopWordsText := &decorator.StopWordsWrapper{Text: separatorText}
	//smallText := &decorator.SmallWrapper{Text: stopWordsText}
	fieldsText := &decorator.FieldsWrapper{Text: stopWordsText}
	return fieldsText.Render()
}