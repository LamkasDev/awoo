package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"golang.org/x/exp/constraints"
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

func GetNodePrimitiveValueFormat[K constraints.Integer](context lexer_context.AwooLexerContext, n *AwooParserNode) K {
	primType := context.Types.All[GetNodePrimitiveType(n)]
	switch primType.Size {
	case 1:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int8))
		} else {
			return K(GetNodePrimitiveValue(n).(uint8))
		}
	case 2:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int16))
		} else {
			return K(GetNodePrimitiveValue(n).(uint16))
		}
	case 4:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int32))
		} else {
			return K(GetNodePrimitiveValue(n).(uint32))
		}
	case 8:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int64))
		} else {
			return K(GetNodePrimitiveValue(n).(uint64))
		}
	}

	return K(0)
}

func GetNodePrimitiveValue(n *AwooParserNode) interface{} {
	return n.Data.(AwooParserNodeDataPrimitive).Value
}

func SetNodePrimitiveValue(n *AwooParserNode, value interface{}) {
	d := n.Data.(AwooParserNodeDataPrimitive)
	d.Value = value
	n.Data = d
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
