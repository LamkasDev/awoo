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

func GetPrimitiveLimits(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (int64, int64) {
	primType := cparser.Context.Lexer.Types.All[lexer_token.GetTokenPrimitiveType(&t)]
	primBytes := float64(8 * primType.Size)
	if primType.Flags&types.AwooTypeFlagsSign == 1 {
		return int64(math.Pow(2, primBytes-1)) - 1, -int64(math.Pow(2, primBytes-1))
	} else {
		return int64(math.Pow(2, primBytes)) - 1, 0
	}
}

func CreateNodePrimitiveSafe(cparser *parser.AwooParser, t lexer_token.AwooLexerToken) (node.AwooParserNodeResult, error) {
	primValue := lexer_token.GetTokenPrimitiveValue(&t).(int64)
	primUp, primDown := GetPrimitiveLimits(cparser, t)
	if primValue > primUp {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s > %s", awerrors.ErrorPrimitiveOverflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(primUp)))
	}
	if primValue < primDown {
		return node.AwooParserNodeResult{}, fmt.Errorf("%w: %s < %s", awerrors.ErrorPrimitiveUnderflow, gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(primDown)))
	}

	return node.CreateNodePrimitive(t), nil
}
