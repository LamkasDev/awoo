package util

import (
	"github.com/LamkasDev/awoo-emu/cmd/arch"
)

func SelectRangeRegister(raw uint32, start uint8, length uint8) arch.AwooRegister {
	return arch.AwooRegister((raw >> start) & (1<<length - 1))
}

func InsertRangeRegister(accumulator uint32, current uint32, start uint8, length uint8) arch.AwooRegister {
	// Zero out the bits where we're planning to insert new bits
	mask := ^(^uint32(0) << length) << start
	accumulator &= ^mask

	// Insert new bits
	current = current << start

	return arch.AwooRegister(accumulator | current)
}
