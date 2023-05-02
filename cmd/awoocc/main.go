package main

import (
	"flag"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_run"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/flags"
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
	flag.Func("i", "path to input .awoo file", func(s string) error {
		passedInputs = append(passedInputs, s)
		return nil
	})
	flag.Func("o", "path to output .awooobj file", func(s string) error {
		passedOutputs = append(passedOutputs, s)
		return nil
	})
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	if len(passedInputs) == 0 {
		passedInputs = append(passedInputs, path.Join(u.HomeDir, "Documents", "awoo", "data", "src"))
	}
	if len(passedOutputs) == 0 {
		passedOutputs = append(passedOutputs, path.Join(u.HomeDir, "Documents", "awoo", "data", "obj"))
	}
	flags.ResolveColor()
	logger.AwooLoggerEnabled = !quiet
	inputs, outputs := paths.ResolveAllPaths(passedInputs, ".awoo", passedOutputs[0], ".awoobj")

	context := map[string]elf.AwooElf{}
	for i, input := range inputs {
		compiler_run.RunCompilerFull(context, input, outputs[i], passedOutputs[0])
	}
}
