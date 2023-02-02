package lexer

import (
	"strconv"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

func CreateTokenNumber(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, error) {
	base, baseMatchedString, baseSkipper, baseValidator := ResolveBase(lexer)
	matchedString := ConstructChunk(lexer, string(lexer.Current), baseSkipper, baseValidator)
	number, err := strconv.ParseInt(matchedString, base, 64)
	if err != nil {
		return lexer_token.AwooLexerToken{}, matchedString, err
	}

	return lexer_token.CreateTokenPrimitive(lexer.Position, types.AwooTypeInt64, number, base), (baseMatchedString + matchedString), nil
}

func BaseSkipper(c rune) bool {
	return c == '_'
}

func BaseValidator(c rune) bool {
	return unicode.IsNumber(c)
}

func Base16Validator(c rune) bool {
	return unicode.IsNumber(c) || (c >= 'a' && c <= 'f')
}

func ResolveBase(lexer *AwooLexer) (int, string, ConstructChunkValidator, ConstructChunkValidator) {
	if lexer.Current == '0' {
		c, ok := PeekLexer(lexer)
		if !ok {
			return 10, "", ConstructChunkSkipperDefault, BaseValidator
		}
		switch unicode.ToLower(c) {
		case 'b':
			AdvanceLexerFor(lexer, 2)
			return 2, "0b", BaseSkipper, BaseValidator
		case 'o':
			AdvanceLexerFor(lexer, 2)
			return 8, "0o", BaseSkipper, BaseValidator
		case 'x':
			AdvanceLexerFor(lexer, 2)
			return 16, "0x", BaseSkipper, Base16Validator
		}
	}

	return 10, "", ConstructChunkSkipperDefault, BaseValidator
}
