package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileNodeNegative(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	err := CompileNodeValue(ccompiler, celf, node.GetNodeSingleValue(&n), details)
	if err != nil {
		return err
	}

	return encoder.Encode(celf, instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionSUB,
		Destination: details.Register,
		SourceTwo:   details.Register,
	})
}
