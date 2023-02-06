package util

import (
	"math"

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

func FillSignBits(imm arch.AwooRegister, start uint8) arch.AwooRegister {
	// Check if immediate's sign bit is set.
	if imm>>start&1 == 1 {
		// Create bitmask with (32 - start) 1s and shift it to the sign bit's position.
		mask := arch.AwooRegister(math.Pow(2, float64(32-start)) - 1)
		mask <<= start
		return imm | mask
	}

	return imm
}
