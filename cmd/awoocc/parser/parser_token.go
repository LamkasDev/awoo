package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func TransformToken(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken) lexer_token.AwooLexerToken {
	switch t.Type {
	case token.TokenTypeIdentifier:
		possibleType, ok := lexer_context.GetContextType(&context.Lexer, lexer_token.GetTokenIdentifierValue(&t))
		if ok {
			t = lexer_token.CreateTokenType(t.Start, possibleType.Id)
		}
	}

	return t
}
