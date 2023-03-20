package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/jwalton/gchalk"
)

func HandleNodeExpressionLeftRight(ins instruction.AwooInstructionDefinition) compiler.AwooCompileNodeExpression {
	return func(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) error {
		return encoder.Encode(elf, encoder.AwooEncodedInstruction{
			Instruction: ins,
			SourceOne:   leftDetails.Register,
			SourceTwo:   rightDetails.Register,
			Destination: leftDetails.Register,
		})
	}
}

func HandleNodeExpressionRightLeft(ins instruction.AwooInstructionDefinition) compiler.AwooCompileNodeExpression {
	return func(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) error {
		return encoder.Encode(elf, encoder.AwooEncodedInstruction{
			Instruction: ins,
			SourceOne:   rightDetails.Register,
			SourceTwo:   leftDetails.Register,
			Destination: leftDetails.Register,
		})
	}
}

func CompileNodeExpressionEqEq(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) error {
	err := HandleNodeExpressionLeftRight(instructions.AwooInstructionSUB)(ccompiler, elf, leftDetails, rightDetails)
	if err != nil {
		return err
	}
	return encoder.Encode(elf, encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionSLTIU,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	})
}

func CompileNodeExpressionNotEq(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) error {
	err := HandleNodeExpressionLeftRight(instructions.AwooInstructionSUB)(ccompiler, elf, leftDetails, rightDetails)
	if err != nil {
		return err
	}
	return encoder.Encode(elf, encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionSLTU,
		SourceOne:   cpu.AwooRegisterZero,
		SourceTwo:   leftDetails.Register,
		Destination: leftDetails.Register,
	})
}

func CompileNodeExpressionLTEQ(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) error {
	err := HandleNodeExpressionRightLeft(instructions.AwooInstructionSLT)(ccompiler, elf, leftDetails, rightDetails)
	if err != nil {
		return err
	}
	return encoder.Encode(elf, encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionXORI,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	})
}

func CompileNodeExpressionGTEQ(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) error {
	err := HandleNodeExpressionLeftRight(instructions.AwooInstructionSLT)(ccompiler, elf, leftDetails, rightDetails)
	if err != nil {
		return err
	}
	return encoder.Encode(elf, encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionXORI,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	})
}

func CompileNodeExpression(ccompiler *compiler.AwooCompiler, elf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error {
	entry, ok := ccompiler.Settings.Mappings.NodeExpression[n.Token.Type]
	if !ok {
		return fmt.Errorf("%w: %s", awerrors.ErrorCantCompileOperator, gchalk.Red(ccompiler.Settings.Parser.Lexer.Tokens.All[n.Token.Type].Name))
	}

	var err error
	left := node.GetNodeExpressionLeft(&n)
	leftDetails := compiler_details.CompileNodeValueDetails{
		Type:     details.Type,
		Register: details.Register,
	}
	if err = CompileNodeValue(ccompiler, elf, left, &leftDetails); err != nil {
		return err
	}
	right := node.GetNodeExpressionRight(&n)
	rightDetails := compiler_details.CompileNodeValueDetails{
		Type:     details.Type,
		Register: cpu.GetNextTemporaryRegister(details.Register),
	}
	if err = CompileNodeValue(ccompiler, elf, right, &rightDetails); err != nil {
		return err
	}

	return entry(ccompiler, elf, &leftDetails, &rightDetails)
}
