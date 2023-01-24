package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementDefinition(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken ConstructStatementFetchToken) (AwooParserStatement, error) {
	statementType := context.Lexer.Types.All[lexer_token.GetTokenTypeType(&t)]
	statement := CreateStatementDefinition(node.CreateNodeType(t))
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return statement, err
	}
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	if _, ok := parser_context.GetContextVariable(context, identifier); ok {
		return statement, fmt.Errorf("already defined identifier %s", gchalk.Red(identifier))
	}
	parser_context.SetContextVariable(context, parser_context.AwooParserContextVariable{
		Name: identifier,
	})
	SetStatementDefinitionIdentifier(&statement, node.CreateNodeIdentifier(t))
	_, err = ExpectToken(fetchToken, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return statement, err
	}
	t, err = ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive}, "primitive")
	if err != nil {
		return statement, err
	}
	primitiveNode, err := node.CreateNodePrimitiveSafe(statementType, t)
	if err != nil {
		return statement, err
	}
	SetStatementDefinitionValue(&statement, primitiveNode)
	_, err = ExpectToken(fetchToken, []uint16{token.TokenOperatorEndStatement}, ";")
	if err != nil {
		return statement, err
	}

	return statement, nil
}
