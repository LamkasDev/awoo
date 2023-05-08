package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooCompilerIfBodiesDescriptor struct {
	Jump   arch.AwooRegister
	Bodies []AwooCompilerIfBodyDescriptor
}

type AwooCompilerIfBodyDescriptor struct {
	Start  arch.AwooRegister
	Length arch.AwooRegister
}

func CompileStatementIfNode(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, descriptor *AwooCompilerIfBodiesDescriptor, s statement.AwooParserStatement) error {
	start := arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))
	switch s.Type {
	case statement.ParserStatementTypeIf:
		// TODO: this could be optimized using top level comparison from value node (because the below instruction can compare).
		valueNode := statement.GetStatementIfValue(&s)
		valueDetails := compiler_details.CompileNodeValueDetails{
			Type:     types.AwooTypeBoolean,
			Register: cpu.AwooRegisterTemporaryZero,
		}
		if err := CompileNodeValue(ccompiler, celf, valueNode, &valueDetails); err != nil {
			return err
		}

		// Reserve instruction for jump to next block.
		jumpToNextBlockStart := arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))
		if err := encoder.Encode(celf, instruction.AwooInstruction{}); err != nil {
			return err
		}

		// Compile the body to determine comparison jump address.
		scope.PushCurrentFunctionBlock(&ccompiler.Context.Scopes, scope.NewScopeBlock("if"))
		if err := CompileStatementGroup(ccompiler, celf, statement.GetStatementIfBody(&s)); err != nil {
			return err
		}
		scope.PopCurrentFunctionBlock(&ccompiler.Context.Scopes)
		bodyEnd := arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))
		bodyLength := bodyEnd - jumpToNextBlockStart

		// Populate reserved instruction with jump.
		jumpToNextBlockInstruction := instruction.AwooInstruction{
			Definition: instructions.AwooInstructionBEQ,
			SourceOne:  valueDetails.Register,
			Immediate:  bodyLength + 8,
		}
		if err := encoder.EncodeAt(celf, jumpToNextBlockStart, jumpToNextBlockInstruction); err != nil {
			return err
		}
	case statement.ParserStatementTypeGroup:
		scope.PushCurrentFunctionBlock(&ccompiler.Context.Scopes, scope.NewScopeBlock("else"))
		if err := CompileStatementGroup(ccompiler, celf, s); err != nil {
			return err
		}
		scope.PopCurrentFunctionBlock(&ccompiler.Context.Scopes)
	}

	// Reserve instruction for jump beyond subsequent blocks.
	if err := encoder.Encode(celf, instruction.AwooInstruction{}); err != nil {
		return err
	}
	end := arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))
	length := end - start
	descriptor.Bodies = append(descriptor.Bodies, AwooCompilerIfBodyDescriptor{
		Start:  start,
		Length: length,
	})
	descriptor.Jump += length

	return nil
}

func CompileStatementIf(ccompiler *compiler.AwooCompiler, celf *elf.AwooElf, ifNode statement.AwooParserStatement) error {
	descriptor := AwooCompilerIfBodiesDescriptor{
		Jump:   arch.AwooRegister(4),
		Bodies: []AwooCompilerIfBodyDescriptor{},
	}
	if err := CompileStatementIfNode(ccompiler, celf, &descriptor, ifNode); err != nil {
		return err
	}
	elseGroups := statement.GetStatementIfElse(&ifNode)
	for _, elseGroup := range elseGroups {
		if err := CompileStatementIfNode(ccompiler, celf, &descriptor, elseGroup); err != nil {
			return err
		}
	}

	// This is used for jumping over subsequent if statements.
	// An extra instruction is added on end of each block, except for the last.
	for _, body := range descriptor.Bodies {
		// TODO: remove jump on last
		descriptor.Jump -= body.Length + 4
		jumpBeyondSubsequentBlocksInstruction := instruction.AwooInstruction{
			Definition: instructions.AwooInstructionJAL,
			Immediate:  descriptor.Jump,
		}
		if err := encoder.EncodeAt(celf, (body.Start+body.Length)-4, jumpBeyondSubsequentBlocksInstruction); err != nil {
			return err
		}
	}

	return nil
}
