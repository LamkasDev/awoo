package statement_parse

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/jwalton/gchalk"
)

func ConstructExpressionEndStatement(n node.AwooParserNode, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if details.PendingBrackets > 0 {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(")"))
	}
	return node.AwooParserNodeResult{
		Node: n,
		End:  true,
	}, nil
}

func ConstructExpressionEndBracket(n node.AwooParserNode, details *ConstructExpressionDetails) (node.AwooParserNodeResult, error) {
	if details.PendingBrackets > 0 {
		details.PendingBrackets--
		return node.AwooParserNodeResult{
			Node:       n,
			EndBracket: true,
		}, nil
	}
	return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s", awerrors.ErrorUnexpectedToken, gchalk.Red(")"))
}
