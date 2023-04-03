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
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func RunCompiler(ccompiler *compiler.AwooCompiler) {
	logger.LogExtra(gchalk.Yellow("\n> Compiler\n"))

	err := os.MkdirAll(filepath.Dir(ccompiler.Settings.Path), 0644)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(ccompiler.Settings.Path)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	celf := commonElf.NewAwooElf(commonElf.AwooElfTypeObject)
	for ok := true; ok; ok = compiler.AdvanceCompiler(ccompiler) {
		start := len(celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents)
		parser.PrintStatement(&ccompiler.Settings.Parser, &ccompiler.Context.Parser, &ccompiler.Current)
		if err := statement_compile.CompileStatement(ccompiler, &celf, ccompiler.Current); err != nil {
			panic(err)
		}
		end := len(celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents)
		compiler.PrintNewCompile(ccompiler, &ccompiler.Current, celf.SectionList.Sections[celf.SectionList.ProgramIndex].Contents[start:end])
	}
	elf.AlignSections(&celf)
	if err := awooElf.AlignSymbols(ccompiler, &celf); err != nil {
		panic(err)
	}
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(celf); err != nil {
		panic(err)
	}
	writer.Write(data.Bytes())
	writer.Flush()
	file.Close()
}
