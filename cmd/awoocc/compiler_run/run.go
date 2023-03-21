package compiler_run

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_compile"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
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
	elf := elf.NewAwooElf(elf.AwooElfTypeObject)
	for ok := true; ok; ok = compiler.AdvanceCompiler(ccompiler) {
		parser.PrintStatement(&ccompiler.Settings.Parser, &ccompiler.Context.Parser, &ccompiler.Current)
		err := statement_compile.CompileStatement(ccompiler, &elf, ccompiler.Current)
		if err != nil {
			panic(err)
		}
		// compiler.PrintNewCompile(ccompiler, &ccompiler.Current, data)
	}
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(elf); err != nil {
		panic(err)
	}
	writer.Write(data.Bytes())
	writer.Flush()
	file.Close()
}
