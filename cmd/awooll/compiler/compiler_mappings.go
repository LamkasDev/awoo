package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

type AwooCompileStatement func(ccompiler *AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error)

type AwooCompileNodeExpression func(ccompiler *AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error)

type AwooCompileNodeValue func(ccompiler *AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error)

type AwooCompilerMappings struct {
	Statement      map[uint16]AwooCompileStatement
	NodeExpression map[uint16]AwooCompileNodeExpression
	NodeValue      map[uint16]AwooCompileNodeValue
}
