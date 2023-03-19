package cpu

type AwooRegisterId uint8

var AwooRegisterNames = map[AwooRegisterId]string{
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

const AwooRegisterZero = AwooRegisterId(0x00)
const AwooRegisterReturnAddress = AwooRegisterId(0x01)
const AwooRegisterStackPointer = AwooRegisterId(0x02)
const AwooRegisterGlobalPointer = AwooRegisterId(0x03)
const AwooRegisterThreadPointer = AwooRegisterId(0x04)
const AwooRegisterTemporaryZero = AwooRegisterId(0x05)
const AwooRegisterTemporaryOne = AwooRegisterId(0x06)
const AwooRegisterTemporaryTwo = AwooRegisterId(0x07)
const AwooRegisterSavedZero = AwooRegisterId(0x08)
const AwooRegisterSavedOne = AwooRegisterId(0x09)
const AwooRegisterFunctionZero = AwooRegisterId(0x0a)
const AwooRegisterFunctionOne = AwooRegisterId(0x0b)
const AwooRegisterFunctionTwo = AwooRegisterId(0x0c)
const AwooRegisterFunctionThree = AwooRegisterId(0x0d)
const AwooRegisterFunctionFour = AwooRegisterId(0x0e)
const AwooRegisterFunctionFive = AwooRegisterId(0x0f)
const AwooRegisterFunctionSix = AwooRegisterId(0x10)
const AwooRegisterFunctionSeven = AwooRegisterId(0x11)
const AwooRegisterSavedTwo = AwooRegisterId(0x12)
const AwooRegisterSavedThree = AwooRegisterId(0x13)
const AwooRegisterSavedFour = AwooRegisterId(0x14)
const AwooRegisterSavedFive = AwooRegisterId(0x15)
const AwooRegisterSavedSix = AwooRegisterId(0x16)
const AwooRegisterSavedSeven = AwooRegisterId(0x17)
const AwooRegisterSavedEight = AwooRegisterId(0x18)
const AwooRegisterSavedNine = AwooRegisterId(0x19)
const AwooRegisterSavedTen = AwooRegisterId(0x1a)
const AwooRegisterSavedEleven = AwooRegisterId(0x1b)
const AwooRegisterTemporaryThree = AwooRegisterId(0x1c)
const AwooRegisterTemporaryFour = AwooRegisterId(0x1d)
const AwooRegisterTemporaryFive = AwooRegisterId(0x1e)
const AwooRegisterTemporarySix = AwooRegisterId(0x1f)

func GetNextTemporaryRegister(r AwooRegisterId) AwooRegisterId {
	if r == AwooRegisterTemporaryTwo {
		return AwooRegisterTemporaryThree
	}
	return r + 1
}
