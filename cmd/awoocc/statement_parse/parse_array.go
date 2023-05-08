package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
)

func CreateNodeArray(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	elements := []node.AwooParserNode{}
	if t := parser.ExpectTokenOptional(cparser, token.TokenTypeBracketCurlyLeft); t == nil {
		elementDetails := &parser_details.ConstructExpressionDetails{
			Type:      details.Type,
			EndTokens: []uint16{token.TokenTypeBracketCurlyRight, token.TokenTypeComma},
		}
		for true {
			elementNode, err := ConstructExpressionStart(cparser, elementDetails)
			if err != nil {
				return node.AwooParserNodeResult{}, err
			}
			elements = append(elements, elementNode.Node)
			if *elementNode.End == token.TokenTypeBracketCurlyRight {
				break
			}
		}
	}

	return node.CreateNodeArray(t, elements), nil
}
