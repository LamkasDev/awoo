package vga

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/driver"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/memory"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/veandco/go-sdl2/sdl"
)

type AwooDriverDataVga struct {
	Ticks    uint32
	Renderer AwooDriverVGARenderer
}

func SetupDriverVga(internal *internal.AwooEmulatorInternal) driver.AwooDriver {
	internal.CPU.Memory.Lockable = append(internal.CPU.Memory.Lockable, memory.AwooMemoryLockable{
		Start: AwooDriverVgaVector,
		End:   AwooDriverVgaVector + AwooDriverVgaSize,
	})
	renderer, err := SetupRenderer()
	if err != nil {
		panic(err)
	}

	return driver.AwooDriver{
		Id:       AwooDriverIdVga,
		Name:     "VGA",
		TickLong: TickLongDriverVga,
		Clean:    CleanDriverVga,
		Data: AwooDriverDataVga{
			Renderer: renderer,
		},
	}
}

func ReadDriverVga(internal *internal.AwooEmulatorInternal, driver *driver.AwooDriver, offset arch.AwooRegister) int16 {
	return memory.ReadMemorySafe(&internal.CPU.Memory, arch.AwooRegister(AwooDriverVgaVector+offset), memory.ReadMemory16)
}

func TickLongDriverVga(internal *internal.AwooEmulatorInternal, driver *driver.AwooDriver) {
	data := driver.Data.(AwooDriverDataVga)
	if data.Ticks > 1000/AwooDriverVgaFps {
		data.Renderer.Surface.FillRect(nil, 0)
		fontX, fontY := 0, 0
		for y := 0; y < AwooDriverVgaFrameHeight; y++ {
			for x := 0; x < AwooDriverVgaFrameWidth; x++ {
				characterOffset := arch.AwooRegister((y * AwooDriverVgaFrameWidth * AwooDriverVgaCharacterSize) + (x * AwooDriverVgaCharacterSize))
				characterData := ReadDriverVga(internal, driver, characterOffset)
				if characterData == 0 {
					continue
				}

				fgColor := uint8(characterData & 0b1111)
				bgColor := uint8((characterData >> 4) & 0b0111)
				asciiCode := uint8(characterData >> 8)
				sheetCode := uint16(asciiCode) + uint16(fgColor)
				text, ok := data.Renderer.Fontsheet[sheetCode]
				if !ok {
					text, _ = data.Renderer.Font.RenderUTF8Blended(string(rune(asciiCode)), AwooDriverVGAColors[fgColor])
					data.Renderer.Fontsheet[sheetCode] = text
				}
				data.Renderer.Surface.FillRect(&sdl.Rect{X: int32(fontX), Y: int32(fontY), W: text.W, H: text.H}, AwooDriverVGAColors[bgColor].Uint32())
				text.Blit(nil, data.Renderer.Surface, &sdl.Rect{X: int32(fontX), Y: int32(fontY), W: 0, H: 0})
				fontX += int(text.W)
			}
			fontY += data.Renderer.Font.Height()
		}
		data.Renderer.Window.UpdateSurface()
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

func CleanDriverVga(_ *internal.AwooEmulatorInternal, driver *driver.AwooDriver) error {
	data := driver.Data.(AwooDriverDataVga)
	return CleanRenderer(&data.Renderer)
}
