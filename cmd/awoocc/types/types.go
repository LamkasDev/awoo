package types

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooType struct {
	Key           string
	Id            types.AwooTypeId
	PrimitiveType types.AwooTypeId
	Length        uint8
	Size          arch.AwooRegister
	Flags         uint64
}
