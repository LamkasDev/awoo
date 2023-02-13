package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

const AwooCPURate = uint32(1000)

// TODO: move memory into internals
type AwooCPU struct {
	Registers [31]arch.AwooRegister
	Counter   arch.AwooRegister
	Advance   bool
	Memory    memory.AwooMemory
	Table     instruction.AwooInstructionTable
}

var AwooRegisterNames = map[arch.AwooRegister]string{
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

const AwooRegisterZero = 0x00
const AwooRegisterReturnAddress = 0x01
const AwooRegisterStackPointer = 0x02
const AwooRegisterGlobalPointer = 0x03
const AwooRegisterThreadPointer = 0x04
const AwooRegisterTemporaryZero = 0x05
const AwooRegisterTemporaryOne = 0x06
const AwooRegisterTemporaryTwo = 0x07
const AwooRegisterSavedZero = 0x08
const AwooRegisterSavedOne = 0x09
const AwooRegisterFunctionZero = 0x0a
const AwooRegisterFunctionOne = 0x0b
const AwooRegisterFunctionTwo = 0x0c
const AwooRegisterFunctionThree = 0x0d
const AwooRegisterFunctionFour = 0x0e
const AwooRegisterFunctionFive = 0x0f
const AwooRegisterFunctionSix = 0x10
const AwooRegisterFunctionSeven = 0x11
const AwooRegisterSavedTwo = 0x12
const AwooRegisterSavedThree = 0x13
const AwooRegisterSavedFour = 0x14
const AwooRegisterSavedFive = 0x15
const AwooRegisterSavedSix = 0x16
const AwooRegisterSavedSeven = 0x17
const AwooRegisterSavedEight = 0x18
const AwooRegisterSavedNine = 0x19
const AwooRegisterSavedTen = 0x1a
const AwooRegisterSavedEleven = 0x1b
const AwooRegisterTemporaryThree = 0x1c
const AwooRegisterTemporaryFour = 0x1d
const AwooRegisterTemporaryFive = 0x1e
const AwooRegisterTemporarySix = 0x1f

func GetNextTemporaryRegister(r uint8) uint8 {
	if r == AwooRegisterTemporaryTwo {
		return AwooRegisterTemporaryThree
	}
	return r + 1
}

func SetupCPU() AwooCPU {
	return AwooCPU{
		Advance: true,
		Memory:  memory.SetupMemory(16777216),
		Table:   SetupDecoderInstructionTable(),
	}
}
