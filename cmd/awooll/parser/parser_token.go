package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func TransformToken(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken) lexer_token.AwooLexerToken {
	switch t.Type {
	case token.TokenTypeIdentifier:
		possibleType, ok := context.Lexer.Types.Lookup[lexer_token.GetTokenIdentifierValue(&t)]
		if ok {
			t = lexer_token.CreateTokenType(t.Start, possibleType.Type)
		}
	}

	return t
}
