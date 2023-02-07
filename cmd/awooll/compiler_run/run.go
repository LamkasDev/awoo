package compiler_run

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement_compile"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func RunCompiler(ccompiler *compiler.AwooCompiler) compiler.AwooCompilerResult {
	result := compiler.AwooCompilerResult{}
	logger.Log(gchalk.Yellow("\n> Compiler\n"))
	logger.Log("Input: %s\n", gchalk.Magenta(fmt.Sprintf("%v", ccompiler.Contents.Statements)))

	err := os.MkdirAll(filepath.Dir(ccompiler.Settings.Path), 0644)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(ccompiler.Settings.Path)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	// TODO: jump beyond function definitions at start.
	for ok := true; ok; ok = compiler.AdvanceCompiler(ccompiler) {
		statement.PrintStatement(&ccompiler.Context.Parser.Lexer, &ccompiler.Current)
		data, err := statement_compile.CompileStatement(&ccompiler.Context, ccompiler.Current, []byte{})
		if err != nil {
			result.Error = err
			break
		}
		compiler.PrintNewCompile(&ccompiler.Context.Parser.Lexer, &ccompiler.Current, data)
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

	logger.Log(gchalk.Yellow("\n> Memory map\n"))
	for _, scope := range ccompiler.Context.Scopes.Entries {
		logger.Log("┏━ %s (%s)\n", scope.Name, gchalk.Green(fmt.Sprintf("%#x", scope.Id)))
		for _, entry := range scope.Memory.Entries {
			t := ccompiler.Context.Parser.Lexer.Types.All[entry.Type]
			logger.Log("┣━ %s %s  %s (%s)\n",
				gchalk.Green(fmt.Sprintf("%#x - %#x", entry.Start, entry.Start+uint16(t.Size)-1)),
				gchalk.Gray("➔"),
				entry.Name,
				gchalk.Cyan(t.Key),
			)
		}
		logger.Log("┗━━► %s entries\n", gchalk.Blue(fmt.Sprint(len(scope.Memory.Entries))))
	}

	return result
}
