package compiler

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
)

type AwooCompiler struct {
	Contents AwooCompilerContents
	Context  AwooCompilerContext
	Settings AwooCompilerSettings
}

type AwooCompilerContents struct {
	Data     parser.AwooParserResult
	Length   uint32
	Position uint32
}

type AwooCompilerContext struct {
	Parser parser.AwooParserContext
	Scopes scope.AwooScopeContainer
}

type AwooCompilerSettings struct {
	Path     string
	Parser   parser.AwooParserSettings
	Mappings AwooCompilerMappings
}

func NewCompiler(settings AwooCompilerSettings, context parser.AwooParserContext, data parser.AwooParserResult) AwooCompiler {
	return AwooCompiler{
		Contents: AwooCompilerContents{
			Data:   data,
			Length: (uint32)(len(data.Statements)),
		},
		Context: AwooCompilerContext{
			Parser: context,
			Scopes: scope.NewScopeContainer(),
		},
		Settings: settings,
	}
}

func GetCompilerStatement(ccompiler *AwooCompiler) statement.AwooParserStatement {
	return ccompiler.Contents.Data.Statements[ccompiler.Contents.Position]
}

func AdvanceCompilerFor(compiler *AwooCompiler, n int32) *statement.AwooParserStatement {
	compiler.Contents.Position = (uint32)((int32)(compiler.Contents.Position) + n)
	if compiler.Contents.Position >= compiler.Contents.Length {
		return nil
	}
	statement := GetCompilerStatement(compiler)
	return &statement
}

func AdvanceCompiler(compiler *AwooCompiler) *statement.AwooParserStatement {
	return AdvanceCompilerFor(compiler, 1)
}

func StepbackCompiler(compiler *AwooCompiler) *statement.AwooParserStatement {
	return AdvanceCompilerFor(compiler, -1)
}
