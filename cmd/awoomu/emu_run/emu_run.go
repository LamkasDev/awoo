package emu_run

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/decoder"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	commonInstruction "github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func Load(emulator *emu.AwooEmulator, path string) {
	osFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var osElf elf.AwooElf
	decoder := gob.NewDecoder(bytes.NewBuffer(osFile))
	if err := decoder.Decode(&osElf); err != nil {
		panic(err)
	}

	memoryLength := arch.AwooRegister(0)

	copy(emulator.Internal.Memory.Data[memoryLength:], osElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents)
	programLength := arch.AwooRegister(len(osElf.SectionList.Sections[elf.AwooElfSectionProgram].Contents))
	memoryLength += programLength
	emulator.Internal.Memory.ProgramEnd = programLength

	copy(emulator.Internal.Memory.Data[memoryLength:], osElf.SectionList.Sections[elf.AwooElfSectionData].Contents)
	dataLength := arch.AwooRegister(len(osElf.SectionList.Sections[elf.AwooElfSectionData].Contents))
	// memoryLength += dataLength

	emulator.Internal.CPU.Counter = osElf.Counter

	logger.Log("loaded %s; program: %s; data: %s\n",
		gchalk.Red(path),
		gchalk.Green(fmt.Sprint(programLength)),
		gchalk.Blue(fmt.Sprint(dataLength)),
	)
}

func Run(emulator *emu.AwooEmulator) {
	go func() {
		cycles := emulator.Config.CPU.Speed / 1000
		for emulator.Internal.Executing {
			for i := uint32(0); i < cycles && emulator.Internal.Executing; i++ {
				if arch.AwooDebug {
					emulator.Internal.CPU.Snapshot = emulator.Internal.CPU.Registers
				}
				ProcessCycle(emulator)
				for _, id := range emulator.TickDrivers {
					emulator.Drivers[id] = emulator.Drivers[id].Tick(&emulator.Internal, emulator.Drivers[id])
				}
				emulator.Internal.CPU.TotalCycles++
				if memory.ReadMemory32(&emulator.Internal.Memory, cpu.AwooCrsMstatus)&cpu.AwooCrsMstatusHalt == cpu.AwooCrsMstatusHalt {
					emulator.Internal.Executing = false
					logger.Log(fmt.Sprintf("received %s, stopping execution o/\n", gchalk.Red("mstatus halt")))
				}
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
	raw := arch.AwooInstruction(memory.ReadMemory32(&emulator.Internal.Memory, emulator.Internal.CPU.Counter))
	ins, err := decoder.Decode(emulator.Table, raw)
	if err != nil {
		panic(err)
	}
	instruction.PrintInternalInstruction(&emulator.Internal, ins)
	ins.Process.(func(*internal.AwooEmulatorInternal, commonInstruction.AwooInstruction))(&emulator.Internal, ins)
	fmt.Printf("\n")

	if emulator.Internal.CPU.Advance {
		emulator.Internal.CPU.Counter += 4
	}
	emulator.Internal.CPU.Advance = true
	emulator.Internal.CPU.Registers[cpu.AwooRegisterZero] = 0
}
