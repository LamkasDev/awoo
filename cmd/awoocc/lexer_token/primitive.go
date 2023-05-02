package lexer_token

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooLexerTokenDataPrimitive struct {
	Type  types.AwooTypeId
	Value interface{}
	Extra interface{}
}

func GetTokenPrimitiveType(t *AwooLexerToken) types.AwooTypeId {
	return t.Data.(AwooLexerTokenDataPrimitive).Type
}

func SetTokenPrimitiveType(t *AwooLexerToken, primitiveType types.AwooTypeId) {
	t.Data.(*AwooLexerTokenDataPrimitive).Type = primitiveType
}

func GetTokenPrimitiveValue(t *AwooLexerToken) interface{} {
	return t.Data.(AwooLexerTokenDataPrimitive).Value
}

func SetTokenPrimitiveValue(t *AwooLexerToken, value interface{}) {
	t.Data.(*AwooLexerTokenDataPrimitive).Value = value
}

func GetTokenPrimitiveExtra(t *AwooLexerToken) interface{} {
	return t.Data.(AwooLexerTokenDataPrimitive).Extra
}

func SetTokenPrimitiveExtra(t *AwooLexerToken, value interface{}) {
	t.Data.(*AwooLexerTokenDataPrimitive).Extra = value
}

func CreateTokenPrimitive(position AwooLexerTokenPosition, primitiveType types.AwooTypeId, value interface{}, extra interface{}) AwooLexerToken {
	return AwooLexerToken{
		Type:     token.TokenTypePrimitive,
		Position: position,
		Data: AwooLexerTokenDataPrimitive{
			Type:  primitiveType,
			Value: value,
			Extra: extra,
		},
	}
}
