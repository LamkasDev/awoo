package node

import (
	"fmt"
	"math"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

type AwooParserNodeDataPrimitive struct {
	Value interface{}
}

func GetNodePrimitiveValue(n *AwooParserNode) interface{} {
	return n.Data.(AwooParserNodeDataPrimitive).Value
}

func SetNodePrimitiveValue(n *AwooParserNode, value interface{}) {
	d := n.Data.(AwooParserNodeDataPrimitive)
	d.Value = n
	n.Data = d
}

func CreateNodePrimitiveSafe(primitiveType types.AwooType, t lexer_token.AwooLexerToken) (AwooParserNode, error) {
	// TODO: fix da overflow logic
	v := lexer_token.GetTokenPrimitiveValue(&t).(int64)
	limitBytes := float64(8 * primitiveType.Size)
	if primitiveType.Flags&types.AwooTypeFlagsSign == 1 {
		limitBytes--
	}
	limit := int64(math.Pow(2, limitBytes))
	if v > limit {
		return AwooParserNode{}, fmt.Errorf("primitive overflow (%s > %s)", gchalk.Red(fmt.Sprint(v)), gchalk.Green(fmt.Sprint(limit)))
	}
	if v < -limit {
		return AwooParserNode{}, fmt.Errorf("primitive underflow (%s > %s)", gchalk.Red(fmt.Sprint(v)), gchalk.Green(fmt.Sprint(-limit)))
	}

	return CreateNodePrimitive(t), nil
}

func CreateNodePrimitive(t lexer_token.AwooLexerToken) AwooParserNode {
	return AwooParserNode{
		Type:  ParserNodeTypePrimitive,
		Token: t,
		Data: AwooParserNodeDataPrimitive{
			Value: lexer_token.GetTokenPrimitiveValue(&t),
		},
	}
}
