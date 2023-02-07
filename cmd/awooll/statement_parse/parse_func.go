package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementFunc(cparser *parser.AwooParser) (statement.AwooParserStatement, error) {
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	if _, ok := parser_context.GetContextVariable(&cparser.Context, identifier); ok {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %s", awerrors.ErrorAlreadyDefinedVariable, gchalk.Red(identifier))
	}
	identifierNode := node.CreateNodeIdentifier(t)
	funcStatement := statement.CreateStatementFunc(identifierNode.Node)
	t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeBracketLeft}, "(")
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeBracketRight}, ")")
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeBracketCurlyLeft}, "{")
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	funcGroup, err := ConstructStatementGroup(cparser, &ConstructStatementDetails{CanReturn: true})
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	statement.SetStatementFuncBody(&funcStatement, funcGroup)
	parser_context.SetContextFunction(&cparser.Context, parser_context.AwooParserContextFunction{
		Name: identifier,
	})

	return funcStatement, nil
}
