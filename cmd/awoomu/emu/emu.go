package emu

import (
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/config"
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
	Config          config.AwooConfig
	Internal        internal.AwooEmulatorInternal
	Drivers         map[uint16]driver.AwooDriver
	TickDrivers     []uint16
	TickLongDrivers []uint16
}

func SetupEmulator() AwooEmulator {
	procTimeBeginPeriod.Call(uintptr(1))
	emulator := AwooEmulator{
		Config: config.NewAwooConfig(),
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

	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	if err = os.MkdirAll(filepath.Join(u.HomeDir, "Documents", "awoo", "config"), 0755); err != nil {
		panic(err)
	}
	if emulator.Config, err = config.ReadConfig(emulator.Config); err != nil {
		panic(err)
	}

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
