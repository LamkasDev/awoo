package compiler_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_memory"
	"github.com/jwalton/gchalk"
)

func PushCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, blockEntry compiler_memory.AwooCompilerMemoryEntry) (compiler_memory.AwooCompilerMemoryEntry, error) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	blockEntry.Symbol.Start = functionBlock.Memory.Position
	functionBlock.Memory.Position += blockEntry.Symbol.Size
	functionBlock.Memory.Entries[blockEntry.Symbol.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, nil
}

func PushCompilerScopeCurrentBlockMemory(context *AwooCompilerContext, blockEntry compiler_memory.AwooCompilerMemoryEntry) (compiler_memory.AwooCompilerMemoryEntry, error) {
	scopeFunction, ok := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	if !ok {
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
	functionBlock.Memory.Position -= blockEntry.Symbol.Size
	delete(functionBlock.Memory.Entries, name)
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return nil
}

func PopCompilerScopeFunctionMemory(context *AwooCompilerContext, name string) error {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			err := PopCompilerScopeBlockMemory(context, uint16(funcId), uint16(blockId), name)
			if err == nil {
				return nil
			}
		}
	}

	return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, name string) (compiler_memory.AwooCompilerMemoryEntry, error) {
	blockEntry, ok := context.Scopes.Functions[funcId].Blocks[blockId].Memory.Entries[name]
	if !ok {
		return compiler_memory.AwooCompilerMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}

	return blockEntry, nil
}

func GetCompilerScopeFunctionMemory(context *AwooCompilerContext, name string) (compiler_memory.AwooCompilerMemoryEntry, error) {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			blockEntry, err := GetCompilerScopeBlockMemory(context, uint16(funcId), uint16(blockId), name)
			if err == nil {
				return blockEntry, nil
			}
		}
	}

	return compiler_memory.AwooCompilerMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}
