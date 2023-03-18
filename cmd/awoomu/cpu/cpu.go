package cpu

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

type AwooCPU struct {
	Registers   [31]arch.AwooRegister
	Counter     arch.AwooRegister
	Advance     bool
	TotalCycles uint64
}

func SetupCPU() AwooCPU {
	return AwooCPU{
		Advance: true,
	}
}
