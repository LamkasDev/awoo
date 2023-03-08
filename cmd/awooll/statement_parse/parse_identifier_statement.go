package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementIdentifier(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, _ *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	switch t.Type {
	case token.TokenOperatorDereference:
		t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier, "identifier")
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		identifierNode, err := CreateNodeIdentifierSafe(cparser, t)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		identifier := node.GetNodeIdentifierValue(&identifierNode.Node)
		identifierNode = node.CreateNodePointer(t, identifierNode.Node)

		return ConstructStatementAssignment(cparser, identifierNode.Node, identifier)
	case token.TokenTypeIdentifier:
		identifierNode, err := CreateNodeIdentifierSafe(cparser, t)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		if identifierNode.Node.Type == node.ParserNodeTypeCall {
			callStatement := statement.CreateStatementCall(identifierNode.Node)
			if _, err := parser.ExpectToken(cparser, token.TokenTypeEndStatement, ";"); err != nil {
				return callStatement, err
			}
			return callStatement, nil
		}

		return ConstructStatementAssignment(cparser, identifierNode.Node, node.GetNodeIdentifierValue(&identifierNode.Node))
	}

	return statement.AwooParserStatement{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red("identifier"))
}
