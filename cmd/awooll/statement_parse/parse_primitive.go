package statement_parse

import (
	"fmt"
	"math"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

func CreateNodePrimitiveSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	primType := cparser.Context.Lexer.Types.All[lexer_token.GetTokenPrimitiveType(&t)]
	primValue := lexer_token.GetTokenPrimitiveValue(&t).(int64)
	primBytes := float64(8 * primType.Size)
	if primType.Flags&types.AwooTypeFlagsSign == 1 {
		up := int64(math.Pow(2, primBytes-1)) - 1
		if primValue > up {
			return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s > %s", awerrors.ErrorPrimitiveOverflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(up)))
		}
		down := -int64(math.Pow(2, primBytes-1))
		if primValue < down {
			return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s < %s", awerrors.ErrorPrimitiveUnderflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(down)))
		}
	} else {
		up := int64(math.Pow(2, primBytes)) - 1
		if primValue > up {
			return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s > %s", awerrors.ErrorPrimitiveOverflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(up)))
		}
		if primValue < 0 {
			return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s < %s", awerrors.ErrorPrimitiveUnderflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green("0"))
		}
	}

	return node.CreateNodePrimitive(t), nil
}
