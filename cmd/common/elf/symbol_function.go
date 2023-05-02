package elf

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooElfSymbolTableEntryFunctionDetails struct {
	ReturnType *types.AwooTypeId
	Arguments  []AwooElfSymbolTableEntry
}

func IsSymbolFunction(symbol AwooElfSymbolTableEntry) bool {
	_, ok := symbol.Details.(AwooElfSymbolTableEntryFunctionDetails)
	return ok
}

func GetSymbolFunctionArgumentsSize(symbol AwooElfSymbolTableEntry) arch.AwooRegister {
	details := symbol.Details.(AwooElfSymbolTableEntryFunctionDetails)
	offset := arch.AwooRegister(0)
	for _, argument := range details.Arguments {
		offset += argument.Size
	}
	return offset
}
