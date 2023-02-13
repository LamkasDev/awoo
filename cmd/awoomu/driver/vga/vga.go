package vga

import (
	"unsafe"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/driver"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/veandco/go-sdl2/sdl"
)

const AwooDriverIdVGA = 0x000
const AwooDriverVGAVector = arch.AwooRegister(0xb8000)

type AwooDriverDataVGA struct {
	Ticks    uint32
	Frame    [AwooDriverVgaFrameSize]uint32
	Renderer AwooDriverVGARenderer
}

func SetupDriverVGA(internal *internal.AwooEmulatorInternal) driver.AwooDriver {
	internal.CPU.Memory.Lockable = append(internal.CPU.Memory.Lockable, memory.AwooMemoryLockable{
		Start: AwooDriverVGAVector,
		End:   AwooDriverVGAVector + 32768,
	})

	return driver.AwooDriver{
		Id:       AwooDriverIdVGA,
		Name:     "VGA",
		TickLong: TickLongDriverVGA,
		Clean:    CleanDriverVGA,
		Data: AwooDriverDataVGA{
			Renderer: SetupRenderer(),
		},
	}
}

func ReadDriverVGA(internal *internal.AwooEmulatorInternal, driver *driver.AwooDriver, pos arch.AwooRegister) int16 {
	return memory.ReadMemorySafe(&internal.CPU.Memory, arch.AwooRegister(AwooDriverVGAVector+pos), memory.ReadMemory16)
}

func TickLongDriverVGA(internal *internal.AwooEmulatorInternal, driver *driver.AwooDriver) {
	data := driver.Data.(AwooDriverDataVGA)
	if data.Ticks > 1000/AwooDriverVgaFps {
		// TODO: construct frame
		for x := 0; x < AwooDriverVgaFrameWidth; x++ {
			for y := 0; y < AwooDriverVgaFrameHeight; y++ {
				sdlAddress := (y * AwooDriverVgaFrameWidth) + x
				memAddress := arch.AwooRegister((y * AwooDriverVgaFrameWidth * 2) + (x * 2))
				characterData := ReadDriverVGA(internal, driver, memAddress)
				if characterData != 0 {
					data.Frame[sdlAddress] = 0xff0000ff
				}
			}
		}

		data.Renderer.Texture.Update(nil, unsafe.Pointer(&data.Frame), data.Renderer.Pitch)
		data.Renderer.Renderer.Clear()
		data.Renderer.Renderer.Copy(data.Renderer.Texture, nil, nil)
		data.Renderer.Renderer.Present()
		data.Ticks = 0
	}
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			internal.Executing = false
			internal.Running = false
		}
	}

	data.Ticks++
	driver.Data = data
}

func CleanDriverVGA(_ *internal.AwooEmulatorInternal, driver *driver.AwooDriver) {
	data := driver.Data.(AwooDriverDataVGA)
	CleanRenderer(&data.Renderer)
}
