package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

func ConstructStatementDefinitionVariable(cparser *parser.AwooParser, t lexer_token.AwooLexerToken, _ *parser_details.ConstructStatementDetails) (statement.AwooParserStatement, error) {
	variableTypeNode := ConstructNodeType(cparser, t)
	variableType := cparser.Context.Lexer.Types.All[lexer_token.GetTokenTypeId(&t)]
	defStatement := statement.CreateStatementDefinitionVariable(variableTypeNode.Node)
	t, err := parser.ExpectTokenParser(cparser, token.TokenTypeIdentifier, "identifier")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	variableNameNode := node.CreateNodeIdentifier(t)
	variableName := lexer_token.GetTokenIdentifierValue(&t)
	if _, ok := parser_context.GetContextVariable(&cparser.Context, variableName); ok {
		return statement.AwooParserStatement{}, fmt.Errorf("%w: %s", awerrors.ErrorAlreadyDefinedVariable, gchalk.Red(variableName))
	}
	statement.SetStatementDefinitionVariableIdentifier(&defStatement, variableNameNode.Node)
	t, err = parser.ExpectTokensParser(cparser, []uint16{token.TokenOperatorEq, token.TokenTypeEndStatement}, "= or ;")
	if err != nil {
		return statement.AwooParserStatement{}, err
	}
	if t.Type == token.TokenOperatorEq {
		variableValueNode, err := ConstructExpressionStart(cparser, &parser_details.ConstructExpressionDetails{
			Type:     variableType,
			EndToken: token.TokenTypeEndStatement,
		})
		if err != nil {
			return statement.AwooParserStatement{}, err
		}
		statement.SetStatementDefinitionVariableValue(&defStatement, variableValueNode.Node)
	} else {
		// TODO: create set for uninitialized nodes
		variableValueNode := node.CreateNodePrimitive(lexer_token.CreateTokenPrimitive(0, types.AwooTypeInt64, int64(0), nil))
		statement.SetStatementDefinitionVariableValue(&defStatement, variableValueNode.Node)
	}
	parser_context.SetContextVariable(&cparser.Context, parser_context.AwooParserContextVariable{
		Name: variableName, Type: variableType.Id,
	})

	return defStatement, nil
}
