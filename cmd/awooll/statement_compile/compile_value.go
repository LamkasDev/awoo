package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/jwalton/gchalk"
)

func CompileNodeValue(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	entry, ok := context.MappingsNodeValue[n.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileNode, gchalk.Red(fmt.Sprintf("%#x", n.Type)))
	}
	d, err := entry(context, n, d, details)
	if err != nil {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorFailedToCompileNode, err)
	}

	return d, nil
}

func CompileNodeValueFast(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeValue(context, n, d, details)
}
