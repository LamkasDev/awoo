package parser_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
	"github.com/jwalton/gchalk"
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

func PushParserScopeBlockMemory(context *AwooParserContext, funcId uint16, blockId uint16, blockEntry AwooParserMemoryEntry) (AwooParserMemoryEntry, error) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	functionBlock.Memory.Entries[blockEntry.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, nil
}

func PushParserScopeCurrentBlockMemory(context *AwooParserContext, blockEntry AwooParserMemoryEntry) (AwooParserMemoryEntry, error) {
	scopeFunction, ok := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	if !ok {
		return blockEntry, nil
	}

	return PushParserScopeBlockMemory(context, scopeFunction.Id, uint16(len(scopeFunction.Blocks)-1), blockEntry)
}

func PopParserScopeBlockMemory(context *AwooParserContext, funcId uint16, blockId uint16, name string) error {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	delete(functionBlock.Memory.Entries, name)
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return nil
}

func PopParserScopeFunctionMemory(context *AwooParserContext, name string) error {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			err := PopParserScopeBlockMemory(context, uint16(funcId), uint16(blockId), name)
			if err == nil {
				return nil
			}
		}
	}

	return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetParserScopeBlockMemory(context *AwooParserContext, funcId uint16, blockId uint16, name string) (AwooParserMemoryEntry, error) {
	blockEntry, ok := context.Scopes.Functions[funcId].Blocks[blockId].Memory.Entries[name]
	if !ok {
		return AwooParserMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}

	return blockEntry, nil
}

func GetParserScopeFunctionMemory(context *AwooParserContext, name string) (AwooParserMemoryEntry, error) {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			blockEntry, err := GetParserScopeBlockMemory(context, uint16(funcId), uint16(blockId), name)
			if err == nil {
				return blockEntry, nil
			}
		}
	}

	return AwooParserMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}
