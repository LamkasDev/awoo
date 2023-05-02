package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

func ConstructStatementReturn(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (*statement.AwooParserStatement, *parser_error.AwooParserError) {
	returnStatement := statement.CreateStatementReturn(nil)
	currentScopeFunction := cparser.Context.Scopes.Functions[uint16(len(cparser.Context.Scopes.Functions)-1)]
	currentPrototypeFunction, ok := parser_context.GetParserScopeBlockSymbol(&cparser.Context, parser_context.AwooCompilerGlobalFunctionId, parser_context.AwooCompilerGlobalBlockId, currentScopeFunction.Name)
	if !ok {
		panic("weh")
	}
	currentFunctionReturnType := currentPrototypeFunction.Details.(elf.AwooElfSymbolTableEntryFunctionDetails).ReturnType
	if currentFunctionReturnType != nil {
		returnValue, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
			Type:      *currentFunctionReturnType,
			EndTokens: []uint16{details.EndToken},
		})
		if err != nil {
			return &returnStatement, err
		}
		statement.SetStatementReturnValue(&returnStatement, &returnValue.Node)
	} else {
		if _, err := parser.ExpectToken(cparser, details.EndToken); err != nil {
			return &returnStatement, err
		}
	}

	return &returnStatement, nil
}
