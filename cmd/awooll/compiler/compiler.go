package compiler

import (
	"bufio"
	"fmt"
	"os"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/print"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

type AwooCompiler struct {
	Contents parser.AwooParserResult
	Length   uint16
	Position uint16
	Current  statement.AwooParserStatement
	Context  compiler_context.AwooCompilerContext
	Settings AwooCompilerSettings
}

type AwooCompilerSettings struct{}

func SetupCompiler(settings AwooCompilerSettings, context parser_context.AwooParserContext) AwooCompiler {
	compiler := AwooCompiler{
		Context: compiler_context.AwooCompilerContext{
			Parser: context,
			Memory: compiler_context.AwooCompilerContextMemory{
				Memory: make(map[string]compiler_context.AwooCompilerContextMemoryEntry),
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

func RunCompiler(compiler *AwooCompiler) AwooCompilerResult {
	result := AwooCompilerResult{}
	fmt.Println(gchalk.Yellow("\n> Compiler"))
	fmt.Printf("Input: %s\n", gchalk.Magenta(fmt.Sprintf("%v", compiler.Contents.Statements)))

	file, err := os.Create("E:\\code\\go\\awoo-emu\\data\\output.bin")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	for ok := true; ok; ok = AdvanceCompiler(compiler) {
		print.PrintStatement(&compiler.Context.Parser.Lexer, &compiler.Current)
		data, err := CompileStatement(&compiler.Context, compiler.Current)
		if err != nil {
			result.Error = err
			break
		}
		print.PrintNewCompile(&compiler.Context.Parser.Lexer, &compiler.Current, data)
		_, err = writer.Write(data)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	file.Close()
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println(gchalk.Yellow("\n> Memory map"))
	for name, entry := range compiler.Context.Memory.Memory {
		fmt.Printf("%s %s  %s\n",
			gchalk.Green(fmt.Sprintf("%#x - %#x", entry.Start, entry.Start+uint16(compiler.Context.Parser.Lexer.Types.All[entry.Type].Length))),
			gchalk.Gray("➔"),
			gchalk.Cyan(name),
		)
	}

	return result
}
