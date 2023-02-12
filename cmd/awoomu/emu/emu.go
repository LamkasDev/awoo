package emu

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/driver"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/driver/vga"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
)

type AwooEmulator struct {
	Running  bool
	Internal internal.AwooEmulatorInternal
	Drivers  map[uint16]driver.AwooDriver
}

func SetupEmulator() AwooEmulator {
	emulator := AwooEmulator{
		Running: true,
		Internal: internal.AwooEmulatorInternal{
			CPU: cpu.SetupCPU(),
		},
		Drivers: map[uint16]driver.AwooDriver{},
	}
	AddEmulatorDriver(&emulator, vga.SetupDriverVGA())

	return emulator
}

func AddEmulatorDriver(emulator *AwooEmulator, driver driver.AwooDriver) {
	emulator.Drivers[driver.Id] = driver
}
