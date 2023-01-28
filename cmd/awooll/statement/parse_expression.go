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
	Type     types.AwooType
	Bracket  uint8
	Value    uint8
	Negative uint8
}

func ConstructExpressionEndStatement(context *parser_context.AwooParserContext, n node.AwooParserNode, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if details.Bracket > 0 {
		return node.AwooParserNodeResult{
			Error: fmt.Errorf("expected a %s", gchalk.Red(")")),
		}
	}
	return node.AwooParserNodeResult{
		Node: n,
		End:  true,
	}
}

func ConstructExpressionBracketRight(context *parser_context.AwooParserContext, n node.AwooParserNode, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if details.Bracket > 0 {
		details.Bracket--
		return node.AwooParserNodeResult{
			Node:       n,
			EndBracket: true,
		}
	}
	return node.AwooParserNodeResult{
		Error: fmt.Errorf("unexpected %s", gchalk.Red(")")),
	}
}

func ConstructExpressionContinue(context *parser_context.AwooParserContext, leftNode node.AwooParserNode, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if details.Value > 0 {
		details.Value--
		return node.AwooParserNodeResult{
			Node: leftNode,
		}
	}
	op, err := fetchToken()
	if err != nil {
		return node.AwooParserNodeResult{
			Node: leftNode,
		}
	}
	switch op.Type {
	case token.TokenTypeEndStatement:
		return ConstructExpressionEndStatement(context, leftNode, fetchToken, details)
	case token.TokenTypeBracketRight:
		return ConstructExpressionBracketRight(context, leftNode, fetchToken, details)
	case token.TokenOperatorAddition,
		token.TokenOperatorSubstraction,
		token.TokenOperatorMultiplication,
		token.TokenOperatorDivision:
		if op.Type == token.TokenOperatorMultiplication || op.Type == token.TokenOperatorDivision {
			details.Value++
		}
		rightNode := ConstructExpressionNegativeFast(context, fetchToken, details)
		if rightNode.Error != nil {
			return rightNode
		}
		resultNode := node.CreateNodeExpression(op, leftNode, rightNode.Node)
		if rightNode.End {
			return node.AwooParserNodeResult{
				Node: resultNode.Node,
				End:  true,
			}
		}
		return ConstructExpressionContinue(context, resultNode.Node, fetchToken, details)
	}

	// TODO: this text should say bracket if bracket > 0
	return node.AwooParserNodeResult{
		Error: fmt.Errorf("expected an %s", gchalk.Red("operator or ;")),
	}
}

func ConstructExpression(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	leftNode := ConstructNodeValue(context, t, fetchToken, details)
	if leftNode.Error != nil || leftNode.End {
		return leftNode
	}
	return ConstructExpressionContinue(context, leftNode.Node, fetchToken, details)
}

func ConstructExpressionFast(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier}, "primitive or identifier")
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}
	return ConstructExpression(context, t, fetchToken, details)
}
