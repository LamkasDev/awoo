package cpu

import (
	"fmt"

	"github.com/jwalton/gchalk"
)

func PrintDecodedInstruction(ins AwooDecodedInstruction) string {
	return fmt.Sprintf(
		"code: %-36s; src: %s; dst: %-15s; im: %s;",
		fmt.Sprintf("%s (%s)", gchalk.Green(fmt.Sprintf("%#4x", ins.Instruction.Code)), gchalk.Blue(ins.Instruction.Name)),
		fmt.Sprintf("%-14s & %-15s", gchalk.Yellow(AwooRegisterNames[ins.SourceOne]), gchalk.Yellow(AwooRegisterNames[ins.SourceTwo])),
		gchalk.Yellow(AwooRegisterNames[ins.Destination]),
		gchalk.Magenta(fmt.Sprintf("%-8d", ins.Immediate)),
	)
}
