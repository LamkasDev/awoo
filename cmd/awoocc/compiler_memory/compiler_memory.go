package compiler_memory

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
)

type AwooCompilerMemory struct {
	Entries  map[string]AwooCompilerMemoryEntry
	Position arch.AwooRegister
}

type AwooCompilerMemoryEntry struct {
	Symbol elf.AwooElfSymbolTableEntry
	Global bool
}
