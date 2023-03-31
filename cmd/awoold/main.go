package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"os"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker_run"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/flags"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/paths"
)

func main() {
	u, _ := user.Current()
	defaultInput := path.Join(u.HomeDir, "Documents", "awoo", "data", "obj", "input.awoobj")
	defaultOutput := path.Join(u.HomeDir, "Documents", "awoo", "data", "bin", "input.awooxe")

	var input string
	var output string
	var quiet bool
	flag.StringVar(&input, "i", defaultInput, "path to input .awoobj file")
	flag.StringVar(&output, "o", defaultOutput, "path to output .awooxe file")
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	logger.AwooLoggerEnabled = !quiet
	input, output = paths.ResolvePaths(input, ".awoobj", output, ".awooxe")
	flags.ResolveColor()

	inputFile, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	var elf elf.AwooElf
	if err := gob.NewDecoder(bytes.NewBuffer(inputFile)).Decode(&elf); err != nil {
		panic(err)
	}

	linkerSettings := linker.AwooLinkerSettings{
		Path: output,
		Mappings: linker.AwooLinkerMappings{
			InstructionTable: instructions.SetupInstructionTable(),
		},
	}
	clinker := linker.SetupLinker(linkerSettings)
	linker.LoadLinker(&clinker, elf)
	linker_run.RunLinker(&clinker)
}
