package memory

import (
	"sync"

	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"golang.org/x/exp/constraints"
)

type WriteMemoryFunc[K constraints.Integer] func(mem *AwooMemory, n arch.AwooRegister, data K)
type ReadMemoryFunc[K constraints.Integer] func(mem *AwooMemory, n arch.AwooRegister) K

type AwooMemory struct {
	Data       []byte
	Lockable   []AwooMemoryLockable
	ProgramEnd arch.AwooRegister
	TotalRead  uint64
	TotalWrite uint64
}

type AwooMemoryLockable struct {
	Start arch.AwooRegister
	End   arch.AwooRegister
	Lock  sync.Mutex
}

func SetupMemory(n arch.AwooRegisterU) AwooMemory {
	return AwooMemory{
		Data: make([]byte, n),
	}
}

func WriteMemorySafe[K constraints.Integer](mem *AwooMemory, n arch.AwooRegister, data K, write WriteMemoryFunc[K]) {
	for i := 0; i < len(mem.Lockable); i++ {
		if n >= mem.Lockable[i].Start && n <= mem.Lockable[i].End {
			mem.Lockable[i].Lock.Lock()
			defer mem.Lockable[i].Lock.Unlock()
			write(mem, n, data)
		}
	}
	write(mem, n, data)
}

func WriteMemory64(mem *AwooMemory, n arch.AwooRegister, data int64) {
	mem.Data[n] = byte(data >> 56)
	mem.Data[n+1] = byte(data >> 48)
	mem.Data[n+2] = byte(data >> 40)
	mem.Data[n+3] = byte(data >> 32)
	mem.Data[n+4] = byte(data >> 24)
	mem.Data[n+5] = byte(data >> 16)
	mem.Data[n+6] = byte(data >> 8)
	mem.Data[n+7] = byte(data)
	mem.TotalWrite += 8
}

func WriteMemory32(mem *AwooMemory, n arch.AwooRegister, data int32) {
	mem.Data[n] = byte(data >> 24)
	mem.Data[n+1] = byte(data >> 16)
	mem.Data[n+2] = byte(data >> 8)
	mem.Data[n+3] = byte(data)
	mem.TotalWrite += 4
}

func WriteMemory16(mem *AwooMemory, n arch.AwooRegister, data int16) {
	mem.Data[n] = byte(data >> 8)
	mem.Data[n+1] = byte(data)
	mem.TotalWrite += 2
}

func WriteMemory8(mem *AwooMemory, n arch.AwooRegister, data int8) {
	mem.Data[n] = byte(data)
	mem.TotalWrite++
}

func ReadMemorySafe[K constraints.Integer](mem *AwooMemory, n arch.AwooRegister, read ReadMemoryFunc[K]) K {
	for i := 0; i < len(mem.Lockable); i++ {
		if n >= mem.Lockable[i].Start && n <= mem.Lockable[i].End {
			mem.Lockable[i].Lock.Lock()
			defer mem.Lockable[i].Lock.Unlock()
			return read(mem, n)
		}
	}
	return read(mem, n)
}

func ReadMemory64(mem *AwooMemory, n arch.AwooRegister) int64 {
	data := int64(mem.Data[n]) << 56
	data |= int64(mem.Data[n+1]) << 48
	data |= int64(mem.Data[n+2]) << 40
	data |= int64(mem.Data[n+3]) << 32
	data |= int64(mem.Data[n+4]) << 24
	data |= int64(mem.Data[n+5]) << 16
	data |= int64(mem.Data[n+6]) << 8
	data |= int64(mem.Data[n+7])
	mem.TotalRead += 8

	return data
}

func ReadMemory32(mem *AwooMemory, n arch.AwooRegister) int32 {
	data := int32(mem.Data[n]) << 24
	data |= int32(mem.Data[n+1]) << 16
	data |= int32(mem.Data[n+2]) << 8
	data |= int32(mem.Data[n+3])
	mem.TotalRead += 4

	return data
}

func ReadMemory16(mem *AwooMemory, n arch.AwooRegister) int16 {
	data := int16(mem.Data[n]) << 8
	data |= int16(mem.Data[n+1])
	mem.TotalRead += 2

	return data
}

func ReadMemory8(mem *AwooMemory, n arch.AwooRegister) int8 {
	mem.TotalRead++
	return int8(mem.Data[n])
}
