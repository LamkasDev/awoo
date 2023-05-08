package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
)

type AwooParseStatement func(cparser *AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError)

type AwooParseNodeExpression func(cparser *AwooParser, leftNode node.AwooParserNodeResult, op lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError)

type AwooParseNodeValue func(cparser *AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError)

type AwooPrintStatement func(settings *AwooParserSettings, context *AwooParserContext, s *statement.AwooParserStatement) string

type AwooParserMappings struct {
	Statement      map[uint16]AwooParseStatement
	NodeExpression map[uint16]AwooParseNodeExpression
	NodeValue      map[uint16]AwooParseNodeValue
	PrintStatement map[uint16]AwooPrintStatement
}
