package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awooll/token"

type AwooLexerToken struct {
	Type  uint16
	Start uint16
	Data  interface{}
}

func CreateToken(start uint16, t *token.AwooToken) AwooLexerToken {
	return AwooLexerToken{
		Type:  t.Type,
		Start: start - uint16(t.Length) + 1,
	}
}
