package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awooll/token"

type AwooLexerTokenDataIdentifier struct {
	Value string
}

func GetTokenIdentifierValue(t *AwooLexerToken) string {
	return t.Data.(AwooLexerTokenDataIdentifier).Value
}

func SetTokenIdentifierValue(t *AwooLexerToken, value string) {
	t.Data.(*AwooLexerTokenDataIdentifier).Value = value
}

func CreateTokenIdentifier(start uint16, text string) AwooLexerToken {
	return AwooLexerToken{
		Type:  token.TokenTypeIdentifier,
		Start: start,
		Data: AwooLexerTokenDataIdentifier{
			Value: text,
		},
	}
}
