package main

import (
	"flag"
	"fmt"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/flags"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/paths"
	"github.com/jwalton/gchalk"
)

func main() {
	logger.Log(fmt.Sprintf("hi from %s :3", gchalk.Red(arch.AwooPlatform)))

	u, _ := user.Current()
	defaultInput := path.Join(u.HomeDir, "Documents", "awoo", "data", "obj", "input.awoobj")
	var input string
	var quiet bool
	flag.StringVar(&input, "i", defaultInput, "path to input .awoobj file")
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	logger.AwooLoggerEnabled = !quiet
	flags.ResolveColor()
	input = paths.ResolvePath(input, ".awoobj")

	emu.Load(input)

	logger.Log("bay! :33")
}
