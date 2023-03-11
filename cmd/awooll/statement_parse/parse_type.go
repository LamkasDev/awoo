package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructNodeType(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) node.AwooParserNodeResult {
	n := node.CreateNodeType(t)
	for dereferenceToken, _ := parser.ExpectTokenOptional(cparser, token.TokenOperatorDereference); dereferenceToken != nil; dereferenceToken, _ = parser.ExpectTokenOptional(cparser, token.TokenOperatorDereference) {
		n = node.CreateNodePointer(t, n.Node)
	}

	return n
}

func ConstructNodeTypeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeType)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructNodeType(cparser, t), nil
}
