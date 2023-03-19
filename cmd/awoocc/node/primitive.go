package node

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
	"golang.org/x/exp/constraints"
)

type AwooParserNodeDataPrimitive struct {
	Type  commonTypes.AwooTypePrimitiveId
	Value interface{}
}

func GetNodePrimitiveType(n *AwooParserNode) commonTypes.AwooTypePrimitiveId {
	return n.Data.(AwooParserNodeDataPrimitive).Type
}

func SetNodePrimitiveType(n *AwooParserNode, t commonTypes.AwooTypePrimitiveId) {
	d := n.Data.(AwooParserNodeDataPrimitive)
	d.Type = t
	n.Data = d
}

func GetNodePrimitiveValueFormat[K constraints.Integer](context lexer_context.AwooLexerContext, n *AwooParserNode) K {
	primType := context.Types.All[commonTypes.AwooTypeId(GetNodePrimitiveType(n))]
	switch primType.Size {
	case 1:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int8))
		}
		return K(GetNodePrimitiveValue(n).(uint8))
	case 2:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int16))
		}
		return K(GetNodePrimitiveValue(n).(uint16))
	case 4:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int32))
		}
		return K(GetNodePrimitiveValue(n).(uint32))
	case 8:
		if primType.Flags&types.AwooTypeFlagsSign == 1 {
			return K(GetNodePrimitiveValue(n).(int64))
		}
		return K(GetNodePrimitiveValue(n).(uint64))
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
