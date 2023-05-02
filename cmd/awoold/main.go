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
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
)

func main() {
	util.RegisterGobTypes()
	u, _ := user.Current()

	passedInputs := []string{}
	passedOutputs := []string{}
	var quiet bool
	flag.Func("i", "path to input .awoobj file", func(s string) error {
		passedInputs = append(passedInputs, s)
		return nil
	})
	flag.Func("o", "path to output .awooxe file", func(s string) error {
		passedOutputs = append(passedOutputs, s)
		return nil
	})
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	if len(passedInputs) == 0 {
		passedInputs = append(passedInputs, path.Join(u.HomeDir, "Documents", "awoo", "data", "obj"))
	}
	if len(passedOutputs) == 0 {
		passedOutputs = append(passedOutputs, path.Join(u.HomeDir, "Documents", "awoo", "data", "bin", "input.awooxe"))
	}
	flags.ResolveColor()
	logger.AwooLoggerEnabled = !quiet
	inputs, outputs := paths.ResolveAllPaths(passedInputs, ".awoobj", passedOutputs[0], ".awooxe")

	elfs := []elf.AwooElf{}
	for _, input := range inputs {
		inputFile, err := os.ReadFile(input.Absolute)
		if err != nil {
			panic(err)
		}

		var elf elf.AwooElf
		decoder := gob.NewDecoder(bytes.NewBuffer(inputFile))
		if err := decoder.Decode(&elf); err != nil {
			panic(err)
		}
		elfs = append(elfs, elf)
	}

	linkerSettings := linker.AwooLinkerSettings{
		Path: outputs[0].Absolute,
		Mappings: linker.AwooLinkerMappings{
			InstructionTable: instructions.SetupInstructionTable(),
		},
	}
	clinker := linker.SetupLinker(linkerSettings)
	linker.LoadLinker(&clinker, elfs)
	linker_run.RunLinker(&clinker)
}
