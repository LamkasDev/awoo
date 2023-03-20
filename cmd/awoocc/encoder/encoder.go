package encoder

import (
	"encoding/binary"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
)

type AwooEncodedInstruction struct {
	Instruction instruction.AwooInstructionDefinition
	SourceOne   cpu.AwooRegisterId
	SourceTwo   cpu.AwooRegisterId
	Destination cpu.AwooRegisterId
	Immediate   uint32
}

func EncodeAt(elf *elf.AwooElf, offset uint32, ins AwooEncodedInstruction) error {
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
	case instruction.AwooInstructionFormatB:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>11, 7, 1)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>1, 8, 4)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Instruction.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceTwo), 20, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>5, 25, 6)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>12, 31, 1)
	case instruction.AwooInstructionFormatU:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>12, 12, 20)
	case instruction.AwooInstructionFormatJ:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>12, 12, 8)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>11, 20, 1)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>1, 21, 10)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Immediate)>>20, 31, 1)
	}
	binary.BigEndian.PutUint32(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents[offset:], uint32(raw))

	return nil
}

func Encode(elf *elf.AwooElf, ins AwooEncodedInstruction) error {
	offset := uint32(len(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents))
	binary.BigEndian.AppendUint32(elf.SectionList.Sections[elf.SectionList.ProgramIndex].Contents, uint32(0))
	return EncodeAt(elf, offset, ins)
}
