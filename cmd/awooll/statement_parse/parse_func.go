package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementFunc(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, _ *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier, "identifier")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	functionNameNode := node.CreateNodeIdentifier(t)
	functionName := lexer_token.GetTokenIdentifierValue(&t)
	functionStatement := statement.CreateStatementFunc(functionNameNode.Node)
	parser_context.PushParserScopeFunction(&cparser.Context, parser_context.AwooParserScopeFunction{
		Name: functionName,
	})

	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketLeft, "("); err != nil {
		return functionStatement, err
	}
	for argumentToken, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeIdentifier); argumentToken != nil; argumentToken, _ = parser.ExpectTokenOptional(cparser, token.TokenTypeIdentifier) {
		argumentName := lexer_token.GetTokenIdentifierValue(argumentToken)
		argumentTypeNode, err := ConstructNodeTypeFast(cparser)
		if err != nil {
			return functionStatement, err
		}
		argumentType := node.GetNodeTypeType(&argumentTypeNode.Node)

		// TODO: support pointers
		statement.SetStatementFuncArguments(&functionStatement, append(statement.GetStatementFuncArguments(&functionStatement), statement.AwooParserStatementFuncArgument{
			Name: argumentName,
			Size: cparser.Context.Lexer.Types.All[argumentType].Size,
			Type: argumentType,
		}))
		_, err = parser_context.PushParserScopeCurrentBlockMemory(&cparser.Context, parser_context.AwooParserMemoryEntry{
			Name: argumentName,
			Type: argumentType,
		})
		if err != nil {
			return functionStatement, err
		}
	}
	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketRight, ")"); err != nil {
		return functionStatement, err
	}

	var functionReturnType *uint16
	if returnTypeToken, _ := parser.ExpectTokenOptional(cparser, token.TokenTypeType); returnTypeToken != nil {
		returnTypeNode := ConstructNodeType(cparser, *returnTypeToken)
		statement.SetStatementFuncReturnType(&functionStatement, &returnTypeNode.Node)
	}

	if _, err = parser.ExpectToken(cparser, token.TokenTypeBracketCurlyLeft, "{"); err != nil {
		return functionStatement, err
	}
	functionBody, err := ConstructStatementGroup(cparser, &parser_details.ConstructStatementDetails{CanReturn: true})
	if err != nil {
		return functionStatement, err
	}
	statement.SetStatementFuncBody(&functionStatement, functionBody)

	parser_context.PopParserScopeCurrentFunction(&cparser.Context)
	parser_context.PushParserFunction(&cparser.Context, parser_context.AwooParserFunction{
		Name:       functionName,
		ReturnType: functionReturnType,
		Arguments:  statement.GetStatementFuncArguments(&functionStatement),
	})

	return functionStatement, nil
}
