package main

import (
	"flag"
	"fmt"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu_run"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/flags"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/paths"
	"github.com/jwalton/gchalk"
)

func main() {
	logger.Log(fmt.Sprintf("hi from %s :3\n", gchalk.Red(arch.AwooPlatform)))
	u, _ := user.Current()
	defaultInput := path.Join(u.HomeDir, "Documents", "awoo", "data", "bin", "input.awooxe")

	var input string
	var quiet bool
	flag.StringVar(&input, "i", defaultInput, "path to input .awooxe file")
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	logger.AwooLoggerEnabled = !quiet
	flags.ResolveColor()
	input = paths.ResolvePath(input, ".awooxe")

	emulator := emu.SetupEmulator()
	emu_run.Load(&emulator, input)
	emu_run.Run(&emulator)

	logger.Log("bay! :33\n")
}
