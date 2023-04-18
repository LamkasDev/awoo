package linker_run

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/elf"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/linker"
	"github.com/LamkasDev/awoo-emu/cmd/awoold/symbol"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	commonElf "github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func RunLinker(clinker *linker.AwooLinker) {
	if err := elf.PrependProgramData(clinker, &clinker.Contents, make([]byte, 12)); err != nil {
		panic(err)
	}
	if err := symbol.ResolveSymbols(clinker, &clinker.Contents); err != nil {
		panic(err)
	}
	commonElf.AlignSections(&clinker.Contents)

	stackOffset := clinker.Contents.SectionList.Sections[clinker.Contents.SectionList.DataIndex].Address + arch.AwooRegister(len(clinker.Contents.SectionList.Sections[clinker.Contents.SectionList.DataIndex].Contents))
	stackAdjustmentInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionADDI,
		SourceOne:   cpu.AwooRegisterSavedZero,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   stackOffset,
	}
	if err := encoder.EncodeAt(&clinker.Contents, 0, stackAdjustmentInstruction); err != nil {
		panic(err)
	}

	jumpGlobalInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
		Immediate:   clinker.Contents.SymbolTable[cc.AwooCompilerGlobalFunctionName].Start,
	}
	if err := encoder.EncodeAt(&clinker.Contents, 4, jumpGlobalInstruction); err != nil {
		panic(err)
	}

	jumpMainInstruction := instruction.AwooInstruction{
		Definition:  instructions.AwooInstructionJALR,
		Destination: cpu.AwooRegisterReturnAddress,
		Immediate:   clinker.Contents.SymbolTable[cc.AwooCompilerMainFunctionName].Start,
	}
	if err := encoder.EncodeAt(&clinker.Contents, 8, jumpMainInstruction); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", clinker.Contents)

	if err := os.MkdirAll(filepath.Dir(clinker.Settings.Path), 0644); err != nil {
		panic(err)
	}
	outputFile, err := os.Create(clinker.Settings.Path)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(outputFile)
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(clinker.Contents); err != nil {
		panic(err)
	}
	writer.Write(data.Bytes())
	writer.Flush()
	outputFile.Close()
}
