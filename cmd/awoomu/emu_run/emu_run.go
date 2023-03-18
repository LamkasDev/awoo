package emu_run

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	commonInstruction "github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func Load(path string) {
	/* program, _ := SelectProgram() */
	emulator := emu.SetupEmulator()
	rom.LoadROMFromPath(&emulator.Internal.ROM, path)
	Run(&emulator)
}

func Run(emulator *emu.AwooEmulator) {
	go func() {
		cycles := emulator.Config.CPU.Speed / 1000
		for emulator.Internal.Executing {
			for i := uint32(0); i < cycles && emulator.Internal.Executing; i++ {
				ProcessCycle(emulator)
				for _, id := range emulator.TickDrivers {
					emulator.Drivers[id] = emulator.Drivers[id].Tick(&emulator.Internal, emulator.Drivers[id])
				}
				emulator.Internal.CPU.TotalCycles++
				emulator.Internal.Executing = emulator.Internal.CPU.Counter < emulator.Internal.ROM.Length
			}
			time.Sleep(time.Millisecond)
		}
	}()
	for emulator.Internal.Running {
		// TODO: this will need a proper lock system, if a driver has both tick and tick long
		for _, id := range emulator.TickLongDrivers {
			emulator.Drivers[id] = emulator.Drivers[id].TickLong(&emulator.Internal, emulator.Drivers[id])
		}
		time.Sleep(time.Millisecond)
	}
	for _, driver := range emulator.Drivers {
		if driver.Clean != nil {
			_, err := driver.Clean(&emulator.Internal, driver)
			if err != nil {
				panic(err)
			}
		}
	}
	logger.Log("cycles: %s; read: %s; write: %s\n",
		gchalk.Red(fmt.Sprint(emulator.Internal.CPU.TotalCycles)),
		gchalk.Green(fmt.Sprint(emulator.Internal.Memory.TotalRead)),
		gchalk.Blue(fmt.Sprint(emulator.Internal.Memory.TotalWrite)),
	)
}

func ProcessCycle(emulator *emu.AwooEmulator) {
	raw := emulator.Internal.ROM.Data[emulator.Internal.CPU.Counter : emulator.Internal.CPU.Counter+4]
	rawIns := arch.AwooInstruction(binary.BigEndian.Uint32(raw))
	ins, err := instruction.Decode(emulator.Table, rawIns)
	if err != nil {
		panic(err)
	}
	instruction.PrintInternalInstruction(&emulator.Internal, raw, ins)
	ins.Process.(func(*internal.AwooEmulatorInternal, commonInstruction.AwooInstruction))(&emulator.Internal, ins)
	fmt.Printf("\n")

	if emulator.Internal.CPU.Advance {
		emulator.Internal.CPU.Counter += 4
	}
	emulator.Internal.CPU.Advance = true
	emulator.Internal.CPU.Registers[cpu.AwooRegisterZero] = 0
}
