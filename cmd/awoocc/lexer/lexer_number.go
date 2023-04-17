package lexer

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
)

func CreateTokenNumber(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, error) {
	tokenPosition := lexer.Current.Position
	base, baseMatchedString, baseSkipper, baseValidator := ResolveBase(lexer)
	matchedString := ConstructChunk(lexer, string(lexer.Current.Character), baseSkipper, baseValidator)
	tokenPosition = lexer_token.ExtendAwooLexerTokenPosition(tokenPosition, lexer_token.AwooLexerTokenPosition{
		Length: uint32(len(baseMatchedString)+len(matchedString)) - tokenPosition.Length,
	})
	number, err := strconv.ParseInt(matchedString, base, 32)
	if err != nil {
		return lexer_token.AwooLexerToken{}, matchedString, fmt.Errorf("%w: %w", awerrors.ErrorFailedToCreateToken, err)
	}

	return lexer_token.CreateTokenPrimitive(tokenPosition, types.AwooTypeInt32, int32(number), base), (baseMatchedString + matchedString), nil
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
	if lexer.Current.Character == '0' {
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
