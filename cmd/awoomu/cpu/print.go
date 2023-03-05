package cpu

import (
	"fmt"

	"github.com/jwalton/gchalk"
)

func PrintDecodedInstruction(ins AwooDecodedInstruction) string {
	return fmt.Sprintf(
		"code: %s (%s); src: %s & %s; dst: %s; im: %s",
		gchalk.Green(fmt.Sprintf("%#x", ins.Instruction.Code)),
		gchalk.Blue(ins.Instruction.Name),
		gchalk.Yellow(AwooRegisterNames[ins.SourceOne]),
		gchalk.Yellow(AwooRegisterNames[ins.SourceTwo]),
		gchalk.Yellow(AwooRegisterNames[ins.Destination]),
		gchalk.Magenta(fmt.Sprintf("%d", ins.Immediate)),
	)
}
