package memory

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooMemory struct {
	Data []byte
}

func SetupMemory(n arch.AwooRegisterU) AwooMemory {
	return AwooMemory{
		Data: make([]byte, n),
	}
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
}

func WriteMemory32(mem *AwooMemory, n arch.AwooRegister, data int32) {
	mem.Data[n] = byte(data >> 24)
	mem.Data[n+1] = byte(data >> 16)
	mem.Data[n+2] = byte(data >> 8)
	mem.Data[n+3] = byte(data)
}

func WriteMemory16(mem *AwooMemory, n arch.AwooRegister, data int16) {
	mem.Data[n] = byte(data >> 8)
	mem.Data[n+1] = byte(data)
}

func WriteMemory8(mem *AwooMemory, n arch.AwooRegister, data byte) {
	mem.Data[n] = byte(data)
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

	return data
}

func ReadMemory32(mem *AwooMemory, n arch.AwooRegister) int32 {
	data := int32(mem.Data[n]) << 24
	data |= int32(mem.Data[n+1]) << 16
	data |= int32(mem.Data[n+2]) << 8
	data |= int32(mem.Data[n+3])

	return data
}

func ReadMemory16(mem *AwooMemory, n arch.AwooRegister) int16 {
	data := int16(mem.Data[n]) << 8
	data |= int16(mem.Data[n+1])

	return data
}

func ReadMemory8(mem *AwooMemory, n arch.AwooRegister) int8 {
	return int8(mem.Data[n])
}
