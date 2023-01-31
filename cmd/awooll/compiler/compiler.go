package compiler

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
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
			Scopes: compiler_context.SetupCompilerScopeContainer(),
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

	u, _ := user.Current()
	file, err := os.Create(path.Join(u.HomeDir, "Documents", "awoo", "data", "output.bin"))
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	for ok := true; ok; ok = AdvanceCompiler(compiler) {
		statement.PrintStatement(&compiler.Context.Parser.Lexer, &compiler.Current)
		data, err := CompileStatement(&compiler.Context, compiler.Current)
		if err != nil {
			result.Error = err
			break
		}
		PrintNewCompile(&compiler.Context.Parser.Lexer, &compiler.Current, data)
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
	for _, scope := range compiler.Context.Scopes.Entries {
		fmt.Printf("┏━ %s (%s)\n", scope.Name, gchalk.Green(fmt.Sprintf("%#x", scope.ID)))
		for _, entry := range scope.Memory.Entries {
			t := compiler.Context.Parser.Lexer.Types.All[entry.Type]
			fmt.Printf("┣━ %s %s  %s (%s)\n",
				gchalk.Green(fmt.Sprintf("%#x - %#x", entry.Start, entry.Start+uint16(t.Size)-1)),
				gchalk.Gray("➔"),
				entry.Name,
				gchalk.Cyan(t.Key),
			)
		}
		fmt.Printf("┗━━► %s entries\n", gchalk.Blue(fmt.Sprint(len(scope.Memory.Entries))))
	}

	return result
}
