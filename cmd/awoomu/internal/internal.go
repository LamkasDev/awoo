package internal

import (
	"encoding/binary"
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

type AwooEmulatorInternal struct {
	Running   bool
	Executing bool
	CPU       cpu.AwooCPU
	ROM       rom.AwooRom
}

func TickInternal(internal *AwooEmulatorInternal) {
	raw := internal.ROM.Data[internal.CPU.Counter : internal.CPU.Counter+4]
	rawIns := arch.AwooInstruction(binary.BigEndian.Uint32(raw))
	ins, err := cpu.Decode(internal.CPU.Table, rawIns)
	if err != nil {
		panic(err)
	}
	logger.Log(
		"c: %s; r: %s; code: %s (%s); src: %s & %s; dst: %s; im: %s\n",
		gchalk.Red(fmt.Sprintf("%#x", internal.CPU.Counter)),
		gchalk.Cyan(fmt.Sprintf("%#x %#x %#x %#x", raw[0:1], raw[1:2], raw[2:3], raw[3:4])),
		gchalk.Green(fmt.Sprintf("%#x", ins.Instruction.Code)),
		gchalk.Blue(ins.Instruction.Name),
		gchalk.Yellow(cpu.AwooRegisterNames[ins.SourceOne]),
		gchalk.Yellow(cpu.AwooRegisterNames[ins.SourceTwo]),
		gchalk.Yellow(cpu.AwooRegisterNames[ins.Destination]),
		gchalk.Magenta(fmt.Sprintf("%d", ins.Immediate)),
	)
	ins.Process(&internal.CPU, ins)
	if internal.CPU.Advance {
		internal.CPU.Counter += 4
	}
	internal.CPU.Advance = true
	internal.CPU.Registers[cpu.AwooRegisterZero] = 0
}
