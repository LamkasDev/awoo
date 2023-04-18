package parser_memory

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooParserMemory struct {
	Entries map[string]AwooParserMemoryEntry
}

type AwooParserMemoryEntry struct {
	Name        string
	Global      bool
	Type        types.AwooTypeId
	TypeDetails *types.AwooTypeId
}
