package compiler_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/jwalton/gchalk"
)

type AwooCompilerMemory struct {
	Entries  map[string]AwooCompilerMemoryEntry
	Position uint32
}

type AwooCompilerMemoryEntry struct {
	Name   string
	Global bool
	Type   uint16
	Data   interface{}
	Start  uint32
	Size   uint32
}

func PushCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, blockEntry AwooCompilerMemoryEntry) (AwooCompilerMemoryEntry, error) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	blockEntry.Start = functionBlock.Memory.Position
	functionBlock.Memory.Position += blockEntry.Size
	functionBlock.Memory.Entries[blockEntry.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, nil
}

func PushCompilerScopeCurrentBlockMemory(context *AwooCompilerContext, blockEntry AwooCompilerMemoryEntry) (AwooCompilerMemoryEntry, error) {
	scopeFunction, ok := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	if !ok {
		blockEntry.Global = true
		blockEntry.Start = context.Scopes.Global.Position
		context.Scopes.Global.Entries[blockEntry.Name] = blockEntry
		context.Scopes.Global.Position += blockEntry.Size
		return blockEntry, nil
	}

	return PushCompilerScopeBlockMemory(context, scopeFunction.Id, uint16(len(scopeFunction.Blocks)-1), blockEntry)
}

func PopCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, name string) error {
	blockEntry, err := GetCompilerScopeBlockMemory(context, funcId, blockId, name)
	if err != nil {
		return err
	}

	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	functionBlock.Memory.Position -= blockEntry.Size
	delete(functionBlock.Memory.Entries, name)
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return nil
}

func PopCompilerScopeFunctionMemory(context *AwooCompilerContext, funcId uint16, name string) error {
	globalEntry, ok := context.Scopes.Global.Entries[name]
	if ok {
		context.Scopes.Global.Position -= globalEntry.Size
		delete(context.Scopes.Global.Entries, name)
		return nil
	}
	for blockId := len(context.Scopes.Functions[funcId].Blocks); blockId >= 0; blockId-- {
		err := PopCompilerScopeBlockMemory(context, funcId, uint16(blockId), name)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func PopCompilerScopeCurrentFunctionMemory(context *AwooCompilerContext, name string) error {
	return PopCompilerScopeFunctionMemory(context, uint16(len(context.Scopes.Functions)-1), name)
}

func GetCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, name string) (AwooCompilerMemoryEntry, error) {
	blockEntry, ok := context.Scopes.Functions[funcId].Blocks[blockId].Memory.Entries[name]
	if !ok {
		return AwooCompilerMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}

	return blockEntry, nil
}

func GetCompilerScopeFunctionMemory(context *AwooCompilerContext, funcId uint16, name string) (AwooCompilerMemoryEntry, error) {
	globalEntry, ok := context.Scopes.Global.Entries[name]
	if ok {
		return globalEntry, nil
	}
	for blockId := len(context.Scopes.Functions[funcId].Blocks); blockId >= 0; blockId-- {
		blockEntry, err := GetCompilerScopeBlockMemory(context, funcId, uint16(blockId), name)
		if err == nil {
			return blockEntry, nil
		}
	}

	return AwooCompilerMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetCompilerScopeCurrentFunctionMemory(context *AwooCompilerContext, name string) (AwooCompilerMemoryEntry, error) {
	return GetCompilerScopeFunctionMemory(context, uint16(len(context.Scopes.Functions)-1), name)
}
