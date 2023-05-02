package lexer

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
	"github.com/jwalton/gchalk"
)

func CreateTokenString(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, error) {
	tokenPosition := lexer_token.ExtendAwooLexerTokenPosition(lexer.Current.Position, lexer_token.AwooLexerTokenPosition{
		Length: 2,
	})
	matchedString := ConstructChunkFast(lexer, string(""), func(c rune) bool {
		return c != '"'
	})
	tokenPosition = lexer_token.ExtendAwooLexerTokenPosition(tokenPosition, lexer_token.AwooLexerTokenPosition{
		Length: uint32(len(matchedString)) - tokenPosition.Length,
	})
	if _, err := AdvanceLexer(lexer); err != nil {
		return lexer_token.AwooLexerToken{}, "", err
	}
	if lexer.Current.Character != '"' {
		return lexer_token.AwooLexerToken{}, "", fmt.Errorf("%w: %s", awerrors.ErrorIllegalCharacter, gchalk.Red((string)(lexer.Current.Character)))
	}

	return lexer_token.CreateTokenPrimitive(tokenPosition, types.AwooTypeString, matchedString, nil), matchedString, nil
}
