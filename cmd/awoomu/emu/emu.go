package emu

import (
	"syscall"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/driver"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/driver/vga"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
)

var (
	winmmDLL            = syscall.NewLazyDLL("winmm.dll")
	procTimeBeginPeriod = winmmDLL.NewProc("timeBeginPeriod")
)

type AwooEmulator struct {
	Internal        internal.AwooEmulatorInternal
	Drivers         map[uint16]driver.AwooDriver
	TickDrivers     []uint16
	TickLongDrivers []uint16
}

func SetupEmulator() AwooEmulator {
	procTimeBeginPeriod.Call(uintptr(1))
	emulator := AwooEmulator{
		Internal: internal.AwooEmulatorInternal{
			Running:   true,
			Executing: true,
			CPU:       cpu.SetupCPU(),
		},
		Drivers:         map[uint16]driver.AwooDriver{},
		TickDrivers:     []uint16{},
		TickLongDrivers: []uint16{},
	}
	AddEmulatorDriver(&emulator, vga.SetupDriverVga(&emulator.Internal))

	return emulator
}

func AddEmulatorDriver(emulator *AwooEmulator, driver driver.AwooDriver) {
	emulator.Drivers[driver.Id] = driver
	if driver.Tick != nil {
		emulator.TickDrivers = append(emulator.TickDrivers, driver.Id)
	}
	if driver.TickLong != nil {
		emulator.TickLongDrivers = append(emulator.TickLongDrivers, driver.Id)
	}
}
