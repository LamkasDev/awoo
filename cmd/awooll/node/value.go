package node

import "github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"

// this is retarded, but i'm too stupid to figure this out
func ProcessPrioritizedExpression(left *AwooParserNodeResult, op lexer_token.AwooLexerToken, right *AwooParserNodeResult) AwooParserNodeResult {
	var result AwooParserNodeResult
	result.End = right.End
	if left.Node.Type != ParserNodeTypeExpression || left.IsBracket {
		if right.Node.Type != ParserNodeTypeExpression || right.IsBracket {
			// prim * prim
			result.Node = CreateNodeExpression(op, left.Node, right.Node)
		} else {
			// prim * expression
			tempRight := GetNodeExpressionLeft(&right.Node)
			result.Node = CreateNodeExpression(op, left.Node, tempRight)
			result.Node = CreateNodeExpression(right.Node.Token, result.Node, GetNodeExpressionRight(&right.Node))
		}
	} else {
		if right.Node.Type != ParserNodeTypeExpression || right.IsBracket {
			// expression * prim
			tempLeft := GetNodeExpressionRight(&left.Node)
			result.Node = CreateNodeExpression(op, tempLeft, right.Node)
			result.Node = CreateNodeExpression(left.Node.Token, GetNodeExpressionLeft(&left.Node), result.Node)
		} else {
			// expression * expression
			tempLeft := GetNodeExpressionRight(&left.Node)
			tempRight := GetNodeExpressionLeft(&right.Node)
			result.Node = CreateNodeExpression(op, tempLeft, tempRight)
			result.Node = CreateNodeExpression(left.Node.Token, GetNodeExpressionLeft(&left.Node), result.Node)
			result.Node = CreateNodeExpression(right.Node.Token, result.Node, GetNodeExpressionRight(&right.Node))
		}
	}

	return result
}
