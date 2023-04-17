package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awoocc/token"

type AwooLexerTokenDataIdentifier struct {
	Value string
}

func GetTokenIdentifierValue(t *AwooLexerToken) string {
	return t.Data.(AwooLexerTokenDataIdentifier).Value
}

func SetTokenIdentifierValue(t *AwooLexerToken, value string) {
	t.Data.(*AwooLexerTokenDataIdentifier).Value = value
}

func CreateTokenIdentifier(position AwooLexerTokenPosition, text string) AwooLexerToken {
	return AwooLexerToken{
		Type:     token.TokenTypeIdentifier,
		Position: position,
		Data: AwooLexerTokenDataIdentifier{
			Value: text,
		},
	}
}
