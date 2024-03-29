package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
	"github.com/jwalton/gchalk"
)

func CreateNodeIdentifierVariableSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	symbol, err := scope.GetCurrentFunctionSymbol(&cparser.Context.Scopes, identifier)
	if err != nil || symbol.Symbol.Type == types.AwooTypeFunction {
		return node.AwooParserNodeResult{}, parser_error.CreateParserErrorText(parser_error.AwooParserErrorUnknownVariable,
			fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorUnknownVariable], gchalk.Red(identifier)),
			t.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorUnknownVariable])
	}
	if arrToken := parser.ExpectTokenOptional(cparser, token.TokenTypeBracketSquareLeft); arrToken != nil {
		arrIndexNode := node.CreateNodeArrayIndex(*arrToken, identifier)
		indexNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
			Type:      types.AwooTypeUInt16,
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

func CreateNodeIdentifierVariableSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypeIdentifier})
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierVariableSafe(cparser, *t)
}

func CreateNodeIdentifierCallSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	callFunctionName := lexer_token.GetTokenIdentifierValue(&t)
	if _, err := parser.ExpectTokens(cparser, []uint16{token.TokenTypeBracketLeft}); err != nil {
		return node.AwooParserNodeResult{}, err
	}
	symbol, err := scope.GetCurrentFunctionSymbol(&cparser.Context.Scopes, callFunctionName)
	if err != nil || symbol.Symbol.Type != types.AwooTypeFunction {
		return node.AwooParserNodeResult{}, parser_error.CreateParserErrorText(parser_error.AwooParserErrorUnknownFunction,
			fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorUnknownFunction], gchalk.Red(callFunctionName)),
			t.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorUnknownFunction])
	}
	callFunctionArguments := symbol.Symbol.Details.(elf.AwooElfSymbolTableEntryFunctionDetails).Arguments

	callNode := node.CreateNodeCall(t)
	for _, arg := range callFunctionArguments {
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
	if len(callFunctionArguments) == 0 {
		if _, err := parser.ExpectToken(cparser, token.TokenTypeBracketRight); err != nil {
			return callNode, err
		}
	}

	return callNode, nil
}

func CreateNodeIdentifierCallSafeFast(cparser *parser.AwooParser) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.ExpectTokens(cparser, []uint16{node.ParserNodeTypeIdentifier})
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierCallSafe(cparser, *t)
}

func CreateNodeIdentifierSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, _ *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
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

func CreateNodeIdentifierSafeFast(cparser *parser.AwooParser, details *parser_details.ConstructExpressionDetails) (node.AwooParserNodeResult, *parser_error.AwooParserError) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier)
	if err != nil {
		return node.AwooParserNodeResult{}, err
	}

	return CreateNodeIdentifierSafe(cparser, *t, details)
}
