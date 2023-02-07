package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/jwalton/gchalk"
)

type CompileNodeValueDetails struct {
	Register uint8
}

func CompileNodeValue(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	switch n.Type {
	case node.ParserNodeTypeIdentifier:
		return CompileNodeIdentifier(context, n, d, details)
	case node.ParserNodeTypePrimitive:
		return CompileNodePrimitive(n, d, details)
	case node.ParserNodeTypeExpression:
		return CompileNodeExpression(context, n, d, details)
	case node.ParserNodeTypeNegative:
		return CompileNodeNegative(context, n, d, details)
	case node.ParserNodeTypeReference:
		return CompileNodeReference(context, n, d, details)
	case node.ParserNodeTypeDereference:
		return CompileNodeDereference(context, n, d, details)
	}

	return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileNode, gchalk.Red(fmt.Sprintf("%#x", n.Type)))
}

func CompileNodeValueFast(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte) ([]byte, error) {
	return CompileNodeValue(context, n, d, CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero})
}
