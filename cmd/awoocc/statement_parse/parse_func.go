package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
	"github.com/jwalton/gchalk"
)

func ConstructStatementFunc(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, _ *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier)
	if err != nil {
		return nil, err
	}
	functionNameNode := node.CreateNodeIdentifier(*t)
	functionName := lexer_token.GetTokenIdentifierValue(t)
	functionStatement := statement.CreateStatementFunc(functionNameNode.Node)
	scope.PushFunction(&cparser.Context.Scopes, scope.AwooScopeFunction{
		Name: functionName,
	})

	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketLeft); err != nil {
		return &functionStatement, err
	}
	for argumentToken := parser.ExpectTokenOptional(cparser, token.TokenTypeIdentifier); argumentToken != nil; argumentToken = parser.ExpectTokenOptional(cparser, token.TokenTypeIdentifier) {
		argumentName := lexer_token.GetTokenIdentifierValue(argumentToken)
		argumentTypeNode, err := ConstructNodeTypeFast(cparser)
		if err != nil {
			return &functionStatement, err
		}
		argumentType := node.GetNodeTypeType(&argumentTypeNode.Node)

		// TODO: support pointers
		statement.SetStatementFuncArguments(&functionStatement, append(statement.GetStatementFuncArguments(&functionStatement), elf.AwooElfSymbolTableEntry{
			Name: argumentName,
			Size: cparser.Context.Lexer.Types.All[argumentType].Size,
			Type: argumentType,
		}))
		_, cerr := scope.PushCurrentFunctionSymbol(&cparser.Context.Scopes, elf.AwooElfSymbolTableEntry{
			Name: argumentName,
			Type: argumentType,
		})
		if cerr != nil {
			return &functionStatement, parser_error.CreateParserErrorText(parser_error.AwooParserErrorAlreadyDefinedVariable,
				fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorAlreadyDefinedVariable], gchalk.Red(argumentName)),
				t.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorAlreadyDefinedVariable])
		}
	}
	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketRight); err != nil {
		return &functionStatement, err
	}

	var functionReturnType *commonTypes.AwooTypeId
	if returnTypeToken := parser.ExpectTokenOptional(cparser, token.TokenTypeType); returnTypeToken != nil {
		returnTypeNode, err := ConstructNodeType(cparser, *returnTypeToken)
		if err != nil {
			return &functionStatement, err
		}
		statement.SetStatementFuncReturnType(&functionStatement, &returnTypeNode.Node)
	}

	_, cerr := scope.PushFunctionBlockSymbol(&cparser.Context.Scopes, scope.AwooScopeGlobalFunctionId, scope.AwooScopeGlobalBlockId, elf.AwooElfSymbolTableEntry{
		Name: functionName,
		Type: commonTypes.AwooTypeFunction,
		Details: elf.AwooElfSymbolTableEntryFunctionDetails{
			ReturnType: functionReturnType,
			Arguments:  statement.GetStatementFuncArguments(&functionStatement),
		},
	})
	if cerr != nil {
		return &functionStatement, parser_error.CreateParserError(parser_error.AwooParserErrorAlreadyDefinedFunction, t.Position)
	}
	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketCurlyLeft); err != nil {
		return &functionStatement, err
	}
	functionBody, err := ConstructStatementGroup(cparser, &parser_details.ConstructStatementDetails{CanReturn: true})
	if err != nil {
		return &functionStatement, err
	}
	statement.SetStatementFuncBody(&functionStatement, *functionBody)

	scope.PopCurrentFunction(&cparser.Context.Scopes)

	return &functionStatement, nil
}
