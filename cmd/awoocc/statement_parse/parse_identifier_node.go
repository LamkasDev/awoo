package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
	"github.com/jwalton/gchalk"
)

func CreateNodeIdentifierVariableSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	if _, err := parser_context.GetParserScopeCurrentFunctionMemory(&cparser.Context, identifier); err != nil {
		return node.AwooParserNodeResult{}, err
	}
	if arrToken, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeBracketSquareLeft); arrToken != nil {
		arrIndexNode := node.CreateNodeArrayIndex(*arrToken, identifier)
		indexNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
			Type:      commonTypes.AwooTypeId(types.AwooTypeUInt16),
			EndTokens: []uint16{token.TokenTypeBracketSquareRight},
		})
		if err != nil {
			return arrIndexNode, err
		}
		node.SetNodeArrayIndexIndex(&arrIndexNode.Node, indexNode.Node)
		return arrIndexNode, nil
	}

	return node.CreateNodeIdentifier(t), nil
}

func CreateNodeIdentifierVariableSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypeIdentifier})
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierVariableSafe(cparser, t)
}

func CreateNodeIdentifierCallSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	callFunctionName := lexer_token.GetTokenIdentifierValue(&t)
	if _, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypeBracketLeft}); err != nil {
		return node.AwooParserNodeResult{}, err
	}
	callFunction, ok := parser_context.GetParserFunction(&cparser.Context, callFunctionName)
	if !ok {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownFunction, gchalk.Red(callFunctionName))
	}

	callNode := node.CreateNodeCall(t)
	for _, arg := range callFunction.Arguments {
		details := parser_details.ConstructExpressionDetails{
			Type:      arg.Type,
			EndTokens: []uint16{token.TokenTypeBracketRight, token.TokenTypeComma},
		}
		argNode, err := ConstructExpressionStart(cparser, &details)
		if err != nil {
			return callNode, err
		}
		node.SetNodeCallArguments(&callNode.Node, append(node.GetNodeCallArguments(&callNode.Node), argNode.Node))
	}
	if len(callFunction.Arguments) == 0 {
		if _, err := parser.ExpectToken(cparser, token.TokenTypeBracketRight); err != nil {
			return callNode, err
		}
	}

	return callNode, nil
}

func CreateNodeIdentifierCallSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectTokens(cparser, []uint16{node.ParserNodeTypeIdentifier})
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierCallSafe(cparser, t)
}

func CreateNodeIdentifierSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, _ *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
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

func CreateNodeIdentifierSafeFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierSafe(cparser, t, details)
}