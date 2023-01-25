package encoder

import (
	"encoding/binary"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
)

type AwooEncodedInstruction struct {
	Instruction instruction.AwooInstruction
	SourceOne   uint8
	SourceTwo   uint8
	Destination uint8
	Immediate   uint32
}

func Encode(ins AwooEncodedInstruction, data []byte) ([]byte, error) {
	// TODO: finish this silly goose
	raw := arch.AwooRegister(ins.Instruction.Code)
	switch ins.Instruction.Format {
	case instruction.AwooInstructionFormatR:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Instruction.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceTwo), 20, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Instruction.Argument)>>3, 25, 7)
	case instruction.AwooInstructionFormatI:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Instruction.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate), 20, 12)
	case instruction.AwooInstructionFormatS:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate), 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Instruction.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceTwo), 20, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>5, 25, 7)
	}

	return binary.BigEndian.AppendUint32(data, uint32(raw)), nil
}
