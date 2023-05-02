package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/jwalton/gchalk"
)

func GetVariableMemoryForAssignment(cparser *parser.AwooParser, identifierNode node.AwooParserNode) (elf.AwooElfSymbolTableEntry, *parser_error.AwooParserError) {
	switch identifierNode.Type {
	case node.ParserNodeTypePointer:
		identifierNode = node.GetNodeSingleValue(&identifierNode)
	case node.ParserNodeTypeArrayIndex:
		identifier := node.GetNodeArrayIndexIdentifier(&identifierNode)
		entry, ok := parser_context.GetParserScopeFunctionSymbol(&cparser.Context, identifier)
		if !ok {
			return entry, parser_error.CreateParserErrorText(parser_error.AwooParserErrorUnknownVariable,
				fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorUnknownVariable], gchalk.Red(identifier)),
				identifierNode.Token.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorUnknownVariable])
		}
		return entry, nil
	}

	identifier := node.GetNodeIdentifierValue(&identifierNode)
	entry, ok := parser_context.GetParserScopeFunctionSymbol(&cparser.Context, identifier)
	if !ok {
		return entry, parser_error.CreateParserErrorText(parser_error.AwooParserErrorUnknownVariable,
			fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorUnknownVariable], gchalk.Red(identifier)),
			identifierNode.Token.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorUnknownVariable])
	}
	return entry, nil
}

func ConstructStatementAssignment(cparser *parser.AwooParser, identifierNode node.AwooParserNode, details *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	variableMemory, err := GetVariableMemoryForAssignment(cparser, identifierNode)
	if err != nil {
		return nil, err
	}
	assignmentStatement := statement.CreateStatementAssignment(identifierNode)
	assignmentOperator, _ := parser.ExpectTokensOptional(cparser, []uint16{token.TokenOperatorAddition, token.TokenOperatorSubstraction, token.TokenOperatorMultiplication, token.TokenOperatorDivision})
	if _, err := parser.ExpectToken(cparser, token.TokenOperatorEq); err != nil {
		return nil, err
	}
	valueNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		Type:      variableMemory.Type,
		EndTokens: []uint16{details.EndToken},
	})
	if err != nil {
		return nil, err
	}

	if assignmentOperator != nil {
		valueNode = node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(*assignmentOperator, identifierNode, valueNode.Node),
		}
	}
	statement.SetStatementAssignmentValue(&assignmentStatement, valueNode.Node)

	return &assignmentStatement, nil
}
