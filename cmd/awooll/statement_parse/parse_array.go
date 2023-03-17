package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func CreateNodeArray(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	elements := []node.AwooParserNode{}
	if t, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeBracketCurlyLeft); t == nil {
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
