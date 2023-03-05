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
	t, err := parser.ExpectTokensParser(cparser, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierVariableSafe(cparser, t)
}

func CreateNodeIdentifierCallSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	_, err := parser.ExpectTokensParser(cparser, []uint16{token.TokenTypeBracketLeft}, "(")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}
	callFunction, ok := parser_context.GetContextFunction(&cparser.Context, identifier)
	if !ok {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownFunction, gchalk.Red(identifier))
	}

	callNode := node.CreateNodeCall(t)
	for i, arg := range callFunction.Arguments {
		details := parser_details.ConstructExpressionDetails{
			Type:     cparser.Contents.Context.Types.All[arg.Type],
			EndToken: token.TokenTypeBracketRight,
		}
		if i < len(callFunction.Arguments)-1 {
			details.EndToken = token.TokenTypeComma
		}
		argNode, err := ConstructExpressionStart(cparser, &details)
		if err != nil {
			return callNode, err
		}
		node.SetNodeCallArguments(&callNode.Node, append(node.GetNodeCallArguments(&callNode.Node), argNode.Node))
	}
	if len(callFunction.Arguments) == 0 {
		if _, err := parser.ExpectTokenParser(cparser, token.TokenTypeBracketRight, ")"); err != nil {
			return callNode, err
		}
	}

	return callNode, nil
}

func CreateNodeIdentifierCallSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokensParser(cparser, []uint16{node.ParserNodeTypeIdentifier}, "identifier")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierCallSafe(cparser, t)
}

func CreateNodeIdentifierSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	variableNode, variableErr := CreateNodeIdentifierVariableSafe(cparser, t)
	if variableErr == nil {
		return variableNode, nil
	}
	callNode, callErr := CreateNodeIdentifierCallSafe(cparser, t)
	if callErr == nil {
		return callNode, nil
	}

	return node.AwooParserNodeResult{}, variableErr
}

func CreateNodeIdentifierSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokenParser(cparser, token.TokenTypeIdentifier, "identifier")
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierSafe(cparser, t)
}
