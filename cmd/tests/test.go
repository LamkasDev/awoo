package tests

import (
	"testing"

	"github.com/LamkasDev/awoo-emu/cmd/emu"
	"github.com/LamkasDev/awoo-emu/cmd/rom"
)

type InstructionTestCheck func(emulator *emu.AwooEmulator) bool
type InstructionTest struct {
	ROM       []byte
	Registers func(emulator *emu.AwooEmulator)
	Results   map[string]InstructionTestCheck
}

func RunInstructionTest(suite *testing.T, test InstructionTest) {
	// Load
	emulator := emu.SetupEmulator()
	rom.LoadROM(&emulator.ROM, test.ROM)
	if test.Registers != nil {
		test.Registers(&emulator)
	}

	// Run
	emu.Run(&emulator)

	// Check
	for name, check := range test.Results {
		if !check(&emulator) {
			suite.Fatalf(name)
		}
	}
}
