package compiler_run

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	awooElf "github.com/LamkasDev/awoo-emu/cmd/awoocc/elf"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement_compile"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
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
	elf := commonElf.NewAwooElf(commonElf.AwooElfTypeObject)
	for ok := true; ok; ok = compiler.AdvanceCompiler(ccompiler) {
		start := len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents)
		parser.PrintStatement(&ccompiler.Settings.Parser, &ccompiler.Context.Parser, &ccompiler.Current)
		if err := statement_compile.CompileStatement(ccompiler, &elf, ccompiler.Current); err != nil {
			panic(err)
		}
		end := len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents)
		compiler.PrintNewCompile(ccompiler, &ccompiler.Current, elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents[start:end])
	}
	awooElf.PopulateSymbolTable(ccompiler, &elf)
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(elf); err != nil {
		panic(err)
	}
	writer.Write(data.Bytes())
	writer.Flush()
	file.Close()
}
