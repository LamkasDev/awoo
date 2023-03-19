//go:build awoo32

package arch

const AwooPlatform = "awoo32"

type AwooRegister int32
type AwooRegisterU uint32

type AwooInstruction uint32

type AwooWordDouble int64
type AwooWordDoubleU uint64
type AwooWord int32
type AwooWordU uint32
type AwooWordHalf int16
type AwooWordHalfU uint16
type AwooWordByte int8
type AwooWordByteU uint8

const AwooImmediateSmallMax = 2047
