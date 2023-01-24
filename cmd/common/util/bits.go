package util

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

func SelectRangeRegister(raw arch.AwooInstruction, start uint8, length uint8) arch.AwooRegister {
	return arch.AwooRegister((raw >> start) & (1<<length - 1))
}

func InsertRangeRegister(accumulator arch.AwooRegister, current arch.AwooRegister, start uint8, length uint8) arch.AwooRegister {
	// Zero out the range where we're planning to insert current range
	mask := ^(^arch.AwooRegister(0) << length) << start
	accumulator &= ^mask

	// Align the current range to where we want to insert it
	current <<= start

	return accumulator | current
}

func FillSignBits(raw arch.AwooRegister, start uint8) arch.AwooRegister {
	// Create regular bitmask and shift it to the sign bit's position
	mask := arch.AwooRegister(1 << (32 - start - 1))
	mask = mask << start

	return raw | mask
}
