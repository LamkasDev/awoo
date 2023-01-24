package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awooll/token"

type AwooLexerTokenDataPrimitive struct {
	Value interface{}
}

func GetTokenPrimitiveValue(t *AwooLexerToken) interface{} {
	return t.Data.(AwooLexerTokenDataPrimitive).Value
}

func SetTokenPrimitiveValue(t *AwooLexerToken, value interface{}) {
	t.Data.(*AwooLexerTokenDataPrimitive).Value = value
}

func CreateTokenPrimitive(start uint16, value interface{}) AwooLexerToken {
	return AwooLexerToken{
		Type:  token.TokenTypePrimitive,
		Start: start,
		Data: AwooLexerTokenDataPrimitive{
			Value: value,
		},
	}
}
