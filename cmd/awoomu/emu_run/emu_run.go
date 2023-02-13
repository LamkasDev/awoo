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
	go func() {
		for emulator.Running {
			internal.TickInternal(&emulator.Internal)
			for i, driver := range emulator.Drivers {
				driver.Tick(&emulator.Internal, &driver)
				emulator.Drivers[i] = driver
			}
			emulator.Running = emulator.Internal.CPU.Counter < emulator.Internal.ROM.Length
		}
	}
	rendererElapsed := uint16(0)
	rendererTiming := 1000 / vga
	for nes.Cycling {
		if rendererElapsed >= nes.Timings.Renderer {
			
		}
		time.Sleep(time.Millisecond)
		rendererElapsed++
	}
	n1 := int(memory.ReadMemory32(&emulator.Internal.CPU.Memory, 0))
	n2 := int(memory.ReadMemory32(&emulator.Internal.CPU.Memory, 4))
	logger.Log("%d %d\n", n1, n2)
}
