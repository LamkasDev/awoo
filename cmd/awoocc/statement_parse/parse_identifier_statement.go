package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementIdentifier(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, *parser_error.AwooParserError) {
	switch t.Type {
	case token.TokenOperatorDereference:
		t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		identifierNode, err := CreateNodeIdentifierSafe(cparser, t, &parser_details.ConstructExpressionDetails{})
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		identifierNode = node.CreateNodePointer(t, identifierNode.Node)

		return ConstructStatementAssignment(cparser, identifierNode.Node, details)
	case token.TokenTypeIdentifier:
		identifierNode, err := CreateNodeIdentifierSafe(cparser, t, &parser_details.ConstructExpressionDetails{})
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		if identifierNode.Node.Type == node.ParserNodeTypeCall {
			callStatement := statement.CreateStatementCall(identifierNode.Node)
			if _, err := parser.ExpectToken(cparser, details.EndToken); err != nil {
				return callStatement, err
			}
			return callStatement, nil
		}

		return ConstructStatementAssignment(cparser, identifierNode.Node, details)
	}

	return statement.AwooParserStatement{}, parser_error.CreateParserErrorText(parser_error.AwooParserErrorExpectedToken,
		fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorExpectedToken], gchalk.Red("identifier")),
		cparser.Current.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorExpectedToken])
}
