package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
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
	Path     string
	Parser   parser.AwooParserSettings
	Mappings AwooCompilerMappings
}

func SetupCompiler(settings AwooCompilerSettings, context parser_context.AwooParserContext) AwooCompiler {
	compiler := AwooCompiler{
		Context: compiler_context.AwooCompilerContext{
			Parser: context,
			Scopes: compiler_context.SetupCompilerScopeContainer(),
			Functions: compiler_context.AwooCompilerFunctionContainer{
				Entries: map[string]compiler_context.AwooCompilerFunction{},
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
