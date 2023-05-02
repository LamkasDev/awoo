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
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
)

func main() {
	util.RegisterGobTypes()
	u, _ := user.Current()

	passedInputs := []string{}
	var quiet bool
	flag.Func("i", "path to input .awoobj file", func(s string) error {
		passedInputs = append(passedInputs, s)
		return nil
	})
	flag.BoolVar(&quiet, "q", false, "set to disable log")
	flag.Parse()
	if len(passedInputs) == 0 {
		passedInputs = append(passedInputs, path.Join(u.HomeDir, "Documents", "awoo", "data", "obj", "input.awoobj"))
	}
	flags.ResolveColor()
	logger.AwooLoggerEnabled = !quiet
	inputs := paths.CreatePathListCombined(passedInputs, ".awoobj")

	for _, input := range inputs {
		inputFile, err := os.ReadFile(input.Absolute)
		if err != nil {
			panic(err)
		}
		var elf elf.AwooElf
		if err := gob.NewDecoder(bytes.NewBuffer(inputFile)).Decode(&elf); err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", elf)
	}
}
