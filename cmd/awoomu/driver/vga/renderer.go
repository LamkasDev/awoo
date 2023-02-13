package vga

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/veandco/go-sdl2/sdl"
)

const AwooDriverVgaFps = 90
const AwooDriverVgaFrameWidth = 80
const AwooDriverVgaFrameHeight = 25
const AwooDriverVgaFrameSize = AwooDriverVgaFrameWidth * AwooDriverVgaFrameHeight

type AwooDriverVGARenderer struct {
	Window   *sdl.Window
	Surface  *sdl.Surface
	Renderer *sdl.Renderer
	Texture  *sdl.Texture
	Pitch    int
}

func SetupRenderer() AwooDriverVGARenderer {
	renderer := AwooDriverVGARenderer{}
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	w, h := int32(800), int32(600)
	renderer.Window, err = sdl.CreateWindow(fmt.Sprintf("Emulator (%s)", arch.AwooPlatform), sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
	renderer.Surface, err = renderer.Window.GetSurface()
	renderer.Renderer, err = sdl.CreateRenderer(renderer.Window, -1, sdl.RENDERER_ACCELERATED)
	renderer.Texture, err = renderer.Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, AwooDriverVgaFrameWidth, AwooDriverVgaFrameHeight)
	renderer.Pitch = int(AwooDriverVgaFrameWidth * 4)

	return renderer
}

func CleanRenderer(renderer *AwooDriverVGARenderer) {
	renderer.Texture.Destroy()
	renderer.Renderer.Destroy()
	renderer.Window.Destroy()
	sdl.Quit()
}
