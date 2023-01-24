package rom

import (
	"os"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
)

type AwooRom struct {
	Data   []byte
	Length arch.AwooRegister
}

func LoadROMFromPath(rom *AwooRom, path string) {
	data, _ := os.ReadFile(path)
	LoadROM(rom, data)
}

func LoadROM(rom *AwooRom, data []byte) {
	rom.Data = data
	rom.Length = arch.AwooRegister(len(rom.Data))
}
