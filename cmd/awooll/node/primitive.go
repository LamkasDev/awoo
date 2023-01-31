package node

import (
	"fmt"
	"math"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/jwalton/gchalk"
)

type AwooParserNodeDataPrimitive struct {
	Type  uint16
	Value interface{}
}

func GetNodePrimitiveType(n *AwooParserNode) uint16 {
	return n.Data.(AwooParserNodeDataPrimitive).Type
}

func SetNodePrimitiveType(n *AwooParserNode, t uint16) {
	d := n.Data.(AwooParserNodeDataPrimitive)
	d.Type = t
	n.Data = d
}

func GetNodePrimitiveValue(n *AwooParserNode) interface{} {
	return n.Data.(AwooParserNodeDataPrimitive).Value
}

func SetNodePrimitiveValue(n *AwooParserNode, value interface{}) {
	d := n.Data.(AwooParserNodeDataPrimitive)
	d.Value = n
	n.Data = d
}

func CreateNodePrimitiveSafe(context *parser_context.AwooParserContext, t lexer_token.AwooLexerToken) AwooParserNodeResult {
	primType := context.Lexer.Types.All[lexer_token.GetTokenPrimitiveType(&t)]
	primValue := lexer_token.GetTokenPrimitiveValue(&t).(int64)
	primBytes := float64(8 * primType.Size)
	if primType.Flags&types.AwooTypeFlagsSign == 1 {
		up := int64(math.Pow(2, primBytes-1)) - 1
		if primValue > up {
			return AwooParserNodeResult{
				Error: fmt.Errorf("primitive overflow (%s > %s)", gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(up))),
			}
		}
		down := -int64(math.Pow(2, primBytes-1))
		if primValue < down {
			return AwooParserNodeResult{
				Error: fmt.Errorf("primitive underflow (%s < %s)", gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(down))),
			}
		}
	} else {
		up := int64(math.Pow(2, primBytes)) - 1
		if primValue > up {
			return AwooParserNodeResult{
				Error: fmt.Errorf("primitive overflow (%s > %s)", gchalk.Red(fmt.Sprint(primValue)), gchalk.Green(fmt.Sprint(up))),
			}
		}
		if primValue < 0 {
			return AwooParserNodeResult{
				Error: fmt.Errorf("primitive underflow (%s < %s)", gchalk.Red(fmt.Sprint(primValue)), gchalk.Green("0")),
			}
		}
	}

	return CreateNodePrimitive(t)
}

func CreateNodePrimitive(t lexer_token.AwooLexerToken) AwooParserNodeResult {
	return AwooParserNodeResult{
		Node: AwooParserNode{
			Type:  ParserNodeTypePrimitive,
			Token: t,
			Data: AwooParserNodeDataPrimitive{
				Type:  lexer_token.GetTokenPrimitiveType(&t),
				Value: lexer_token.GetTokenPrimitiveValue(&t),
			},
		},
	}
}
