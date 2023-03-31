package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/jwalton/gchalk"
)

func CompileNodeValue(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	entry, ok := ccompiler.Settings.Mappings.NodeValue[n.Type]
	if !ok {
		return fmt.Errorf("%w: %s", awerrors.ErrorCantCompileNode, gchalk.Red(fmt.Sprintf("%#x", n.Type)))
	}

	return entry(ccompiler, celf, n, details)
}
