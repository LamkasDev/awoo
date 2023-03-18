package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

func ConstructExpressionAccumulate(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	op, err := parser.FetchToken(cparser)
	if err != nil {
		return leftNode, err
	}
	if util.Contains(details.EndTokens, op.Type) {
		return ConstructExpressionEndStatement(cparser, leftNode, op, details)
	}
	entry, ok := cparser.Settings.Mappings.NodeExpression[op.Type]
	if !ok {
		expectedTypes := append([]uint16{token.TokenOperatorLT, token.TokenOperatorGT}, details.EndTokens...)
		if details.PendingBrackets > 0 {
			expectedTypes = append(expectedTypes, token.TokenTypeBracketRight)
		}
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(fmt.Sprintf("operator, %s", lexer.PrintTokenTypes(&cparser.Settings.Lexer, expectedTypes))))
	}

	return entry(cparser, leftNode, op, details)
}

func ConstructExpressionBracket(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	leftNode, err := ConstructExpressionReference(cparser, t, details)
	for err == nil && leftNode.End == nil {
		leftNode, err = ConstructExpressionAccumulate(cparser, leftNode, details)
	}
	if err != nil {
		return leftNode, err
	}
	if leftNode.Node.Type == node.ParserNodeTypeExpression {
		node.SetNodeExpressionIsBracket(&leftNode.Node, true)
	}
	leftNode.End = nil

	return leftNode, err
}

func ConstructExpressionBracketFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.FetchToken(cparser)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return ConstructExpressionBracket(cparser, t, details)
}

func ConstructExpressionStart(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	leftNode, err := ConstructExpressionReferenceFast(cparser, details)
	for err == nil && leftNode.End == nil {
		leftNode, err = ConstructExpressionAccumulate(cparser, leftNode, details)
	}
	if err != nil {
		return leftNode, err
	}

	return leftNode, nil
}
