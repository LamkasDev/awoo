package encoder

import (
	"encoding/binary"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
)

func EncodeAt(celf *elf.AwooElf, offset arch.AwooRegister, ins instruction.AwooInstruction) error {
	raw := arch.AwooRegister(ins.Definition.Code)
	switch ins.Definition.Format {
	case instruction.AwooInstructionFormatR:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Definition.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceTwo), 20, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Definition.Argument)>>3, 25, 7)
	case instruction.AwooInstructionFormatI:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Definition.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, ins.Immediate, 20, 12)
	case instruction.AwooInstructionFormatS:
		raw = util.InsertRangeRegister(raw, ins.Immediate, 7, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Definition.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceTwo), 20, 5)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>5, 25, 7)
	case instruction.AwooInstructionFormatB:
		raw = util.InsertRangeRegister(raw, ins.Immediate>>11, 7, 1)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>1, 8, 4)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Definition.Argument), 12, 3)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceOne), 15, 5)
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.SourceTwo), 20, 5)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>5, 25, 6)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>12, 31, 1)
	case instruction.AwooInstructionFormatU:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>12, 12, 20)
	case instruction.AwooInstructionFormatJ:
		raw = util.InsertRangeRegister(raw, arch.AwooRegister(ins.Destination), 7, 5)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>12, 12, 8)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>11, 20, 1)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>1, 21, 10)
		raw = util.InsertRangeRegister(raw, ins.Immediate>>20, 31, 1)
	}
	binary.BigEndian.PutUint32(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents[offset:], uint32(raw))

	return nil
}

func Encode(celf *elf.AwooElf, ins instruction.AwooInstruction) error {
	offset := arch.AwooRegister(len(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))
	celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents = binary.BigEndian.AppendUint32(celf.SectionList.Sections[elf.AwooElfSectionProgram].Contents, uint32(0))
	return EncodeAt(celf, offset, ins)
}
