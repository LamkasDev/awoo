package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/jwalton/gchalk"
)

func CompileNodeValue(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	entry, ok := ccompiler.Settings.Mappings.NodeValue[n.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileNode, gchalk.Red(fmt.Sprintf("%#x", n.Type)))
	}

	return entry(ccompiler, n, d, details)
}
