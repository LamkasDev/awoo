package instruction

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
)

const AwooInstructionCodeLength = 7

type AwooInstructionRange struct {
	Start  uint8
	Length uint8
}
type AwooInstructionRangeExtended struct {
	Offset uint8
	Ranges []AwooInstructionRange
}

type AwooInstruction struct {
	Code     uint8
	Argument uint16
	Format   uint8
	Name     string
	Advance  bool
}

func ProcessExtendedRange(raw arch.AwooInstruction, rangeExtended AwooInstructionRangeExtended, extendSign bool) arch.AwooRegister {
	if len(rangeExtended.Ranges) == 0 {
		return 0
	}
	value := arch.AwooRegister(0)
	offset := rangeExtended.Offset
	for _, currentRange := range rangeExtended.Ranges {
		currentRangeValue := util.SelectRangeRegister(raw, currentRange.Start, currentRange.Length)
		value = util.InsertRangeRegister(value, currentRangeValue, offset, currentRange.Length)
		offset += currentRange.Length
	}
	if extendSign {
		value = util.FillSignBits(value, offset-1)
	}

	return value
}
