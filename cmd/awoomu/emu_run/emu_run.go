package emu_run

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
)

func Load(path string) {
	/* program, _ := SelectProgram() */
	emulator := emu.SetupEmulator()
	rom.LoadROMFromPath(&emulator.Internal.ROM, path)
	Run(&emulator)
}

func Run(emulator *emu.AwooEmulator) {
	for emulator.Running {
		internal.TickInternal(&emulator.Internal)
		for i, driver := range emulator.Drivers {
			driver.Tick(&emulator.Internal, &driver)
			emulator.Drivers[i] = driver
		}
		emulator.Running = emulator.Internal.CPU.Counter < emulator.Internal.ROM.Length
	}
	n1 := int(memory.ReadMemory32(&emulator.Internal.CPU.Memory, 0))
	n2 := int(memory.ReadMemory32(&emulator.Internal.CPU.Memory, 4))
	logger.Log("%d %d\n", n1, n2)
}
