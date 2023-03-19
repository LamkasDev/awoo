package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
)

func ConstructStatementReturn(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	returnStatement := statement.CreateStatementReturn(nil)
	currentScopeFunction := cparser.Context.Scopes.Functions[uint16(len(cparser.Context.Scopes.Functions)-1)]
	currentFunction := cparser.Context.Functions.Entries[currentScopeFunction.Name]
	if currentFunction.ReturnType != nil {
		returnValue, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
			Type:      *currentFunction.ReturnType,
			EndTokens: []uint16{details.EndToken},
		})
		if err != nil {
			return returnStatement, err
		}
		statement.SetStatementReturnValue(&returnStatement, &returnValue.Node)
	} else {
		if _, err := parser.ExpectToken(cparser, details.EndToken); err != nil {
			return returnStatement, err
		}
	}

	return returnStatement, nil
}
