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

func CompileStatementAssignmentIdentifier(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeIdentifierValue(&identifierNode))
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type]

	valueNode := statement.GetStatementAssignmentValue(&s)
	valueDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Symbol.Type,
		Register: cpu.AwooRegisterTemporaryZero,
		Address: compiler_details.CompileNodeValueDetailsAddress{
			Immediate: variableMemory.Symbol.Start,
		},
	}
	if !variableMemory.Global {
		valueDetails.Address.Register = cpu.AwooRegisterSavedZero
	}
	if err = CompileNodeValue(ccompiler, elf, valueNode, &valueDetails); err != nil {
		return err
	}
	if valueDetails.Address.Used {
		return nil
	}

	saveInstruction := encoder.AwooEncodedInstruction{
		Instruction: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:   valueDetails.Address.Register,
		SourceTwo:   valueDetails.Register,
		Immediate:   valueDetails.Address.Immediate,
	}
	return encoder.Encode(elf, saveInstruction)
}
