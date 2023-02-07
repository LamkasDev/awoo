package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

func ConstructStatementReturn(cparser *parser.AwooParser) (statement.AwooParserStatement, error) {
	// TODO: add return type.
	n, err := ConstructExpressionStart(cparser, &ConstructExpressionDetails{Type: cparser.Context.Lexer.Types.All[types.AwooTypeInt32]})
	if err != nil {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %w", awerrors.ErrorFailedToConstructStatement, err)
	}
	retStatement := statement.CreateStatementReturn(n.Node)

	return retStatement, nil
}
