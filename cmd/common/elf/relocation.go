package elf

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooElfRelocationList []AwooElfRelocationListEntry

type AwooElfRelocationListEntry struct {
	Offset arch.AwooRegister
	Name   string
}
