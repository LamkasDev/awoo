package vga

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/driver"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

const AwooDriverIdVGA = 0x000
const AwooDriverVGAVector = uint32(0xb8000)

type AwooDriverDataVGA struct {
	Position uint32
	Renderer AwooDriverVGARenderer
}

func SetupDriverVGA() driver.AwooDriver {
	return driver.AwooDriver{
		Id:   AwooDriverIdVGA,
		Name: "VGA",
		Tick: TickDriverVGA,
		Data: AwooDriverDataVGA{},
	}
}

func ReadDriverVGA(internal *internal.AwooEmulatorInternal, driver *driver.AwooDriver) int32 {
	return memory.ReadMemory32(&internal.CPU.Memory, arch.AwooRegister(AwooDriverVGAVector+driver.Data.(AwooDriverDataVGA).Position))
}

func TickDriverVGA(internal *internal.AwooEmulatorInternal, driver *driver.AwooDriver) {
	for v := ReadDriverVGA(internal, driver); v != 0; v = ReadDriverVGA(internal, driver) {
		print(string(v))
		data := driver.Data.(AwooDriverDataVGA)
		data.Position += 4
		driver.Data = data
	}
}
