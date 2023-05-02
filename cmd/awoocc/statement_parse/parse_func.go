package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func ConstructStatementFunc(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, _ *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier)
	if err != nil {
		return nil, err
	}
	functionNameNode := node.CreateNodeIdentifier(t)
	functionName := lexer_token.GetTokenIdentifierValue(&t)
	functionStatement := statement.CreateStatementFunc(functionNameNode.Node)
	parser_context.PushParserScopeFunction(&cparser.Context, parser_context.AwooParserScopeFunction{
		Name: functionName,
	})

	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketLeft); err != nil {
		return &functionStatement, err
	}
	for argumentToken, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeIdentifier); argumentToken != nil; argumentToken, _ = parser.ExpectTokenOptional(cparser, token.TokenTypeIdentifier) {
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
		_, ok := parser_context.PushParserScopeCurrentBlockSymbol(&cparser.Context, elf.AwooElfSymbolTableEntry{
			Name: argumentName,
			Type: argumentType,
		})
		if !ok {
			return &functionStatement, err
		}
	}
	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketRight); err != nil {
		return &functionStatement, err
	}

	var functionReturnType *commonTypes.AwooTypeId
	if returnTypeToken, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeType); returnTypeToken != nil {
		returnTypeNode, err := ConstructNodeType(cparser, *returnTypeToken)
		if err != nil {
			return &functionStatement, err
		}
		statement.SetStatementFuncReturnType(&functionStatement, &returnTypeNode.Node)
	}

	parser_context.PushParserScopeBlockSymbol(&cparser.Context, parser_context.AwooCompilerGlobalFunctionId, parser_context.AwooCompilerGlobalBlockId, elf.AwooElfSymbolTableEntry{
		Name: functionName,
		Type: commonTypes.AwooTypeFunction,
		Details: elf.AwooElfSymbolTableEntryFunctionDetails{
			ReturnType: functionReturnType,
			Arguments:  statement.GetStatementFuncArguments(&functionStatement),
		},
	})
	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketCurlyLeft); err != nil {
		return &functionStatement, err
	}
	functionBody, err := ConstructStatementGroup(cparser, &parser_details.ConstructStatementDetails{CanReturn: true})
	if err != nil {
		return &functionStatement, err
	}
	statement.SetStatementFuncBody(&functionStatement, *functionBody)

	parser_context.PopParserScopeCurrentFunction(&cparser.Context)

	return &functionStatement, nil
}
