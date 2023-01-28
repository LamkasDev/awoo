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
	n := node.CreateNodeType(t)
	if n.Error != nil {
		return AwooParserStatement{}, n.Error
	}
	statement := CreateStatementDefinition(n.Node)
	statementType := context.Lexer.Types.All[lexer_token.GetTokenTypeType(&t)]
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return AwooParserStatement{}, err
	}
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	if _, ok := parser_context.GetContextVariable(context, identifier); ok {
		return AwooParserStatement{}, fmt.Errorf("already defined identifier %s", gchalk.Red(identifier))
	}
	n = node.CreateNodeIdentifier(t)
	if n.Error != nil {
		return AwooParserStatement{}, n.Error
	}
	SetStatementDefinitionIdentifier(&statement, n.Node)
	_, err = ExpectToken(fetchToken, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return AwooParserStatement{}, err
	}
	n = ConstructExpressionNegativeFast(context, fetchToken, &ConstructExpressionDetails{Type: statementType})
	if n.Error != nil {
		return AwooParserStatement{}, n.Error
	}
	SetStatementDefinitionValue(&statement, n.Node)
	parser_context.SetContextVariable(context, parser_context.AwooParserContextVariable{
		Name: identifier, Type: statementType.Type,
	})

	return statement, nil
}
