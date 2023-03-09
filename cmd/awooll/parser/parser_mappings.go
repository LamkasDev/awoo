package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

type AwooParseStatement func(cparser *AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error)

type AwooParseNodeExpression func(cparser *AwooParser, leftNode node.AwooParserNodeResult, op lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error)

type AwooParseNodeValue func(cparser *AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error)

type AwooPrintStatement func(settings *AwooParserSettings, context *parser_context.AwooParserContext, s *statement.AwooParserStatement) string

type AwooParserMappings struct {
	Statement      map[uint16]AwooParseStatement
	NodeExpression map[uint16]AwooParseNodeExpression
	NodeValue      map[uint16]AwooParseNodeValue
	PrintStatement map[uint16]AwooPrintStatement
}
