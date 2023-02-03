package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructExpressionNegative(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	if t.Type == token.TokenOperatorSubstraction {
		n := ConstructExpressionNegativeFast(cparser, details)
		if n.Error != nil {
			return node.AwooParserNodeResult{
				Error: n.Error,
			}
		}
		return node.CreateNodeNegative(t, n.Node)
	}
	return ConstructExpressionPriority(cparser, t, details)
}

func ConstructExpressionNegativeFast(cparser *parser.AwooParser, details *ConstructExpressionDetails) node.AwooParserNodeResult {
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypePrimitive, node.ParserNodeTypeIdentifier, token.TokenTypeBracketLeft, token.TokenOperatorSubstraction, token.TokenOperatorEq, token.TokenOperatorLT, token.TokenOperatorGT}, "primitive, identifier, (, -, < or >")
	if err != nil {
		return node.AwooParserNodeResult{
			Error: err,
		}
	}
	return ConstructExpressionNegative(cparser, t, details)
}
