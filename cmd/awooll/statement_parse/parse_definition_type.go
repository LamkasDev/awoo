package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

func ConstructStatementDefinitionType(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (statement.AwooParserStatement, error) {
	t, err := parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeIdentifier}, "identifier")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	newIdentifier := lexer_token.GetTokenIdentifierValue(&t)
	newType := types.AwooType{
		Key: newIdentifier,
	}
	newIdentifierNode := node.CreateNodeIdentifier(t)
	if newIdentifierNode.Error != nil {
		return statement.AwooParserStatement{}, newIdentifierNode.Error
	}
	defStatement := statement.CreateStatementDefinitionType(newIdentifierNode.Node)

	// TODO: this is a placeholder i promise
	t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeType}, "type")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	originalIdentifier := lexer_token.GetTokenTypeId(&t)
	originalType, _ := lexer_context.GetContextTypeId(&cparser.Context.Lexer, originalIdentifier)
	originalIdentifierNode := node.CreateNodeType(t)
	if originalIdentifierNode.Error != nil {
		return statement.AwooParserStatement{}, originalIdentifierNode.Error
	}
	statement.SetStatementDefinitionTypeValue(&defStatement, originalIdentifierNode.Node)
	lexer_context.AddContextTypeAlias(&cparser.Context.Lexer, originalType, newType)
	// end

	t, err = parser.ExpectTokenParser(cparser, []uint16{token.TokenTypeEndStatement}, ";")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}

	return defStatement, nil
}
