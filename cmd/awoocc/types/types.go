package types

import "github.com/LamkasDev/awoo-emu/cmd/common/types"

type AwooType struct {
	Key           string
	Id            types.AwooTypeId
	PrimitiveType types.AwooTypePrimitiveId
	Length        uint8
	Size          uint32
	Flags         uint64
}
