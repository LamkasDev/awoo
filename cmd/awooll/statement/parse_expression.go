package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

type ConstructExpressionDetails struct {
	Type            types.AwooType
	PendingBrackets uint8
}

func ConstructExpressionEndStatement(context *parser_context.AwooParserContext, n node.AwooParserNode, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
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

func ConstructExpressionEndBracket(context *parser_context.AwooParserContext, n node.AwooParserNode, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
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

func ConstructExpressionAccumulate(context *parser_context.AwooParserContext, leftNode node.AwooParserNodeResult, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	op, err := fetchToken()
	if err != nil {
		return leftNode
	}
	switch op.Type {
	case token.TokenTypeEndStatement:
		return ConstructExpressionEndStatement(context, leftNode.Node, fetchToken, details)
	case token.TokenTypeBracketRight:
		return ConstructExpressionEndBracket(context, leftNode.Node, fetchToken, details)
	case token.TokenOperatorAddition,
		token.TokenOperatorSubstraction,
		token.TokenOperatorMultiplication,
		token.TokenOperatorDivision:
		rightNode := ConstructExpressionNegativeFast(context, fetchToken, details)
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
		op, err = ExpectToken(fetchToken, []uint16{token.TokenOperatorEq}, "==")
		if err != nil {
			return node.AwooParserNodeResult{
				Error: err,
			}
		}
		rightNode := ConstructExpressionNegativeFast(context, fetchToken, details)
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
		op, err = ExpectToken(fetchToken, []uint16{token.TokenOperatorEq}, "!=")
		if err != nil {
			return node.AwooParserNodeResult{
				Error: err,
			}
		}
		rightNode := ConstructExpressionNegativeFast(context, fetchToken, details)
		if rightNode.Error != nil {
			return rightNode
		}
		return node.AwooParserNodeResult{
			Node: node.CreateNodeExpression(lexer_token.AwooLexerToken{
				Type:  token.TokenOperatorNotEq,
				Start: op.Start - 1,
			}, leftNode.Node, rightNode.Node),
		}
	}

	if details.PendingBrackets > 0 {
		return node.AwooParserNodeResult{
			Error: fmt.Errorf("expected an %s", gchalk.Red("operator, ) or ;")),
		}
	}
	return node.AwooParserNodeResult{
		Error: fmt.Errorf("expected an %s", gchalk.Red("operator or ;")),
	}
}

func ConstructExpressionBracket(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	leftNode := ConstructExpressionNegative(context, t, fetchToken, details)
	for leftNode.Error == nil && !leftNode.End && !leftNode.EndBracket {
		leftNode = ConstructExpressionAccumulate(context, leftNode, fetchToken, details)
	}
	if leftNode.Node.Type == node.ParserNodeTypeExpression {
		node.SetNodeExpressionIsBracket(&leftNode.Node, true)
	}
	leftNode.EndBracket = false

	return leftNode
}

func ConstructExpressionBracketFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := fetchToken()
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}

	return ConstructExpressionBracket(context, t, fetchToken, details)
}

func ConstructExpressionStart(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	leftNode := ConstructExpressionNegativeFast(context, fetchToken, details)
	for leftNode.Error == nil && !leftNode.End {
		leftNode = ConstructExpressionAccumulate(context, leftNode, fetchToken, details)
	}

	return leftNode
}
