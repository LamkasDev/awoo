package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

func ConstructStatementReturn(cparser *parser.AwooParser) (statement.AwooParserStatement, error) {
	// TODO: add return type.
	n, err := ConstructExpressionStart(cparser, &ConstructExpressionDetails{Type: cparser.Context.Lexer.Types.All[types.AwooTypeInt32]})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	retStatement := statement.CreateStatementReturn(n.Node)

	return retStatement, nil
}
