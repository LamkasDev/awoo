package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func CreateNodeIdentifierVariableSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	_, ok := parser_context.GetContextVariable(&cparser.Context, identifier)
	if ok {
		return node.CreateNodeIdentifier(t), nil
	}

	return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(identifier))
}

func CreateNodeIdentifierVariableSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokensParser(cparser, []uint16{node.ParserNodeTypeIdentifier}, "identifier")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return CreateNodeIdentifierVariableSafe(cparser, t)
}

// TODO: this should call above.
func CreateNodeIdentifierSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	_, ok := parser_context.GetContextVariable(&cparser.Context, identifier)
	if ok {
		return node.CreateNodeIdentifier(t), nil
	}
	if tlb, ok := parser.PeekParser(cparser); ok && tlb.Type == token.TokenTypeBracketLeft {
		f, ok := parser_context.GetContextFunction(&cparser.Context, identifier)
		if !ok {
			return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownFunction, gchalk.Red(identifier))
		}
		callNode := node.CreateNodeCall(t)
		parser.AdvanceParser(cparser)
		for i, arg := range f.Arguments {
			details := parser_details.ConstructExpressionDetails{
				Type:     cparser.Contents.Context.Types.All[arg.Type],
				EndToken: token.TokenTypeBracketRight,
			}
			if i < len(f.Arguments)-1 {
				details.EndToken = token.TokenTypeComma
			}
			argNode, err := ConstructExpressionStart(cparser, &details)
			if err != nil {
				return node.AwooParserNodeResult{}, err
			}
			node.SetNodeCallArguments(&callNode.Node, append(node.GetNodeCallArguments(&callNode.Node), argNode.Node))
		}
		if len(f.Arguments) == 0 {
			if _, err := parser.ExpectTokenParser(cparser, token.TokenTypeBracketRight, ")"); err != nil {
				return node.AwooParserNodeResult{}, err
			}
		}

		return callNode, nil
	}

	return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(identifier))
}

func CreateNodeIdentifierSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokenParser(cparser, node.ParserNodeTypeIdentifier, "identifier")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	return CreateNodeIdentifierSafe(cparser, t)
}
