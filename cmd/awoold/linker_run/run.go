package linker_run

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"

	"github.com/LamkasDev/awoo-emu/cmd/awoold/elf"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/header"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/symbol"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/davecgh/go-spew/spew"
)

func RunLinker(clinker *linker.AwooLinker) {
	celf := commonElf.NewAwooElf(filepath.Base(clinker.Settings.Path), commonElf.AwooElfTypeExecutable)
	header.ReserveHeader(clinker, &celf)
	if err := elf.ConcatObjects(clinker, &celf); err != nil {
		panic(err)
	}
	commonElf.AlignSections(&celf)
	if err := symbol.ResolveSymbols(clinker, &celf); err != nil {
		panic(err)
	}
	if err := header.PopulateHeader(clinker, &celf); err != nil {
		panic(err)
	}
	spew.Dump(celf)

	if err := os.MkdirAll(filepath.Dir(clinker.Settings.Path), 0644); err != nil {
		panic(err)
	}
	file, err := os.Create(clinker.Settings.Path)
	if err != nil {
		panic(err)
	}
	var data bytes.Buffer
	encoder := gob.NewEncoder(&data)
	if err := encoder.Encode(celf); err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	writer.Write(data.Bytes())
	writer.Flush()
	file.Close()
}
