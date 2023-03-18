package compiler_run

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_compile"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func RunCompiler(ccompiler *compiler.AwooCompiler) {
	logger.Log(gchalk.Yellow("\n> Compiler\n"))

	err := os.MkdirAll(filepath.Dir(ccompiler.Settings.Path), 0644)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(ccompiler.Settings.Path)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	for ok := true; ok; ok = compiler.AdvanceCompiler(ccompiler) {
		parser.PrintStatement(&ccompiler.Settings.Parser, &ccompiler.Context.Parser, &ccompiler.Current)
		data, err := statement_compile.CompileStatement(ccompiler, ccompiler.Current, []byte{})
		if err != nil {
			panic(err)
		}
		compiler.PrintNewCompile(ccompiler, &ccompiler.Current, data)
		_, err = writer.Write(data)
		if err != nil {
			panic(err)
		}
		ccompiler.Context.CurrentAddress += uint16(len(data))
	}
	writer.Flush()
	err = CompileProgramHeader(ccompiler, file, writer)
	if err != nil {
		panic(err)
	}
	file.Close()
}
