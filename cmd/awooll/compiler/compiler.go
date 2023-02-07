package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement_compile"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
)

type AwooCompiler struct {
	Contents parser.AwooParserResult
	Length   uint16
	Position uint16
	Current  statement.AwooParserStatement
	Context  compiler_context.AwooCompilerContext
	Settings AwooCompilerSettings
}

type AwooCompilerSettings struct {
	Path string
}

func SetupCompiler(settings AwooCompilerSettings, context parser_context.AwooParserContext) AwooCompiler {
	compiler := AwooCompiler{
		Context: compiler_context.AwooCompilerContext{
			Parser: context,
			Scopes: compiler_context.SetupCompilerScopeContainer(),
			Functions: compiler_context.AwooCompilerFunctionContainer{
				Entries: map[string]compiler_context.AwooCompilerFunction{},
			},
			MappingsStatement: map[uint16]compiler_context.AwooCompileStatement{
				statement.ParserStatementTypeDefinitionVariable: statement_compile.CompileStatementDefinition,
				statement.ParserStatementTypeAssignment:         statement_compile.CompileStatementAssignment,
				statement.ParserStatementTypeDefinitionType: func(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
					return []byte{}, nil
				},
				statement.ParserStatementTypeIf:     statement_compile.CompileStatementIf,
				statement.ParserStatementTypeGroup:  statement_compile.CompileStatementGroup,
				statement.ParserStatementTypeFunc:   statement_compile.CompileStatementFunc,
				statement.ParserStatementTypeReturn: statement_compile.CompileStatementReturn,
			},
			MappingsNodeValue: map[uint16]compiler_context.AwooCompileNodeValue{
				node.ParserNodeTypeIdentifier:  statement_compile.CompileNodeIdentifier,
				node.ParserNodeTypePrimitive:   statement_compile.CompileNodePrimitive,
				node.ParserNodeTypeExpression:  statement_compile.CompileNodeExpression,
				node.ParserNodeTypeNegative:    statement_compile.CompileNodeNegative,
				node.ParserNodeTypeReference:   statement_compile.CompileNodeReference,
				node.ParserNodeTypeDereference: statement_compile.CompileNodeDereference,
				node.ParserNodeTypeCall:        statement_compile.CompileNodeCall,
			},
			MappingsNodeExpression: map[uint16]compiler_context.AwooCompileNodeExpression{
				token.TokenOperatorAddition:       statement_compile.CompileNodeExpressionAdd,
				token.TokenOperatorSubstraction:   statement_compile.CompileNodeExpressionSubstract,
				token.TokenOperatorMultiplication: statement_compile.CompileNodeExpressionMultiply,
				token.TokenOperatorDivision:       statement_compile.CompileNodeExpressionDivide,
				token.TokenOperatorEqEq:           statement_compile.CompileNodeExpressionEqEq,
				token.TokenOperatorNotEq:          statement_compile.CompileNodeExpressionNotEq,
				token.TokenOperatorLT:             statement_compile.CompileNodeExpressionLT,
				token.TokenOperatorLTEQ:           statement_compile.CompileNodeExpressionLTEQ,
				token.TokenOperatorGT:             statement_compile.CompileNodeExpressionGT,
				token.TokenOperatorGTEQ:           statement_compile.CompileNodeExpressionGTEQ,
			},
		},
		Settings: settings,
	}
	return compiler
}

func LoadCompiler(compiler *AwooCompiler, contents parser.AwooParserResult) {
	compiler.Contents = contents
	compiler.Length = (uint16)(len(contents.Statements))
	compiler.Position = 0
	compiler.Current = compiler.Contents.Statements[compiler.Position]
}

func AdvanceCompilerFor(compiler *AwooCompiler, n int16) bool {
	compiler.Position = (uint16)((int16)(compiler.Position) + n)
	if compiler.Position >= compiler.Length {
		return false
	}
	compiler.Current = compiler.Contents.Statements[compiler.Position]
	return true
}

func AdvanceCompiler(compiler *AwooCompiler) bool {
	return AdvanceCompilerFor(compiler, 1)
}

func StepbackCompiler(compiler *AwooCompiler) bool {
	return AdvanceCompilerFor(compiler, -1)
}
