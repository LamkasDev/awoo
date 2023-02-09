package compiler_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

type CompileNodeValueDetails struct {
	Register uint8
}

type AwooCompileStatement func(context *AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error)

type AwooCompileNodeValue func(context *AwooCompilerContext, n node.AwooParserNode, d []byte, details *CompileNodeValueDetails) ([]byte, error)

type AwooCompileNodeExpression func(context *AwooCompilerContext, d []byte, leftDetails *CompileNodeValueDetails, rightDetails *CompileNodeValueDetails) ([]byte, error)

// TODO: create a stack class that saves position of stack after all constants have been written
type AwooCompilerContext struct {
	Position               uint16
	Parser                 parser_context.AwooParserContext
	Scopes                 AwooCompilerScopeContainer
	Registers              AwooCompilerRegisters
	Functions              AwooCompilerFunctionContainer
	MappingsStatement      map[uint16]AwooCompileStatement
	MappingsNodeValue      map[uint16]AwooCompileNodeValue
	MappingsNodeExpression map[uint16]AwooCompileNodeExpression
}

func GetProgramHeaderSize() uint16 {
	return 4
}
