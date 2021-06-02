package common

import (
	decorator "ClinicalIndication/cleanupText"
)

func FormatText(s string) string {
	wrapperText := &decorator.TextWrapper{Text: s}
	lowerText := &decorator.LowerWrapper{Text: wrapperText}
	normalizeText := &decorator.NormalizeWrapper{Text: lowerText}
	punctuationText := &decorator.PunctuationWrapper{Text: normalizeText}
	separatorText := &decorator.SeparatorWrapper{Text: punctuationText}
	abbreviationsText := &decorator.AbbreviationWrapper{Text: separatorText}
	stopWordsText := &decorator.StopWordsWrapper{Text: abbreviationsText}
	//smallText := &decorator.SmallWrapper{Text: stopWordsText}
	fieldsText := &decorator.FieldsWrapper{Text: stopWordsText}
	return fieldsText.Render()
}