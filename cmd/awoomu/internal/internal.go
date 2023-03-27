package internal

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
)

type AwooEmulatorInternal struct {
	Running   bool
	Executing bool
	CPU       cpu.AwooCPU
	Memory    memory.AwooMemory
}

func SetupInternal() AwooEmulatorInternal {
	return AwooEmulatorInternal{
		Running:   true,
		Executing: true,
		CPU:       cpu.SetupCPU(),
		Memory:    memory.SetupMemory(16777216),
	}
}
