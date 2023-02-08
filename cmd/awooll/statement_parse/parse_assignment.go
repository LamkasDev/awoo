package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementAssignment(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (statement.AwooParserStatement, error) {
	n, err := CreateNodeIdentifierSafe(cparser, t)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	identifierVariable, ok := parser_context.GetContextVariable(&cparser.Context, node.GetNodeIdentifierValue(&n.Node))
	if !ok {
		return statement.AwooParserStatement{}, err
	}
	asStatement := statement.CreateStatementAssignment(n.Node)
	_, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenOperatorEq}, "=")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	n, err = ConstructExpressionStart(cparser, &ConstructExpressionDetails{Type: cparser.Context.Lexer.Types.All[identifierVariable.Type]})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	statement.SetStatementAssignmentValue(&asStatement, n.Node)

	return asStatement, nil
}
