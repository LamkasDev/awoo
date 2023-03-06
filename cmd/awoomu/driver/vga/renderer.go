package vga

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type AwooDriverVGARenderer struct {
	Window    *sdl.Window
	Surface   *sdl.Surface
	Renderer  *sdl.Renderer
	Texture   *sdl.Texture
	Font      *ttf.Font
	Fontsheet map[uint16]*sdl.Surface
	Pitch     int
}

func SetupRenderer() (AwooDriverVGARenderer, error) {
	renderer := AwooDriverVGARenderer{
		Fontsheet: map[uint16]*sdl.Surface{},
	}
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return renderer, fmt.Errorf("failed to initialize SDL: %w", err)
	}

	w, h := int32(800), int32(600)
	renderer.Window, err = sdl.CreateWindow(fmt.Sprintf("Emulator (%s)", arch.AwooPlatform), sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		return renderer, fmt.Errorf("failed to create window: %w", err)
	}
	renderer.Surface, err = renderer.Window.GetSurface()
	if err != nil {
		return renderer, fmt.Errorf("failed to get surface: %w", err)
	}
	renderer.Renderer, err = sdl.CreateRenderer(renderer.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return renderer, fmt.Errorf("failed to create renderer: %w", err)
	}
	renderer.Texture, err = renderer.Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, AwooDriverVgaFrameWidth, AwooDriverVgaFrameHeight)
	if err != nil {
		return renderer, fmt.Errorf("failed to create surface texture: %w", err)
	}
	renderer.Pitch = int(AwooDriverVgaFrameWidth * 4)

	err = ttf.Init()
	if err != nil {
		return renderer, fmt.Errorf("failed to initialize TTF: %w", err)
	}
	renderer.Font, err = ttf.OpenFont("../../resources/fonts/Roboto-Regular.ttf", 32)
	if err != nil {
		return renderer, fmt.Errorf("failed to load font: %w", err)
	}

	return renderer, nil
}

func CleanRenderer(renderer *AwooDriverVGARenderer) error {
	err := renderer.Texture.Destroy()
	if err != nil {
		return fmt.Errorf("failed to destroy surface texture: %w", err)
	}
	err = renderer.Renderer.Destroy()
	if err != nil {
		return fmt.Errorf("failed to destroy renderer: %w", err)
	}
	err = renderer.Window.Destroy()
	if err != nil {
		return fmt.Errorf("failed to destroy window: %w", err)
	}
	renderer.Font.Close()
	for _, c := range renderer.Fontsheet {
		c.Free()
	}

	sdl.Quit()
	return nil
}
