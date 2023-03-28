package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	awooElf "github.com/LamkasDev/awoo-emu/cmd/awoocc/elf"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementAssignmentPointer(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, s statement.AwooParserStatement) error {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	identifierNode = node.GetNodeSingleValue(&identifierNode)
	variableMemory, err := compiler_context.GetCompilerScopeCurrentFunctionMemory(&ccompiler.Context, node.GetNodeIdentifierValue(&identifierNode))
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[*variableMemory.Symbol.TypeDetails]

	valueNode := statement.GetStatementAssignmentValue(&s)
	details := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Symbol.Type,
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if err = CompileNodeValue(ccompiler, elf, valueNode, &details); err != nil {
		return err
	}

	addressRegister := cpu.GetNextTemporaryRegister(details.Register)
	loadInstruction := instruction.AwooInstruction{
		Definition:  *instructions.AwooInstructionsLoad[ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type].Size],
		Destination: addressRegister,
		Immediate:   variableMemory.Symbol.Start,
	}
	if !variableMemory.Global {
		loadInstruction.SourceOne = cpu.AwooRegisterSavedZero
	} else {
		awooElf.PushRelocationEntry(elf, variableMemory.Symbol.Name)
	}
	if err = encoder.Encode(elf, loadInstruction); err != nil {
		return err
	}

	return encoder.Encode(elf, instruction.AwooInstruction{
		Definition: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:  addressRegister,
		SourceTwo:  details.Register,
	})
}
