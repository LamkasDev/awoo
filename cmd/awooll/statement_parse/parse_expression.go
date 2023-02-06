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

func ConstructExpressionEndStatement(n node.AwooParserNode, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if details.PendingBrackets > 0 {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(")"))
	}
	return node.AwooParserNodeResult{
		Node: n,
		End:  true,
	}, nil
}

func ConstructExpressionEndBracket(n node.AwooParserNode, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if details.PendingBrackets > 0 {
		details.PendingBrackets--
		return node.AwooParserNodeResult{
			Node:       n,
			EndBracket: true,
		}, nil
	}
	return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnxpectedToken, gchalk.Red(")"))
}

func ConstructExpressionAccumulate(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	op, err := parser.FetchTokenParser(cparser)
	if err != nil {
		return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
	}
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
		rightNode, err := ConstructExpressionReferenceFast(cparser, details)
		if err != nil {
			return rightNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
		if leftNode.Node.Type == node.ParserNodeTypeExpression &&
			!node.GetNodeExpressionIsBracket(&leftNode.Node) && token.DoesTokenTakePrecendence(op.Type, leftNode.Node.Token.Type) {
			n := node.CreateNodeExpression(op, node.GetNodeExpressionRight(&leftNode.Node), rightNode.Node)
			return node.AwooParserNodeResult{
				Node: node.CreateNodeExpression(leftNode.Node.Token, node.GetNodeExpressionLeft(&leftNode.Node), n),
			}, nil
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
		}, nil
	case token.TokenOperatorEq:
		op, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "==")
		if err != nil {
			return node.AwooParserNodeResult{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
		rightNode, err := ConstructExpressionReferenceFast(cparser, details)
		if err != nil {
			return rightNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
				Type:  token.TokenOperatorEqEq,
				Start: op.Start - 1,
			}, leftNode.Node, rightNode.Node),
		}, nil
	case token.TokenTypeNot:
		op, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "!=")
		if err != nil {
			return node.AwooParserNodeResult{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
		rightNode, err := ConstructExpressionReferenceFast(cparser, details)
		if err != nil {
			return rightNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
				Type:  token.TokenOperatorNotEq,
				Start: op.Start - 1,
			}, leftNode.Node, rightNode.Node),
		}, nil
	case token.TokenOperatorLT,
		token.TokenOperatorGT:
		t, err := parser.FetchTokenParser(cparser)
		if err != nil {
			return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
		if t.Type == token.TokenOperatorEq {
			if op.Type == token.TokenOperatorLT {
				op.Type = token.TokenOperatorLTEQ
			} else {
				op.Type = token.TokenOperatorGTEQ
			}
			t, err = parser.FetchTokenParser(cparser)
			if err != nil {
				return leftNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
			}
		}
		rightNode, err := ConstructExpressionReference(cparser, t, details)
		if err != nil {
			return rightNode, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructExpression, err)
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
		}, nil
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
