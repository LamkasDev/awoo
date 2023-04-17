package parser_context

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

func PushParserScopeBlockMemory(context *AwooParserContext, funcId uint16, blockId uint16, blockEntry AwooParserMemoryEntry) (AwooParserMemoryEntry, bool) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	if _, ok := functionBlock.Memory.Entries[blockEntry.Name]; ok {
		return AwooParserMemoryEntry{}, false
	}
	functionBlock.Memory.Entries[blockEntry.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, true
}

func PushParserScopeCurrentBlockMemory(context *AwooParserContext, blockEntry AwooParserMemoryEntry) (AwooParserMemoryEntry, bool) {
	scopeFunction, ok := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	if !ok {
		return blockEntry, false
	}

	return PushParserScopeBlockMemory(context, scopeFunction.Id, uint16(len(scopeFunction.Blocks)-1), blockEntry)
}

func PopParserScopeBlockMemory(context *AwooParserContext, funcId uint16, blockId uint16, name string) bool {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	if _, ok := functionBlock.Memory.Entries[name]; !ok {
		return false
	}
	delete(functionBlock.Memory.Entries, name)
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return true
}

func PopParserScopeFunctionMemory(context *AwooParserContext, name string) bool {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			ok := PopParserScopeBlockMemory(context, uint16(funcId), uint16(blockId), name)
			if ok {
				return true
			}
		}
	}

	return false
}

func GetParserScopeBlockMemory(context *AwooParserContext, funcId uint16, blockId uint16, name string) (AwooParserMemoryEntry, bool) {
	blockEntry, ok := context.Scopes.Functions[funcId].Blocks[blockId].Memory.Entries[name]
	if !ok {
		return AwooParserMemoryEntry{}, false
	}

	return blockEntry, true
}

func GetParserScopeFunctionMemory(context *AwooParserContext, name string) (AwooParserMemoryEntry, bool) {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			blockEntry, ok := GetParserScopeBlockMemory(context, uint16(funcId), uint16(blockId), name)
			if ok {
				return blockEntry, true
			}
		}
	}

	return AwooParserMemoryEntry{}, false
}
