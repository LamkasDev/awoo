package internal

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintInternalInstruction(internal *AwooEmulatorInternal, raw []byte, ins cpu.AwooDecodedInstruction) {
	logger.Log(
		"c: %s; r: %s; %s",
		gchalk.Red(fmt.Sprintf("%#6x", internal.CPU.Counter)),
		gchalk.Cyan(fmt.Sprintf("%#x %#x %#x %#x", raw[0:1], raw[1:2], raw[2:3], raw[3:4])),
		cpu.PrintDecodedInstruction(ins),
	)
}
