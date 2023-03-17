package statement_parse

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

func ConstructStatementDefinitionVariable(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, details *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	variableTypeNode, err := ConstructNodeType(cparser, t)
	variableType := cparser.Context.Lexer.Types.All[lexer_token.GetTokenTypeId(&t)]
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	definitionStatement := statement.CreateStatementDefinitionVariable(variableTypeNode.Node)

	t, err = parser.ExpectToken(cparser, token.TokenTypeIdentifier)
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	variableNameNode := node.CreateNodeIdentifier(t)
	variableName := node.GetNodeIdentifierValue(&variableNameNode.Node)
	_, err = parser_context.PushParserScopeCurrentBlockMemory(&cparser.Context, parser_context.AwooParserMemoryEntry{
		Name: variableName,
		Type: variableType.Id,
	})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	statement.SetStatementDefinitionVariableIdentifier(&definitionStatement, variableNameNode.Node)

	t, err = parser.ExpectTokens(cparser, []uint16{token.TokenOperatorEq, details.EndToken})
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	if t.Type == token.TokenOperatorEq {
		valueDetails := parser_details.ConstructExpressionDetails{
			Type:      variableType,
			EndTokens: []uint16{details.EndToken},
		}
		valueNode, err := ConstructExpressionStart(cparser, &valueDetails)
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		statement.SetStatementDefinitionVariableValue(&definitionStatement, &valueNode.Node)
	}

	return definitionStatement, nil
}
