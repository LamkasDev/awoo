package tests

import (
	"testing"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
)

// ADD a0, a1, a2
func TestProcessADD(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x00, 0xc5, 0x85, 0x33},
		Registers: func(emulator *emu.AwooEmulator) {
			emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionOne] = 1
			emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionTwo] = 2
		},
		Results: map[string]InstructionTestCheck{
			"Expected 3 in a0": func(emulator *emu.AwooEmulator) bool {
				return emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionZero] == 3
			},
		},
	}
	RunInstructionTest(t, test)
}

// SUB a0, a1, a2
func TestProcessSUB(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x40, 0xc5, 0x85, 0x33},
		Registers: func(emulator *emu.AwooEmulator) {
			emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionOne] = 1
			emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionTwo] = 2
		},
		Results: map[string]InstructionTestCheck{
			"Expected 1 in a0": func(emulator *emu.AwooEmulator) bool {
				return emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionZero] == 1
			},
		},
	}
	RunInstructionTest(t, test)
}

// ADDI a0, a1, 3
func TestProcessADDI(t *testing.T) {
	test := InstructionTest{
		ROM:       []byte{0x00, 0x15, 0x85, 0x13},
		Registers: func(emulator *emu.AwooEmulator) { emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionOne] = 2 },
		Results: map[string]InstructionTestCheck{
			"Expected 3 in a0": func(emulator *emu.AwooEmulator) bool {
				return emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionZero] == 3
			},
		},
	}
	RunInstructionTest(t, test)
}

// SLT a0, a1, a2
func TestProcessSLT(t *testing.T) {
	test := InstructionTest{
		ROM:       []byte{0x00, 0xc5, 0xa5, 0x33},
		Registers: func(emulator *emu.AwooEmulator) { emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionTwo] = 1 },
		Results: map[string]InstructionTestCheck{
			"Expected 1 in a0": func(emulator *emu.AwooEmulator) bool {
				return emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionZero] == 1
			},
		},
	}
	RunInstructionTest(t, test)
}

// SLTI a0, a1, 1
func TestProcessSLTI(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x00, 0x15, 0xa5, 0x13},
		Results: map[string]InstructionTestCheck{
			"Expected 1 in a0": func(emulator *emu.AwooEmulator) bool {
				return emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionZero] == 1
			},
		},
	}
	RunInstructionTest(t, test)
}

// LUI a0, 1
func TestProcessLUI(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x00, 0x00, 0x15, 0x37},
		Results: map[string]InstructionTestCheck{
			"Expected 4096 in a0": func(emulator *emu.AwooEmulator) bool {
				return emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionZero] == 4096
			},
		},
	}
	RunInstructionTest(t, test)
}

// AUIPC a0, 1
func TestProcessAUIPC(t *testing.T) {
	test := InstructionTest{
		ROM: []byte{0x00, 0x00, 0x15, 0x17},
		Results: map[string]InstructionTestCheck{
			"Expected 4096 in a0": func(emulator *emu.AwooEmulator) bool {
				return emulator.Internal.CPU.Registers[cpu.AwooRegisterFunctionZero] == 4096
			},
		},
	}
	RunInstructionTest(t, test)
}
