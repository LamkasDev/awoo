package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func ConstructNodeType(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	n := node.CreateNodeType(t)
	for dereferenceToken, _ := parser.ExpectTokenOptional(cparser, token.TokenOperatorDereference); dereferenceToken != nil; dereferenceToken, _ = parser.ExpectTokenOptional(cparser, token.TokenOperatorDereference) {
		n = node.CreateNodePointer(t, n.Node)
	}
	if arrToken, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeBracketSquareLeft); arrToken != nil {
		n = node.CreateNodeTypeArray(t, n.Node)
		sizeToken, err := parser.ExpectToken(cparser, token.TokenTypePrimitive)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		node.SetNodeTypeArraySize(&n.Node, GetPrimitiveValue[uint32](cparser.Context.Lexer, sizeToken))
		if _, err := parser.ExpectToken(cparser, token.TokenTypeBracketSquareRight); err != nil {
			return node.AwooParserNodeResult{}, err
		}
	}

	return n, nil
}

func ConstructNodeTypeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeType)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructNodeType(cparser, t)
}
