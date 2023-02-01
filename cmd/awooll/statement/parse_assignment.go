package statement

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementAssignment(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken) (AwooParserStatement, error) {
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	identifierVariable, ok := parser_context.GetContextVariable(context, identifier)
	if !ok {
		return AwooParserStatement{}, fmt.Errorf("unknown identifier %s", gchalk.Red(identifier))
	}
	n := node.CreateNodeIdentifier(t)
	if n.Error != nil {
		return AwooParserStatement{}, n.Error
	}
	statement := CreateStatementAssignment(n.Node)
	_, err := ExpectToken(fetchToken, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return AwooParserStatement{}, err
	}
	n = ConstructExpressionStart(context, fetchToken, &ConstructExpressionDetails{Type: context.Lexer.Types.All[identifierVariable.Type]})
	if n.Error != nil {
		return AwooParserStatement{}, n.Error
	}
	SetStatementAssignmentValue(&statement, n.Node)

	return statement, nil
}
