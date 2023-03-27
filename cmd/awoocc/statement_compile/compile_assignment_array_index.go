package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementAssignmentArrayIndex(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeArrayIndexIdentifier(&identifierNode))
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type]

	valueNode := statement.GetStatementAssignmentValue(&s)
	valueDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Symbol.Type,
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if err = CompileNodeValue(ccompiler, elf, valueNode, &valueDetails); err != nil {
		return err
	}

	addressDetails := compiler_details.CompileNodeValueDetails{
		Register: cpu.GetNextTemporaryRegister(valueDetails.Register),
	}
	if err = CompileArrayIndexAddress(ccompiler, elf, identifierNode, &addressDetails); err != nil {
		return err
	}
	if !variableMemory.Global {
		addressAdjustmentInstruction := encoder.AwooEncodedInstruction{
			Instruction: instructions.AwooInstructionADD,
			SourceOne:   addressDetails.Register,
			SourceTwo:   cpu.AwooRegisterSavedZero,
			Destination: addressDetails.Register,
		}
		if err = encoder.Encode(elf, addressAdjustmentInstruction); err != nil {
			return err
		}
	}

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:   addressDetails.Register,
		SourceTwo:   valueDetails.Register,
		Immediate:   uint32(variableMemory.Symbol.Start),
	}
	return encoder.Encode(elf, saveInstruction)
}
