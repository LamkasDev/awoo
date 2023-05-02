package util

import (
	"encoding/gob"

	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func RegisterGobTypes() {
	gob.Register(types.AwooTypeId(0))
	gob.Register(elf.AwooElfSymbolTableEntryFunctionDetails{})
}
