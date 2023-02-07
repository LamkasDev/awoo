package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

type ConstructExpressionDetails struct {
	Type            types.AwooType
	PendingBrackets uint8
	EndWithCurly    bool
}

func ConstructExpressionAccumulate(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	op, err := parser.FetchTokenParser(cparser)
	if err != nil {
		return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	// TODO: refactor using a map
	switch op.Type {
	case token.TokenTypeEndStatement:
		if !details.EndWithCurly {
			return ConstructExpressionEndStatement(leftNode.Node, details)
		}
	case token.TokenTypeBracketCurlyLeft:
		if details.EndWithCurly {
			return ConstructExpressionEndStatement(leftNode.Node, details)
		}
	case token.TokenTypeBracketRight:
		return ConstructExpressionEndBracket(leftNode.Node, details)
	case token.TokenOperatorAddition,
		token.TokenOperatorSubstraction,
		token.TokenOperatorMultiplication,
		token.TokenOperatorDivision:
		return ConstructExpressionUnary(cparser, leftNode, op, details)
	case token.TokenOperatorEq:
		return ConstructExpressionEquality(cparser, leftNode, details)
	case token.TokenTypeNot:
		return ConstructExpressionNotEquality(cparser, leftNode, details)
	case token.TokenOperatorLT,
		token.TokenOperatorGT:
		return ConstructExpressionComparison(cparser, leftNode, op, details)
	}

	opSymbol := "operator, <, >"
	if details.PendingBrackets > 0 {
		opSymbol += ", )"
	}
	endSymbol := ";"
	if details.EndWithCurly {
		endSymbol = "{"
	}
	return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(fmt.Sprintf("%s or %s", opSymbol, endSymbol)))
}

func ConstructExpressionBracket(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	leftNode, err := ConstructExpressionReference(cparser, t, details)
	for err == nil && !leftNode.End && !leftNode.EndBracket {
		leftNode, err = ConstructExpressionAccumulate(cparser, leftNode, details)
	}
	if err != nil {
		return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
	if leftNode.Node.Type == node.ParserNodeTypeExpression {
		node.SetNodeExpressionIsBracket(&leftNode.Node, true)
	}
	leftNode.EndBracket = false

	return leftNode, err
}

func ConstructExpressionBracketFast(cparser *parser.AwooParser, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.FetchTokenParser(cparser)
	if err != nil {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}

	return ConstructExpressionBracket(cparser, t, details)
}

func ConstructExpressionStart(cparser *parser.AwooParser, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	leftNode, err := ConstructExpressionReferenceFast(cparser, details)
	for err == nil && !leftNode.End {
		leftNode, err = ConstructExpressionAccumulate(cparser, leftNode, details)
	}
	if err != nil {
		return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}

	return leftNode, nil
}
