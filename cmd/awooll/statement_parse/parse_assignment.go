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
	variableNameNode, err := CreateNodeIdentifierSafe(cparser, t)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	contextVariable, ok := parser_context.GetContextVariable(&cparser.Context, node.GetNodeIdentifierValue(&variableNameNode.Node))
	if !ok {
		return statement.AwooParserStatement{}, err
	}
	assignmentStatement := statement.CreateStatementAssignment(variableNameNode.Node)
	if _, err = parser.ExpectTokenParser(cparser, token.TokenOperatorEq, "="); err != nil {
		return statement.AwooParserStatement{}, err
	}
	variableValueNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		Type:     cparser.Context.Lexer.Types.All[contextVariable.Type],
		EndToken: token.TokenTypeEndStatement,
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	statement.SetStatementAssignmentValue(&assignmentStatement, variableValueNode.Node)

	return assignmentStatement, nil
}
