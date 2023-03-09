package lexer

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

func CreateTokenNumber(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, error) {
	base, baseMatchedString, baseSkipper, baseValidator := ResolveBase(lexer)
	matchedString := ConstructChunk(lexer, string(lexer.Current), baseSkipper, baseValidator)
	number, err := strconv.ParseInt(matchedString, base, 32)
	if err != nil {
		return lexer_token.AwooLexerToken{}, matchedString, fmt.Errorf("%w: %w", awerrors.ErrorFailedToParse, err)
	}

	return lexer_token.CreateTokenPrimitive(lexer.Position, types.AwooTypeInt32, int32(number), base), (baseMatchedString + matchedString), nil
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
		c, err := PeekLexer(lexer)
		if err != nil {
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
