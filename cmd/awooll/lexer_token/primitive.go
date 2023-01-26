package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awooll/token"

type AwooLexerTokenDataPrimitive struct {
	Type  uint16
	Value interface{}
}

func GetTokenPrimitiveType(t *AwooLexerToken) uint16 {
	return t.Data.(AwooLexerTokenDataPrimitive).Type
}

func SetTokenPrimitiveType(t *AwooLexerToken, primitiveType uint16) {
	t.Data.(*AwooLexerTokenDataPrimitive).Type = primitiveType
}

func GetTokenPrimitiveValue(t *AwooLexerToken) interface{} {
	return t.Data.(AwooLexerTokenDataPrimitive).Value
}

func SetTokenPrimitiveValue(t *AwooLexerToken, value interface{}) {
	t.Data.(*AwooLexerTokenDataPrimitive).Value = value
}

func CreateTokenPrimitive(start uint16, primitiveType uint16, value interface{}) AwooLexerToken {
	return AwooLexerToken{
		Type:  token.TokenTypePrimitive,
		Start: start,
		Data: AwooLexerTokenDataPrimitive{
			Type:  primitiveType,
			Value: value,
		},
	}
}
