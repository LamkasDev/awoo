package lexer_token

import "github.com/LamkasDev/awoo-emu/cmd/awoocc/token"

type AwooLexerToken struct {
	Type     uint16
	Position AwooLexerTokenPosition
	Data     interface{}
}

type FetchToken func() (AwooLexerToken, error)

func NewAwooLexerToken(position AwooLexerTokenPosition, t *token.AwooToken) AwooLexerToken {
	return AwooLexerToken{
		Type:     t.Id,
		Position: position,
	}
}

type AwooLexerTokenPosition struct {
	Index  uint32
	Line   uint32
	Column uint32
	Length uint32
}

func ExtendAwooLexerTokenPosition(previous AwooLexerTokenPosition, current AwooLexerTokenPosition) AwooLexerTokenPosition {
	return AwooLexerTokenPosition{
		Index:  current.Index,
		Line:   previous.Line,
		Column: previous.Column,
		Length: previous.Length + current.Length,
	}
}
