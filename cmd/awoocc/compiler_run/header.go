package compiler_run

import (
	"bufio"
	"os"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
)

func CompileProgramHeader(ccompiler *compiler.AwooCompiler, file *os.File, writer *bufio.Writer) error {
	// TODO: this is very stupid way to skip to the main function
	mainFunc, ok := compiler_context.GetCompilerFunction(&ccompiler.Context, "main")
	if !ok {
		return awerrors.ErrorFailedToCompileProgramHeader
	}
	firstFuncStart := ccompiler.Context.Functions.Entries[ccompiler.Context.Functions.Start].Start - compiler_context.GetProgramHeaderSize()
	file.Seek(int64(firstFuncStart), 0)
	d, err := encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionADDI,
		Destination: cpu.AwooRegisterSavedZero,
		Immediate:   uint32(ccompiler.Context.Scopes.Global.Position),
	}, []byte{})
	if err != nil {
		return err
	}
	// TODO: this is not correct
	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionJAL,
		Destination: cpu.AwooRegisterReturnAddress,
		Immediate:   uint32(mainFunc.Start - firstFuncStart - 12),
	}, d)
	if err != nil {
		return err
	}
	_, err = writer.Write(d)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
