package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementAssignment(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, _ *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	var variableNameNode node.AwooParserNodeResult
	var variableName string
	var err error
	switch t.Type {
	case token.TokenOperatorDereference:
		idToken, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier, "identifier")
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		variableNameNode, err = CreateNodeIdentifierSafe(cparser, idToken)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		variableName = node.GetNodeIdentifierValue(&variableNameNode.Node)
		variableNameNode = node.CreateNodePointer(t, variableNameNode.Node)
	case token.TokenTypeIdentifier:
		variableNameNode, err = CreateNodeIdentifierSafe(cparser, t)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		variableName = node.GetNodeIdentifierValue(&variableNameNode.Node)
	}
	variableMemory, err := parser_context.GetParserScopeCurrentFunctionMemory(&cparser.Context, variableName)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	assignmentStatement := statement.CreateStatementAssignment(variableNameNode.Node)
	if _, err := parser.ExpectToken(cparser, token.TokenOperatorEq, "="); err != nil {
		return statement.AwooParserStatement{}, err
	}
	variableValueNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		Type:     cparser.Context.Lexer.Types.All[variableMemory.Type],
		EndToken: token.TokenTypeEndStatement,
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	statement.SetStatementAssignmentValue(&assignmentStatement, variableValueNode.Node)

	return assignmentStatement, nil
}
