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

func ReadCharacterDriverVga(internal *internal.AwooEmulatorInternal, data *AwooDriverDataVga, offset arch.AwooRegister) (*sdl.Surface, uint8, bool) {
	characterData := memory.ReadMemorySafe(&internal.CPU.Memory, arch.AwooRegister(AwooDriverVgaVector+offset), memory.ReadMemory16)
	internal.CPU.Memory.TotalRead -= 2
	if characterData == 0 {
		return nil, 0, false
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

	return text, bgColor, true
}

func HandleEventsDriverVga(internal *internal.AwooEmulatorInternal, data *AwooDriverDataVga) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			internal.Executing = false
			internal.Running = false
		case *sdl.WindowEvent:
			switch event.(*sdl.WindowEvent).Event {
			case sdl.WINDOWEVENT_RESIZED:
				var err error
				if data.Renderer.Surface, err = data.Renderer.Window.GetSurface(); err != nil {
					panic(err)
				}
			}
		}
	}
}

func TickLongDriverVga(internal *internal.AwooEmulatorInternal, driver driver.AwooDriver) driver.AwooDriver {
	data := driver.Data.(AwooDriverDataVga)
	if data.Ticks > 1000/AwooDriverVgaFps {
		if err := data.Renderer.Surface.FillRect(nil, 0); err != nil {
			panic(err)
		}
		fontX, fontY := 0, 0
		for y := 0; y < AwooDriverVgaFrameHeight; y++ {
			for x := 0; x < AwooDriverVgaFrameWidth; x++ {
				offset := arch.AwooRegister((y * AwooDriverVgaFrameWidth * AwooDriverVgaCharacterSize) + (x * AwooDriverVgaCharacterSize))
				text, bgColor, ok := ReadCharacterDriverVga(internal, &data, offset)
				if !ok {
					break
				}
				if err := data.Renderer.Surface.FillRect(&sdl.Rect{X: int32(fontX), Y: int32(fontY), W: text.W, H: text.H}, AwooDriverVGAColors[bgColor].Uint32()); err != nil {
					panic(err)
				}
				if err := text.Blit(nil, data.Renderer.Surface, &sdl.Rect{X: int32(fontX), Y: int32(fontY), W: 0, H: 0}); err != nil {
					panic(err)
				}
				fontX += int(text.W)
			}
			fontX = 0
			fontY += data.Renderer.Font.Height()
		}
		if err := data.Renderer.Window.UpdateSurface(); err != nil {
			panic(err)
		}
		data.Ticks = 0
	}
	HandleEventsDriverVga(internal, &data)
	data.Ticks++
	driver.Data = data

	return driver
}

func CleanDriverVga(_ *internal.AwooEmulatorInternal, driver driver.AwooDriver) (driver.AwooDriver, error) {
	data := driver.Data.(AwooDriverDataVga)
	return driver, CleanRenderer(&data.Renderer)
}
