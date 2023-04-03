package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
)

func ConstructStatementDefinitionType(cparser *parser.AwooParser, _ lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, *parser_error.AwooParserError) {
	t, err := parser.ExpectToken(cparser, token.TokenTypeIdentifier)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	newIdentifier := lexer_token.GetTokenIdentifierValue(&t)
	newType := types.AwooType{
		Key: newIdentifier,
	}
	newIdentifierNode := node.CreateNodeIdentifier(t)
	defStatement := statement.CreateStatementDefinitionType(newIdentifierNode.Node)

	// TODO: this is a placeholder i promise.
	t, err = parser.ExpectToken(cparser, token.TokenTypeType)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	originalIdentifier := lexer_token.GetTokenTypeId(&t)
	originalType, _ := lexer_context.GetContextTypeId(&cparser.Context.Lexer, originalIdentifier)
	originalIdentifierNode, err := ConstructNodeType(cparser, t)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	statement.SetStatementDefinitionTypeValue(&defStatement, originalIdentifierNode.Node)
	lexer_context.AddContextTypeAlias(&cparser.Context.Lexer, originalType, newType)
	// end
	if _, err = parser.ExpectToken(cparser, details.EndToken); err != nil {
		return statement.AwooParserStatement{}, err
	}

	return defStatement, nil
}
