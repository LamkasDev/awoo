package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awoocc/token"

type AwooLexerTokenDataType struct {
	Id uint16
}

func GetTokenTypeId(t *AwooLexerToken) uint16 {
	return t.Data.(AwooLexerTokenDataType).Id
}

func SetTokenTypeId(t *AwooLexerToken, id uint16) {
	t.Data.(*AwooLexerTokenDataType).Id = id
}

func CreateTokenType(start uint16, value uint16) AwooLexerToken {
	return AwooLexerToken{
		Type:  token.TokenTypeType,
		Start: start,
		Data: AwooLexerTokenDataType{
			Id: value,
		},
	}
}
