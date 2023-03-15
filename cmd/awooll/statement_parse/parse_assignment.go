package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func GetVariableMemoryForAssignment(cparser *parser.AwooParser, identifierNode node.AwooParserNode) (parser_context.AwooParserMemoryEntry, error) {
	switch identifierNode.Type {
	case node.ParserNodeTypePointer:
		identifierNode = node.GetNodeSingleValue(&identifierNode)
		return parser_context.GetParserScopeCurrentFunctionMemory(&cparser.Context, node.GetNodeIdentifierValue(&identifierNode))
	case node.ParserNodeTypeArrayIndex:
		return parser_context.GetParserScopeCurrentFunctionMemory(&cparser.Context, node.GetNodeArrayIndexIdentifier(&identifierNode))
	}

	return parser_context.GetParserScopeCurrentFunctionMemory(&cparser.Context, node.GetNodeIdentifierValue(&identifierNode))
}

func ConstructStatementAssignment(cparser *parser.AwooParser, identifierNode node.AwooParserNode, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	variableMemory, err := GetVariableMemoryForAssignment(cparser, identifierNode)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	assignmentStatement := statement.CreateStatementAssignment(identifierNode)
	assignmentOperator, _ := parser.ExpectTokensOptional(cparser, []uint16{token.TokenOperatorAddition, token.TokenOperatorSubstraction, token.TokenOperatorMultiplication, token.TokenOperatorDivision})
	if _, err := parser.ExpectToken(cparser, token.TokenOperatorEq); err != nil {
		return statement.AwooParserStatement{}, err
	}
	variableValueNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		Type:     cparser.Context.Lexer.Types.All[variableMemory.Type],
		EndToken: details.EndToken,
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	if assignmentOperator != nil {
		variableValueNode = node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(*assignmentOperator, identifierNode, variableValueNode.Node),
		}
	}
	statement.SetStatementAssignmentValue(&assignmentStatement, variableValueNode.Node)

	return assignmentStatement, nil
}
