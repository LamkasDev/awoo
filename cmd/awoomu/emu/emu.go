package emu

import (
	"encoding/binary"
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/arch"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
	"github.com/jwalton/gchalk"
)

type AwooEmulator struct {
	Running bool
	CPU     cpu.AwooCPU
	ROM     rom.AwooRom
}

func SetupEmulator() AwooEmulator {
	return AwooEmulator{
		Running: true,
		CPU:     cpu.SetupCPU(),
	}
}

func Load() {
	println(fmt.Sprintf("hi from %s :3", gchalk.Red(arch.AwooPlatform)))

	/* program, _ := SelectProgram() */
	program := "E:\\code\\go\\awoo-emu\\data\\mul.bin"
	emulator := SetupEmulator()
	rom.LoadROMFromPath(&emulator.ROM, program)
	Run(&emulator)

	println(fmt.Sprintf("bay! :33"))
}

func Run(emulator *AwooEmulator) {
	for emulator.Running {
		raw := emulator.ROM.Data[emulator.CPU.Counter : emulator.CPU.Counter+4]
		rawIns := arch.AwooInstruction(binary.BigEndian.Uint32(raw))
		ins, err := cpu.Decode(emulator.CPU.Table, rawIns)
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"c: %s; r: %s %s %s %s; code: %s (%s); src: %s & %s; dst: %s; im: %s\n",
			gchalk.Red(fmt.Sprintf("0x%x", emulator.CPU.Counter)),
			gchalk.Cyan(fmt.Sprintf("0x%x", raw[0:1])),
			gchalk.Cyan(fmt.Sprintf("0x%x", raw[1:2])),
			gchalk.Cyan(fmt.Sprintf("0x%x", raw[2:3])),
			gchalk.Cyan(fmt.Sprintf("0x%x", raw[3:4])),
			gchalk.Green(fmt.Sprintf("0x%x", ins.Instruction.Code)),
			gchalk.Blue(ins.Instruction.Name),
			gchalk.Yellow(cpu.AwooRegisterNames[ins.SourceOne]),
			gchalk.Yellow(cpu.AwooRegisterNames[ins.SourceTwo]),
			gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%d", ins.Immediate)),
		)
		ins.Process(&emulator.CPU, ins)

		if emulator.CPU.Advance {
			emulator.CPU.Counter += 4
		}
		emulator.CPU.Advance = true
		emulator.Running = emulator.CPU.Counter < emulator.ROM.Length
	}
}
