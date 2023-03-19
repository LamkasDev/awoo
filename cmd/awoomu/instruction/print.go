package instruction

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintDecodedInstruction(ins instruction.AwooInstruction) string {
	return fmt.Sprintf(
		"code: %-36s; src: %s; dst: %-15s; im: %s; ",
		fmt.Sprintf("%s (%s)", gchalk.Green(fmt.Sprintf("%#4x", ins.Definition.Code)), gchalk.Blue(ins.Definition.Name)),
		fmt.Sprintf("%-14s & %-15s", gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.SourceOne)]), gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.SourceTwo)])),
		gchalk.Yellow(cpu.AwooRegisterNames[cpu.AwooRegisterId(ins.Destination)]),
		gchalk.Magenta(fmt.Sprintf("%-8d", ins.Immediate)),
	)
}

func PrintInternalInstruction(internal *internal.AwooEmulatorInternal, raw []byte, ins instruction.AwooInstruction) {
	logger.Log(
		"c: %s; r: %s; %s",
		gchalk.Red(fmt.Sprintf("%#6x", internal.CPU.Counter)),
		gchalk.Cyan(fmt.Sprintf("%#x %#x %#x %#x", raw[0:1], raw[1:2], raw[2:3], raw[3:4])),
		PrintDecodedInstruction(ins),
	)
}
