package compiler_run

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement_compile"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
	"golang.org/x/exp/maps"
)

func RunCompiler(ccompiler *compiler.AwooCompiler) {
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
	file.Seek(int64(compiler_context.GetProgramHeaderSize()), 0)
	for ok := true; ok; ok = compiler.AdvanceCompiler(ccompiler) {
		statement.PrintStatement(&ccompiler.Context.Parser.Lexer, &ccompiler.Current)
		data, err := statement_compile.CompileStatement(&ccompiler.Context, ccompiler.Current, []byte{})
		if err != nil {
			panic(err)
		}
		compiler.PrintNewCompile(&ccompiler.Context.Parser.Lexer, &ccompiler.Current, data)
		_, err = writer.Write(data)
		if err != nil {
			panic(err)
		}
		ccompiler.Context.Position += uint16(len(data))
	}
	writer.Flush()
	file.Seek(0, 0)
	d, err := CompileProgramHeader(ccompiler)
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(d)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	file.Close()

	logger.Log(gchalk.Yellow("\n> Memory map\n"))
	for _, f := range ccompiler.Context.Scopes.Functions {
		logger.Log("┏━ %s (%s)\n", f.Name, gchalk.Green(fmt.Sprintf("%#x", f.Id)))
		blocks := maps.Values(f.Blocks)
		sort.Slice(blocks, func(i, j int) bool {
			return blocks[i].Memory.Position < blocks[j].Memory.Position
		})
		for _, block := range blocks {
			entries := maps.Values(block.Memory.Entries)
			sort.Slice(entries, func(i, j int) bool {
				return entries[i].Start < entries[j].Start
			})
			for _, entry := range entries {
				t := ccompiler.Context.Parser.Lexer.Types.All[entry.Type]
				logger.Log("┣━ %s %s  %s (%s)\n",
					gchalk.Green(fmt.Sprintf("%#x - %#x", entry.Start, entry.Start+uint16(entry.Size)-1)),
					gchalk.Gray("➔"),
					entry.Name,
					gchalk.Cyan(t.Key),
				)
				logger.Log("┗━━► %s entries\n", gchalk.Blue(fmt.Sprint(len(block.Memory.Entries))))
			}
		}
		logger.Log("┗━━► %s blocks\n", gchalk.Blue(fmt.Sprint(len(f.Blocks))))
	}
	logger.Log("┗━━► %s functions\n", gchalk.Blue(fmt.Sprint(len(ccompiler.Context.Scopes.Functions))))
}
