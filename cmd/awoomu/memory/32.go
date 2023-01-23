//go:build awoo32

package memory

import "github.com/LamkasDev/awoo-emu/cmd/awoomu/arch"

func WriteMemoryWordDouble(mem *AwooMemory, n arch.AwooRegister, data arch.AwooWordDouble) {
	WriteMemory64(mem, n, (uint64)(data))
}

func WriteMemoryWord(mem *AwooMemory, n arch.AwooRegister, data arch.AwooWord) {
	WriteMemory32(mem, n, (uint32)(data))
}

func WriteMemoryWordHalf(mem *AwooMemory, n arch.AwooRegister, data arch.AwooWordHalf) {
	WriteMemory16(mem, n, (uint16)(data))
}

func ReadMemoryWordDouble(mem *AwooMemory, n arch.AwooRegister) arch.AwooWordDouble {
	return arch.AwooWordDouble(ReadMemory64(mem, n))
}

func ReadMemoryWord(mem *AwooMemory, n arch.AwooRegister) arch.AwooWord {
	return arch.AwooWord(ReadMemory32(mem, n))
}

func ReadMemoryWordHalf(mem *AwooMemory, n arch.AwooRegister) arch.AwooWordHalf {
	return arch.AwooWordHalf(ReadMemory16(mem, n))
}
