package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileNodeNegative(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeValue(ccompiler, node.GetNodeSingleValue(&n), d, details)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSUB,
		Destination: details.Register,
		SourceTwo:   details.Register,
	}, d)
}
