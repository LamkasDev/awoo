package elf

type AwooElfRelocationList []AwooElfRelocationListEntry

type AwooElfRelocationListEntry struct {
	Offset uint32
	Name   string
}
