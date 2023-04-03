package lexer_token

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooLexerTokenDataType struct {
	Id types.AwooTypeId
}

func GetTokenTypeId(t *AwooLexerToken) types.AwooTypeId {
	return t.Data.(AwooLexerTokenDataType).Id
}

func SetTokenTypeId(t *AwooLexerToken, id types.AwooTypeId) {
	t.Data.(*AwooLexerTokenDataType).Id = id
}

func CreateTokenType(start uint32, value types.AwooTypeId) AwooLexerToken {
	return AwooLexerToken{
		Type:  token.TokenTypeType,
		Start: start,
		Data: AwooLexerTokenDataType{
			Id: value,
		},
	}
}
