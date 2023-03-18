package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

var AwooRegisterNames = map[arch.AwooRegisterIndex]string{
	0x00: "zero",
	0x01: "ra",
	0x02: "sp",
	0x03: "gp",
	0x04: "tp",
	0x05: "t0",
	0x06: "t1",
	0x07: "t2",
	0x08: "s0",
	0x09: "s1",
	0x0a: "a0",
	0x0b: "a1",
	0x0c: "a2",
	0x0d: "a3",
	0x0e: "a4",
	0x0f: "a5",
	0x10: "a6",
	0x11: "a7",
	0x12: "s2",
	0x13: "s3",
	0x14: "s4",
	0x15: "s5",
	0x16: "s6",
	0x17: "s7",
	0x18: "s8",
	0x19: "s9",
	0x1a: "s10",
	0x1b: "s11",
	0x1c: "t3",
	0x1d: "t4",
	0x1e: "t5",
	0x1f: "t6",
}

const AwooRegisterZero = arch.AwooRegisterIndex(0x00)
const AwooRegisterReturnAddress = arch.AwooRegisterIndex(0x01)
const AwooRegisterStackPointer = arch.AwooRegisterIndex(0x02)
const AwooRegisterGlobalPointer = arch.AwooRegisterIndex(0x03)
const AwooRegisterThreadPointer = arch.AwooRegisterIndex(0x04)
const AwooRegisterTemporaryZero = arch.AwooRegisterIndex(0x05)
const AwooRegisterTemporaryOne = arch.AwooRegisterIndex(0x06)
const AwooRegisterTemporaryTwo = arch.AwooRegisterIndex(0x07)
const AwooRegisterSavedZero = arch.AwooRegisterIndex(0x08)
const AwooRegisterSavedOne = arch.AwooRegisterIndex(0x09)
const AwooRegisterFunctionZero = arch.AwooRegisterIndex(0x0a)
const AwooRegisterFunctionOne = arch.AwooRegisterIndex(0x0b)
const AwooRegisterFunctionTwo = arch.AwooRegisterIndex(0x0c)
const AwooRegisterFunctionThree = arch.AwooRegisterIndex(0x0d)
const AwooRegisterFunctionFour = arch.AwooRegisterIndex(0x0e)
const AwooRegisterFunctionFive = arch.AwooRegisterIndex(0x0f)
const AwooRegisterFunctionSix = arch.AwooRegisterIndex(0x10)
const AwooRegisterFunctionSeven = arch.AwooRegisterIndex(0x11)
const AwooRegisterSavedTwo = arch.AwooRegisterIndex(0x12)
const AwooRegisterSavedThree = arch.AwooRegisterIndex(0x13)
const AwooRegisterSavedFour = arch.AwooRegisterIndex(0x14)
const AwooRegisterSavedFive = arch.AwooRegisterIndex(0x15)
const AwooRegisterSavedSix = arch.AwooRegisterIndex(0x16)
const AwooRegisterSavedSeven = arch.AwooRegisterIndex(0x17)
const AwooRegisterSavedEight = arch.AwooRegisterIndex(0x18)
const AwooRegisterSavedNine = arch.AwooRegisterIndex(0x19)
const AwooRegisterSavedTen = arch.AwooRegisterIndex(0x1a)
const AwooRegisterSavedEleven = arch.AwooRegisterIndex(0x1b)
const AwooRegisterTemporaryThree = arch.AwooRegisterIndex(0x1c)
const AwooRegisterTemporaryFour = arch.AwooRegisterIndex(0x1d)
const AwooRegisterTemporaryFive = arch.AwooRegisterIndex(0x1e)
const AwooRegisterTemporarySix = arch.AwooRegisterIndex(0x1f)

func GetNextTemporaryRegister(r arch.AwooRegisterIndex) arch.AwooRegisterIndex {
	if r == AwooRegisterTemporaryTwo {
		return AwooRegisterTemporaryThree
	}
	return r + 1
}
