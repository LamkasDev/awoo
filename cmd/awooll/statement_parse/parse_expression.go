package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/jwalton/gchalk"
)

func ConstructExpressionAccumulate(cparser *parser.AwooParser, leftNode node.AwooParserNodeResult, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	op, err := parser.FetchToken(cparser)
	if err != nil {
		return leftNode, err
	}
	if op.Type == details.EndToken {
		return ConstructExpressionEndStatement(cparser, leftNode, op, details)
	}
	entry, ok := cparser.Settings.Mappings.NodeExpression[op.Type]
	if !ok {
		opSymbol := "operator, <, >"
		if details.PendingBrackets > 0 {
			opSymbol += ", )"
		}
		endSymbol := cparser.Context.Lexer.Tokens.All[details.EndToken].Name
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(fmt.Sprintf("%s or %s", opSymbol, endSymbol)))
	}

	return entry(cparser, leftNode, op, details)
}

func ConstructExpressionBracket(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	leftNode, err := ConstructExpressionReference(cparser, t, details)
	for err == nil && !leftNode.End && !leftNode.EndBracket {
		leftNode, err = ConstructExpressionAccumulate(cparser, leftNode, details)
	}
	if err != nil {
		return leftNode, err
	}
	if leftNode.Node.Type == node.ParserNodeTypeExpression {
		node.SetNodeExpressionIsBracket(&leftNode.Node, true)
	}
	leftNode.EndBracket = false

	return leftNode, err
}

func ConstructExpressionBracketFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.FetchToken(cparser)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return ConstructExpressionBracket(cparser, t, details)
}

func ConstructExpressionStart(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	leftNode, err := ConstructExpressionReferenceFast(cparser, details)
	for err == nil && !leftNode.End {
		leftNode, err = ConstructExpressionAccumulate(cparser, leftNode, details)
	}
	if err != nil {
		return leftNode, err
	}

	return leftNode, nil
}
