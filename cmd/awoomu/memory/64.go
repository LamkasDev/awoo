//go:build awoo64

package memory

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

func WriteMemoryWordDouble(mem *AwooMemory, n arch.AwooRegister, data arch.AwooWordDouble) {
	WriteMemory128(mem, n, (int128)(data))
}

func WriteMemoryWord(mem *AwooMemory, n arch.AwooRegister, data arch.AwooWord) {
	WriteMemory64(mem, n, (int64)(data))
}

func WriteMemoryWordHalf(mem *AwooMemory, n arch.AwooRegister, data arch.AwooWordHalf) {
	WriteMemory32(mem, n, (int32)(data))
}

func ReadMemoryWordDouble(mem *AwooMemory, n arch.AwooRegister) arch.AwooWordDouble {
	return arch.AwooWordDouble(ReadMemory128(mem, n))
}

func ReadMemoryWord(mem *AwooMemory, n arch.AwooRegister) arch.AwooWord {
	return arch.AwooWord(ReadMemory64(mem, n))
}

func ReadMemoryWordHalf(mem *AwooMemory, n arch.AwooRegister) arch.AwooWordHalf {
	return arch.AwooWordHalf(ReadMemory32(mem, n))
}
