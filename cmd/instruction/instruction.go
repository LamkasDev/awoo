package instruction

import (
	"github.com/LamkasDev/awoo-emu/cmd/arch"
	"github.com/LamkasDev/awoo-emu/cmd/util"
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
	Code   uint8
	Format uint8
	Name   string
}

func ProcessExtendedRange(raw uint32, rangeExtended AwooInstructionRangeExtended, extendSign bool) arch.AwooRegister {
	value := arch.AwooRegister(0)
	offset := rangeExtended.Offset
	for _, currentRange := range rangeExtended.Ranges {
		currentRangeValue := util.SelectRangeRegister(raw, currentRange.Start, currentRange.Length)
		value = util.InsertRangeRegister((uint32)(value), (uint32)(currentRangeValue), offset, currentRange.Length)
		offset += currentRange.Length
	}
	if extendSign && offset < 32 {
		util.InsertRangeRegister((uint32)(value), 1, offset, 32-offset)
	}

	return value
}
