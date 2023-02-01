package statement

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

func ConstructStatementDefinitionType(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken, fetchToken lexer_token.FetchToken) (AwooParserStatement, error) {
	t, err := ExpectToken(fetchToken, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return AwooParserStatement{}, err
	}
	newIdentifier := lexer_token.GetTokenIdentifierValue(&t)
	newType := types.AwooType{
		Key: newIdentifier,
	}
	newIdentifierNode := node.CreateNodeIdentifier(t)
	if newIdentifierNode.Error != nil {
		return AwooParserStatement{}, newIdentifierNode.Error
	}
	statement := CreateStatementDefinitionType(newIdentifierNode.Node)

	// TODO: this is a placeholder i promise
	t, err = ExpectToken(fetchToken, []uint16{token.TokenTypeType}, "type")
	if err != nil {
		return AwooParserStatement{}, err
	}
	originalIdentifier := lexer_token.GetTokenTypeId(&t)
	originalType, _ := lexer_context.GetContextTypeId(&context.Lexer, originalIdentifier)
	originalIdentifierNode := node.CreateNodeType(t)
	if originalIdentifierNode.Error != nil {
		return AwooParserStatement{}, originalIdentifierNode.Error
	}
	SetStatementDefinitionTypeValue(&statement, originalIdentifierNode.Node)
	lexer_context.AddContextTypeAlias(&context.Lexer, originalType, newType)
	// end

	t, err = ExpectToken(fetchToken, []uint16{token.TokenTypeEndStatement}, ";")
	if err != nil {
		return AwooParserStatement{}, err
	}

	return statement, nil
}
