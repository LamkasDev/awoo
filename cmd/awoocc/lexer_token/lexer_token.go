package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awoocc/token"

type AwooLexerToken struct {
	Type  uint16
	Start uint32
	Data  interface{}
}

type FetchToken func() (AwooLexerToken, error)

func CreateToken(start uint32, t *token.AwooToken) AwooLexerToken {
	return AwooLexerToken{
		Type:  t.Id,
		Start: start - uint32(t.Length) + 1,
	}
}
