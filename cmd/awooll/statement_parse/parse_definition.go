package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func ConstructStatementDefinitionVariable(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (statement.AwooParserStatement, error) {
	n := node.CreateNodeType(t)
	if n.Error != nil {
		return statement.AwooParserStatement{}, n.Error
	}
	defStatement := statement.CreateStatementDefinitionVariable(n.Node)
	statementType := cparser.Context.Lexer.Types.All[lexer_token.GetTokenTypeId(&t)]
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	identifier := lexer_token.GetTokenIdentifierValue(&t)
	if _, ok := parser_context.GetContextVariable(&cparser.Context, identifier); ok {
		return statement.AwooParserStatement{}, fmt.Errorf("already defined identifier %s", gchalk.Red(identifier))
	}
	n = node.CreateNodeIdentifier(t)
	if n.Error != nil {
		return statement.AwooParserStatement{}, n.Error
	}
	statement.SetStatementDefinitionVariableIdentifier(&defStatement, n.Node)
	_, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	n = ConstructExpressionStart(cparser, &ConstructExpressionDetails{Type: statementType})
	if n.Error != nil {
		return statement.AwooParserStatement{}, n.Error
	}
	statement.SetStatementDefinitionVariableValue(&defStatement, n.Node)
	parser_context.SetContextVariable(&cparser.Context, parser_context.AwooParserContextVariable{
		Name: identifier, Type: statementType.Id,
	})

	return defStatement, nil
}
