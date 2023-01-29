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
	QueuedBrackets  uint8
	PendingBrackets uint8
	Value           uint8
	ValueBracket    uint8
	Negative        uint8
	NegativeBracket uint8
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
			Node:      n,
			IsBracket: true,
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
		var resultNode node.AwooParserNodeResult
		if op.Type == token.TokenOperatorMultiplication || op.Type == token.TokenOperatorDivision {
			resultNode = node.ProcessPrioritizedExpression(&leftNode, op, &rightNode)
		} else {
			resultNode.Node = node.CreateNodeExpression(op, leftNode.Node, rightNode.Node)
			resultNode.End = rightNode.End
			resultNode.IsBracket = rightNode.IsBracket
		}
		return resultNode
	}

	// TODO: this text should say bracket if bracket > 0
	return node.AwooParserNodeResult{
		Error: fmt.Errorf("expected an %s", gchalk.Red("operator or ;")),
	}
}

func ConstructExpressionContinue(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	leftNode := ConstructNodeValue(context, t, fetchToken, details)
	for leftNode.Error == nil && !leftNode.End && !leftNode.IsBracket {
		leftNode = ConstructExpressionAccumulate(context, leftNode, fetchToken, details)
	}

	return leftNode
}

func ConstructExpressionStart(context *parser_context.AwooParserContext, fetchToken lexer_token.FetchToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	leftNode := ConstructExpressionNegativeFast(context, fetchToken, details)
	for leftNode.Error == nil && !leftNode.End {
		leftNode = ConstructExpressionAccumulate(context, leftNode, fetchToken, details)
		leftNode.IsBracket = false
	}

	return leftNode
}
