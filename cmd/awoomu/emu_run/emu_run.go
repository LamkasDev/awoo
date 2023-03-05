package emu_run

import (
	"time"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
)

func Load(path string) {
	/* program, _ := SelectProgram() */
	emulator := emu.SetupEmulator()
	rom.LoadROMFromPath(&emulator.Internal.ROM, path)
	Run(&emulator)
}

func Run(emulator *emu.AwooEmulator) {
	go func() {
		cycles := cpu.AwooCPURate / 1000
		for emulator.Internal.Executing {
			for i := uint32(0); i < cycles; i++ {
				internal.TickInternal(&emulator.Internal)
				for i, driver := range emulator.Drivers {
					if driver.Tick != nil {
						driver.Tick(&emulator.Internal, &driver)
						emulator.Drivers[i] = driver
					}
				}
				emulator.Internal.Executing = emulator.Internal.CPU.Counter < emulator.Internal.ROM.Length
			}
			time.Sleep(time.Millisecond)
		}
	}()
	for emulator.Internal.Running {
		// TODO: this will need a proper lock system, if a driver has both tick and tick long
		for i, driver := range emulator.Drivers {
			if driver.TickLong != nil {
				driver.TickLong(&emulator.Internal, &driver)
				emulator.Drivers[i] = driver
			}
		}
		time.Sleep(time.Millisecond)
	}
	for _, driver := range emulator.Drivers {
		if driver.Clean != nil {
			driver.Clean(&emulator.Internal, &driver)
		}
	}
}
