package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func GetVariableMemoryForAssignment(cparser *parser.AwooParser, identifierNode node.AwooParserNode) (parser_context.AwooParserMemoryEntry, *parser_error.AwooParserError) {
	switch identifierNode.Type {
	case node.ParserNodeTypePointer:
		identifierNode = node.GetNodeSingleValue(&identifierNode)
		return parser_context.GetParserScopeFunctionMemory(&cparser.Context, node.GetNodeIdentifierValue(&identifierNode))
	case node.ParserNodeTypeArrayIndex:
		return parser_context.GetParserScopeFunctionMemory(&cparser.Context, node.GetNodeArrayIndexIdentifier(&identifierNode))
	}

	return parser_context.GetParserScopeFunctionMemory(&cparser.Context, node.GetNodeIdentifierValue(&identifierNode))
}

func ConstructStatementAssignment(cparser *parser.AwooParser, identifierNode node.AwooParserNode, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, *parser_error.AwooParserError) {
	variableMemory, err := GetVariableMemoryForAssignment(cparser, identifierNode)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	assignmentStatement := statement.CreateStatementAssignment(identifierNode)
	assignmentOperator, _ := parser.ExpectTokensOptional(cparser, []uint16{token.TokenOperatorAddition, token.TokenOperatorSubstraction, token.TokenOperatorMultiplication, token.TokenOperatorDivision})
	if _, err := parser.ExpectToken(cparser, token.TokenOperatorEq); err != nil {
		return statement.AwooParserStatement{}, err
	}
	valueNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		Type:      variableMemory.Type,
		EndTokens: []uint16{details.EndToken},
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}

	if assignmentOperator != nil {
		valueNode = node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(*assignmentOperator, identifierNode, valueNode.Node),
		}
	}
	statement.SetStatementAssignmentValue(&assignmentStatement, valueNode.Node)

	return assignmentStatement, nil
}
