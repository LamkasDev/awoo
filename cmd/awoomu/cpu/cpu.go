package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

type AwooCPURegisters [31]arch.AwooRegister

type AwooCPU struct {
	Snapshot    AwooCPURegisters
	Registers   AwooCPURegisters
	Counter     arch.AwooRegister
	Advance     bool
	TotalCycles uint64
}

func SetupCPU() AwooCPU {
	return AwooCPU{
		Advance: true,
	}
}
