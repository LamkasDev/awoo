package lexer

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
	"github.com/jwalton/gchalk"
)

func CreateTokenCharacter(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, error) {
	tokenPosition := lexer_token.ExtendAwooLexerTokenPosition(lexer.Current.Position, lexer_token.AwooLexerTokenPosition{
		Length: 2,
	})
	character, err := AdvanceLexer(lexer)
	if err != nil {
		return lexer_token.AwooLexerToken{}, "", err
	}
	if _, err := AdvanceLexer(lexer); err != nil {
		return lexer_token.AwooLexerToken{}, "", err
	}
	if lexer.Current.Character != '\'' {
		return lexer_token.AwooLexerToken{}, "", fmt.Errorf("%w: %s", awerrors.ErrorIllegalCharacter, gchalk.Red((string)(lexer.Current.Character)))
	}

	return lexer_token.CreateTokenPrimitive(tokenPosition, types.AwooTypeChar, character, nil), string(character), nil
}
