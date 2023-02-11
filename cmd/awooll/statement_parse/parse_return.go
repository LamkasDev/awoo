package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

func ConstructStatementReturn(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	if !details.CanReturn {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %s", awerrors.ErrorUnexpectedStatement, gchalk.Red("return"))
	}
	details.CanReturn = false

	// TODO: add return type.
	n, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
		Type:     cparser.Context.Lexer.Types.All[types.AwooTypeInt32],
		EndToken: token.TokenTypeEndStatement,
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	retStatement := statement.CreateStatementReturn(n.Node)

	return retStatement, nil
}
