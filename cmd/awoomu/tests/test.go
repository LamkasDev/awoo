package tests

import (
	"testing"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu_run"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
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
	rom.LoadROM(&emulator.Internal.ROM, test.ROM)
	if test.Registers != nil {
		test.Registers(&emulator)
	}

	// Run
	emu_run.Run(&emulator)

	// Check
	for name, check := range test.Results {
		if !check(&emulator) {
			suite.Fatalf(name)
		}
	}
}
