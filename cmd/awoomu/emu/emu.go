package emu

import (
	"encoding/binary"
	"fmt"
	"os/user"
	"path"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
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
	u, _ := user.Current()
	emulator := SetupEmulator()
	rom.LoadROMFromPath(&emulator.ROM, path.Join(u.HomeDir, "Documents", "awoo", "data", "output.bin"))
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
			"c: %s; r: %s; code: %s (%s); src: %s & %s; dst: %s; im: %s\n",
			gchalk.Red(fmt.Sprintf("%#x", emulator.CPU.Counter)),
			gchalk.Cyan(fmt.Sprintf("%#x %#x %#x %#x", raw[0:1], raw[1:2], raw[2:3], raw[3:4])),
			gchalk.Green(fmt.Sprintf("%#x", ins.Instruction.Code)),
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
	n1 := int(memory.ReadMemory32(&emulator.CPU.Memory, 0))
	n2 := int(memory.ReadMemory32(&emulator.CPU.Memory, 4))
	fmt.Printf("%d %d\n", n1, n2)
}
