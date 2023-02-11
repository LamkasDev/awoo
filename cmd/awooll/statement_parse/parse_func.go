package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementFunc(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, _ *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	t, err := parser.ExpectTokenParser(cparser, token.TokenTypeIdentifier, "identifier")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	if _, ok := parser_context.GetContextFunction(&cparser.Context, identifier); ok {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %s", awerrors.ErrorAlreadyDefinedFunction, gchalk.Red(identifier))
	}
	identifierNode := node.CreateNodeIdentifier(t)
	funcStatement := statement.CreateStatementFunc(identifierNode.Node)
	contextFunc := parser_context.AwooParserContextFunction{
		Name:      identifier,
		Arguments: []parser_context.AwooParserContextVariable{},
	}
	if _, err = parser.ExpectTokenParser(cparser, token.TokenTypeBracketLeft, "("); err != nil {
		return funcStatement, err
	}
	for t, ok := parser.PeekParser(cparser); ok && t.Type == token.TokenTypeIdentifier; t, ok = parser.PeekParser(cparser) {
		parser.AdvanceParser(cparser)
		argumentIdentifier := lexer_token.GetTokenIdentifierValue(&t)
		t, err = parser.ExpectTokenParser(cparser, token.TokenTypeType, "type")
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		argumentType := lexer_token.GetTokenTypeId(&t)
		statement.SetStatementFuncArguments(&funcStatement, append(statement.GetStatementFuncArguments(&funcStatement), statement.AwooParserStatementFuncArgument{
			Name: argumentIdentifier,
			Size: cparser.Context.Lexer.Types.All[argumentType].Size,
			Type: argumentType,
		}))
		// TODO: setup a proper scoped system for variables.
		contextArg := parser_context.AwooParserContextVariable{
			Name: argumentIdentifier, Type: argumentType,
		}
		parser_context.SetContextVariable(&cparser.Context, contextArg)
		contextFunc.Arguments = append(contextFunc.Arguments, contextArg)
	}
	if _, err = parser.ExpectTokenParser(cparser, token.TokenTypeBracketRight, ")"); err != nil {
		return funcStatement, err
	}
	if _, err = parser.ExpectTokenParser(cparser, token.TokenTypeBracketCurlyLeft, "{"); err != nil {
		return funcStatement, err
	}
	funcGroup, err := ConstructStatementGroup(cparser, &parser_details.ConstructStatementDetails{CanReturn: true})
	if err != nil {
		return funcStatement, err
	}
	statement.SetStatementFuncBody(&funcStatement, funcGroup)
	parser_context.SetContextFunction(&cparser.Context, contextFunc)

	return funcStatement, nil
}
