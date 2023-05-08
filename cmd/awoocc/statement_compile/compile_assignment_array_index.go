package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileStatementAssignmentArrayIndex(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error {
	identifierNode := statement.GetStatementAssignmentIdentifier(&s)
	variableMemory, err := scope.GetCurrentFunctionSymbol(&ccompiler.Context.Scopes, node.GetNodeArrayIndexIdentifier(&identifierNode))
	if err != nil {
		return err
	}
	variableType := ccompiler.Context.Parser.Lexer.Types.All[variableMemory.Symbol.Type]

	valueNode := statement.GetStatementAssignmentValue(&s)
	valueDetails := compiler_details.CompileNodeValueDetails{
		Type:     variableMemory.Symbol.Type,
		Register: cpu.AwooRegisterTemporaryZero,
	}
	if err = CompileNodeValue(ccompiler, celf, valueNode, &valueDetails); err != nil {
		return err
	}

	addressDetails := compiler_details.CompileNodeValueDetails{
		Register: cpu.GetNextTemporaryRegister(valueDetails.Register),
	}
	if err = CompileArrayIndexAddress(ccompiler, celf, identifierNode, &addressDetails); err != nil {
		return err
	}
	if !variableMemory.Global {
		addressAdjustmentInstruction := instruction.AwooInstruction{
			Definition:  instructions.AwooInstructionADD,
			SourceOne:   addressDetails.Register,
			SourceTwo:   cpu.AwooRegisterSavedZero,
			Destination: addressDetails.Register,
		}
		if err = encoder.Encode(celf, addressAdjustmentInstruction); err != nil {
			return err
		}
	}

	saveInstruction := instruction.AwooInstruction{
		Definition: *instructions.AwooInstructionsSave[variableType.Size],
		SourceOne:  addressDetails.Register,
		SourceTwo:  valueDetails.Register,
		Immediate:  variableMemory.Symbol.Start,
	}
	if variableMemory.Global {
		elf.PushRelocationEntry(celf, variableMemory.Symbol.Name)
	}
	return encoder.Encode(celf, saveInstruction)
}
