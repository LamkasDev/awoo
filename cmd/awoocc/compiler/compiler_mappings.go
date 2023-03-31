package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

type AwooCompileStatement func(ccompiler *AwooCompiler, celf *elf.AwooElf, s statement.AwooParserStatement) error

type AwooCompileNodeExpression func(ccompiler *AwooCompiler, celf *elf.AwooElf, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) error

type AwooCompileNodeValue func(ccompiler *AwooCompiler, celf *elf.AwooElf, n node.AwooParserNode, details *compiler_details.CompileNodeValueDetails) error

type AwooCompilerMappings struct {
	Statement        map[uint16]AwooCompileStatement
	NodeExpression   map[uint16]AwooCompileNodeExpression
	NodeValue        map[uint16]AwooCompileNodeValue
	InstructionTable instructions.AwooInstructionTable
}
