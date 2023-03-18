package parser_context

import "github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"

type AwooParserContext struct {
	Lexer     lexer_context.AwooLexerContext
	Scopes    AwooParserScopeContainer
	Functions AwooParserFunctionContainer
}
