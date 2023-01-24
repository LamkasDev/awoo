package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awooll/token"

type AwooLexerTokenDataType struct {
	Value uint16
}

func GetTokenTypeType(t *AwooLexerToken) uint16 {
	return t.Data.(AwooLexerTokenDataType).Value
}

func SetTokenTypeType(t *AwooLexerToken, value uint16) {
	t.Data.(*AwooLexerTokenDataType).Value = value
}

func CreateTokenType(start uint16, value uint16) AwooLexerToken {
	return AwooLexerToken{
		Type:  token.TokenTypeType,
		Start: start,
		Data: AwooLexerTokenDataType{
			Value: value,
		},
	}
}
