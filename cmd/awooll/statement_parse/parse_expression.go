package statement_parse

import (
	"fmt"

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

func ConstructExpressionEndStatement(cparser *parser.AwooParser, n node.AwooParserNode, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if details.PendingBrackets > 0 {
		return node.AwooParserNodeResult{
			Error: fmt.Errorf("expected a %s", gchalk.Red(")")),
		}
	}
	return node.AwooParserNodeResult{
		Node: n,
		End:  true,
	}
}

func ConstructExpressionEndBracket(cparser *parser.AwooParser, n node.AwooParserNode, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if details.PendingBrackets > 0 {
		details.PendingBrackets--
		return node.AwooParserNodeResult{
			Node:       n,
			EndBracket: true,
		}
	}
	return node.AwooParserNodeResult{
		Error: fmt.Errorf("unexpected %s", gchalk.Red(")")),
	}
}

func ConstructExpressionAccumulate(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	op, err := parser.FetchTokenParser(cparser)
	if err != nil {
		return leftNode
	}
	switch op.Type {
	case token.TokenTypeEndStatement:
		if !details.EndWithCurly {
			return ConstructExpressionEndStatement(cparser, leftNode.Node, details)
		}
	case token.TokenTypeBracketCurlyLeft:
		if details.EndWithCurly {
			return ConstructExpressionEndStatement(cparser, leftNode.Node, details)
		}
	case token.TokenTypeBracketRight:
		return ConstructExpressionEndBracket(cparser, leftNode.Node, details)
	case token.TokenOperatorAddition,
		token.TokenOperatorSubstraction,
		token.TokenOperatorMultiplication,
		token.TokenOperatorDivision:
		rightNode := ConstructExpressionNegativeFast(cparser, details)
		if rightNode.Error != nil {
			return rightNode
		}
		if leftNode.Node.Type == node.ParserNodeTypeExpression &&
			!node.GetNodeExpressionIsBracket(&leftNode.Node) && token.DoesTokenTakePrecendence(op.Type, leftNode.Node.Token.Type) {
			n := node.CreateNodeExpression(op, node.GetNodeExpressionRight(&leftNode.Node), rightNode.Node)
			return node.AwooParserNodeResult{
				Node: node.CreateNodeExpression(leftNode.Node.Token, node.GetNodeExpressionLeft(&leftNode.Node), n),
			}
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
		}
	case token.TokenOperatorEq:
		op, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "==")
		if err != nil {
			return node.AwooParserNodeResult{
				Error: err,
			}
		}
		rightNode := ConstructExpressionNegativeFast(cparser, details)
		if rightNode.Error != nil {
			return rightNode
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
				Type:  token.TokenOperatorEqEq,
				Start: op.Start - 1,
			}, leftNode.Node, rightNode.Node),
		}
	case token.TokenTypeNot:
		op, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "!=")
		if err != nil {
			return node.AwooParserNodeResult{
				Error: err,
			}
		}
		rightNode := ConstructExpressionNegativeFast(cparser, details)
		if rightNode.Error != nil {
			return rightNode
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
				Type:  token.TokenOperatorNotEq,
				Start: op.Start - 1,
			}, leftNode.Node, rightNode.Node),
		}
	case token.TokenOperatorLT,
		token.TokenOperatorGT:
		t, err := parser.FetchTokenParser(cparser)
		if err != nil {
			return leftNode
		}
		if t.Type == token.TokenOperatorEq {
			if op.Type == token.TokenOperatorLT {
				op.Type = token.TokenOperatorLTEQ
			} else {
				op.Type = token.TokenOperatorGTEQ
			}
			t, err = parser.FetchTokenParser(cparser)
			if err != nil {
				return leftNode
			}
		}
		rightNode := ConstructExpressionNegative(cparser, t, details)
		if rightNode.Error != nil {
			return rightNode
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(op, leftNode.Node, rightNode.Node),
		}
	}

	opSymbol := "operator, <, >"
	if details.PendingBrackets > 0 {
		opSymbol += ", )"
	}
	endSymbol := ";"
	if details.EndWithCurly {
		endSymbol = "{"
	}
	return node.AwooParserNodeResult{
		Error: fmt.Errorf("expected an %s", gchalk.Red(fmt.Sprintf("%s or %s", opSymbol, endSymbol))),
	}
}

func ConstructExpressionBracket(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	leftNode := ConstructExpressionNegative(cparser, t, details)
	for leftNode.Error == nil && !leftNode.End && !leftNode.EndBracket {
		leftNode = ConstructExpressionAccumulate(cparser, leftNode, details)
	}
	if leftNode.Node.Type == node.ParserNodeTypeExpression {
		node.SetNodeExpressionIsBracket(&leftNode.Node, true)
	}
	leftNode.EndBracket = false

	return leftNode
}

func ConstructExpressionBracketFast(cparser *parser.AwooParser, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := parser.FetchTokenParser(cparser)
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}

	return ConstructExpressionBracket(cparser, t, details)
}

func ConstructExpressionStart(cparser *parser.AwooParser, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	leftNode := ConstructExpressionNegativeFast(cparser, details)
	for leftNode.Error == nil && !leftNode.End {
		leftNode = ConstructExpressionAccumulate(cparser, leftNode, details)
	}

	return leftNode
}
