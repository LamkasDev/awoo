package lexer

import "unicode"

type ConstructChunkValidator func(rune) bool

func ConstructChunkSkipperDefault(r rune) bool { return false }

func ConstructChunk(lexer *AwooLexer, cs string, skip ConstructChunkValidator, validate ConstructChunkValidator) string {
	for _, err := AdvanceLexer(lexer); err == nil; _, err = AdvanceLexer(lexer) {
		if skip(unicode.ToLower(lexer.Current.Character)) {
			continue
		}
		if !validate(unicode.ToLower(lexer.Current.Character)) {
			break
		}
		cs += (string)(lexer.Current.Character)
	}
	StepbackLexer(lexer)

	return cs
}

func ConstructChunkFast(lexer *AwooLexer, cs string, validate ConstructChunkValidator) string {
	return ConstructChunk(lexer, cs, ConstructChunkSkipperDefault, validate)
}
