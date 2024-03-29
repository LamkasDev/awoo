//go:build awoo64

package arch

const AwooPlatform = "awoo64"

type AwooRegisterIndex uint8
type AwooRegister int64
type AwooRegisterU uint64

type AwooInstruction uint32

type AwooWordDouble int64
type AwooWordDoubleU uint64
type AwooWord int64
type AwooWordU uint64
type AwooWordHalf int32
type AwooWordHalfU uint32
type AwooWordByte int8
type AwooWordByteU uint8

const AwooImmediateSmallMax = 2047
