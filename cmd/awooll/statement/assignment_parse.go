package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementAssignment(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken ConstructStatementFetchToken) (AwooParserStatement, error) {
	statement := CreateStatementAssignment(node.CreateNodeIdentifier(t))
	_, err := ExpectToken(fetchToken, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return statement, err
	}
	t, err = ExpectToken(fetchToken, []uint16{token.TokenTypePrimitive}, "primitive")
	if err != nil {
		return statement, err
	}
	SetStatementAssignmentValue(&statement, node.CreateNodePrimitive(t))
	_, err = ExpectToken(fetchToken, []uint16{token.TokenOperatorEndStatement}, ";")
	if err != nil {
		return statement, err
	}

	return statement, nil
}
