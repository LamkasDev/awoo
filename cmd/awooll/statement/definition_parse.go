package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementDefinition(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken) (AwooParserStatement, error) {
	statement := CreateStatementDefinition(node.CreateNodeType(t))
	statementType := context.Lexer.Types.All[lexer_token.GetTokenTypeType(&t)]
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return statement, err
	}
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	if _, ok := parser_context.GetContextVariable(context, identifier); ok {
		return statement, fmt.Errorf("already defined identifier %s", gchalk.Red(identifier))
	}
	SetStatementDefinitionIdentifier(&statement, node.CreateNodeIdentifier(t))
	_, err = ExpectToken(fetchToken, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return statement, err
	}
	n, err := ConstructExpressionFast(context, fetchToken, statementType)
	if err != nil {
		return statement, err
	}
	SetStatementDefinitionValue(&statement, n)
	parser_context.SetContextVariable(context, parser_context.AwooParserContextVariable{
		Name: identifier, Type: statementType.Type,
	})

	return statement, nil
}
