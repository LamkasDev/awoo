package internal

import (
	"encoding/binary"
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
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
	PrintInternalInstruction(internal, raw, ins)
	ins.Process(&internal.CPU, ins)
	fmt.Printf("\n")

	if internal.CPU.Advance {
		internal.CPU.Counter += 4
	}
	internal.CPU.Advance = true
	internal.CPU.Registers[cpu.AwooRegisterZero] = 0
}
