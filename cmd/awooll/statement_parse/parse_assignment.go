package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementAssignment(cparser *parser.AwooParser, variableNameNode node.AwooParserNode, variableName string) (statement.AwooParserStatement, error) {
	variableMemory, err := parser_context.GetParserScopeCurrentFunctionMemory(&cparser.Context, variableName)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	assignmentStatement := statement.CreateStatementAssignment(variableNameNode)
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
