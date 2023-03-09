package lexer

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

func CreateTokenChar(lexer *AwooLexer) (lexer_token.AwooLexerToken, string, error) {
	character, err := AdvanceLexer(lexer)
	if err != nil {
		return lexer_token.AwooLexerToken{}, string(character), err
	}
	if _, err := AdvanceLexer(lexer); err != nil {
		return lexer_token.AwooLexerToken{}, string(character), err
	}
	if lexer.Current != '\'' {
		return lexer_token.AwooLexerToken{}, string(character), fmt.Errorf("%w: %s", awerrors.ErrorIllegalCharacter, gchalk.Red((string)(lexer.Current)))
	}

	return lexer_token.CreateTokenPrimitive(lexer.Position, types.AwooTypeChar, character, nil), string(character), nil
}
