package parser_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/jwalton/gchalk"
)

type AwooParserMemory struct {
	Entries map[string]AwooParserMemoryEntry
}

type AwooParserMemoryEntry struct {
	Name   string
	Global bool
	Type   uint16
	Data   interface{}
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
		blockEntry.Global = true
		context.Scopes.Global.Entries[blockEntry.Name] = blockEntry
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

func PopParserScopeFunctionMemory(context *AwooParserContext, funcId uint16, name string) error {
	_, ok := context.Scopes.Global.Entries[name]
	if ok {
		delete(context.Scopes.Global.Entries, name)
		return nil
	}
	for blockId := len(context.Scopes.Functions[funcId].Blocks); blockId >= 0; blockId-- {
		err := PopParserScopeBlockMemory(context, funcId, uint16(blockId), name)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func PopParserScopeCurrentFunctionMemory(context *AwooParserContext, name string) error {
	return PopParserScopeFunctionMemory(context, uint16(len(context.Scopes.Functions)-1), name)
}

func GetParserScopeBlockMemory(context *AwooParserContext, funcId uint16, blockId uint16, name string) (AwooParserMemoryEntry, error) {
	blockEntry, ok := context.Scopes.Functions[funcId].Blocks[blockId].Memory.Entries[name]
	if !ok {
		return AwooParserMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}

	return blockEntry, nil
}

func GetParserScopeFunctionMemory(context *AwooParserContext, funcId uint16, name string) (AwooParserMemoryEntry, error) {
	globalEntry, ok := context.Scopes.Global.Entries[name]
	if ok {
		return globalEntry, nil
	}
	for blockId := len(context.Scopes.Functions[funcId].Blocks); blockId >= 0; blockId-- {
		blockEntry, err := GetParserScopeBlockMemory(context, funcId, uint16(blockId), name)
		if err == nil {
			return blockEntry, nil
		}
	}

	return AwooParserMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetParserScopeCurrentFunctionMemory(context *AwooParserContext, name string) (AwooParserMemoryEntry, error) {
	return GetParserScopeFunctionMemory(context, uint16(len(context.Scopes.Functions)-1), name)
}
