package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

func ConstructNodeType(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	n := node.CreateNodeType(t)
	for dereferenceToken := parser.ExpectTokenOptional(cparser, token.TokenOperatorDereference); dereferenceToken != nil; dereferenceToken = parser.ExpectTokenOptional(cparser, token.TokenOperatorDereference) {
		n = node.CreateNodePointer(t, n.Node)
	}
	if arrToken := parser.ExpectTokenOptional(cparser, token.TokenTypeBracketSquareLeft); arrToken != nil {
		n = node.CreateNodeTypeArray(t, n.Node)
		sizeToken, err := parser.ExpectToken(cparser, token.TokenTypePrimitive)
		if err != nil {
			return node.AwooParserNodeResult{}, err
		}
		node.SetNodeTypeArraySize(&n.Node, GetPrimitiveValue[arch.AwooRegister](cparser.Context.Lexer, *sizeToken))
		if _, err := parser.ExpectToken(cparser, token.TokenTypeBracketSquareRight); err != nil {
			return node.AwooParserNodeResult{}, err
		}
	}

	return n, nil
}

func ConstructNodeTypeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeType)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return ConstructNodeType(cparser, *t)
}
