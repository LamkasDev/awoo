package encoder

import (
	"encoding/binary"

	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

type AwooEncodedInstruction struct {
	Instruction instruction.AwooInstruction
	SourceOne   uint8
	SourceTwo   uint8
	Destination uint8
	Immediate   uint32
}

func Encode(ins AwooEncodedInstruction, data []byte) error {
	raw := uint32(ins.Instruction.Code)
	switch ins.Instruction.Format {
	case instruction.AwooInstructionFormatI:
		raw |= (uint32(ins.Destination) << 7)
		raw |= (uint32(ins.Instruction.Argument) << 12)
		raw |= (uint32(ins.SourceOne) << 15)
		raw |= (uint32(ins.Immediate) << 20)
	case instruction.AwooInstructionFormatS:
		raw |= (uint32(ins.Immediate) << 7)
		raw |= (uint32(ins.Instruction.Argument) << 12)
		raw |= (uint32(ins.SourceOne) << 15)
		raw |= (uint32(ins.SourceTwo) << 20)
		raw |= ((uint32(ins.Immediate) >> 5) << 25)
	}

	binary.BigEndian.PutUint32(data, raw)
	return nil
}
