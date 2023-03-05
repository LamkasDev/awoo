package compiler_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/jwalton/gchalk"
)

type AwooCompilerMemory struct {
	Entries  map[string]AwooCompilerContextMemoryEntry
	Position uint16
}

type AwooCompilerContextMemoryEntry struct {
	Name   string
	Global bool
	Start  uint16
	Size   uint16
	Type   uint16
	Data   interface{}
}

func PushCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, entry AwooCompilerContextMemoryEntry) (AwooCompilerContextMemoryEntry, error) {
	block := context.Scopes.Functions[funcId].Blocks[blockId]
	entry.Start = block.Memory.Position
	block.Memory.Position += entry.Size
	block.Memory.Entries[entry.Name] = entry
	context.Scopes.Functions[funcId].Blocks[blockId] = block

	return entry, nil
}

func PushCompilerScopeCurrentBlockMemory(context *AwooCompilerContext, entry AwooCompilerContextMemoryEntry) (AwooCompilerContextMemoryEntry, error) {
	f, ok := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	if !ok {
		entry.Start = context.Scopes.Global.Position
		entry.Global = true
		context.Scopes.Global.Position += entry.Size
		context.Scopes.Global.Entries[entry.Name] = entry
		return entry, nil
	}
	return PushCompilerScopeBlockMemory(context, f.Id, uint16(len(f.Blocks)-1), entry)
}

func PopCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, name string) error {
	entry, err := GetCompilerScopeBlockMemory(context, funcId, blockId, name)
	if err != nil {
		return err
	}

	block := context.Scopes.Functions[funcId].Blocks[blockId]
	block.Memory.Position -= entry.Size
	delete(block.Memory.Entries, name)
	context.Scopes.Functions[funcId].Blocks[blockId] = block

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

func GetCompilerScopeBlockMemory(context *AwooCompilerContext, funcId uint16, blockId uint16, name string) (AwooCompilerContextMemoryEntry, error) {
	entry, ok := context.Scopes.Functions[funcId].Blocks[blockId].Memory.Entries[name]
	if !ok {
		return AwooCompilerContextMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}

	return entry, nil
}

func GetCompilerScopeFunctionMemory(context *AwooCompilerContext, funcId uint16, name string) (AwooCompilerContextMemoryEntry, error) {
	globalEntry, ok := context.Scopes.Global.Entries[name]
	if ok {
		return globalEntry, nil
	}
	for blockId := len(context.Scopes.Functions[funcId].Blocks); blockId >= 0; blockId-- {
		dest, err := GetCompilerScopeBlockMemory(context, funcId, uint16(blockId), name)
		if err == nil {
			return dest, nil
		}
	}

	return AwooCompilerContextMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetCompilerScopeCurrentFunctionMemory(context *AwooCompilerContext, name string) (AwooCompilerContextMemoryEntry, error) {
	return GetCompilerScopeFunctionMemory(context, uint16(len(context.Scopes.Functions)-1), name)
}
