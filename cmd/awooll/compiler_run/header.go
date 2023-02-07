package compiler_run

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileProgramHeader(ccompiler *compiler.AwooCompiler) ([]byte, error) {
	f, ok := compiler_context.GetCompilerFunction(&ccompiler.Context, "main")
	if !ok {
		return []byte{}, awerrors.ErrorFailedToCompileProgramHeader
	}
	d, err := encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionJAL,
		Destination: cpu.AwooRegisterStackPointer,
		Immediate:   uint32(f.Start),
	}, []byte{})
	if err != nil {
		return d, fmt.Errorf("%w: %w", awerrors.ErrorFailedToCompileProgramHeader, err)
	}

	return d, nil
}
