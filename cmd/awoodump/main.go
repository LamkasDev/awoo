package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/flags"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/paths"
)

func main() {
	u, _ := user.Current()
	defaultInput := path.Join(u.HomeDir, "Documents", "awoo", "data", "obj", "input.awoobj")

	var input string
	var quiet bool
	flag.StringVar(&input, "i", defaultInput, "path to input .awoobj file")
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	logger.AwooLoggerEnabled = !quiet
	input = paths.ResolvePath(input, ".awoobj")
	flags.ResolveColor()

	inputFile, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	var elf elf.AwooElf
	if err := gob.NewDecoder(bytes.NewBuffer(inputFile)).Decode(&elf); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", elf)
}
