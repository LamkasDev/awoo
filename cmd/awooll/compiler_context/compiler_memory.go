package compiler_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/jwalton/gchalk"
)

const AwooCompilerContextMemoryPageSize = 1024

type AwooCompilerMemory struct {
	Entries  map[string]AwooCompilerContextMemoryEntry
	Position uint16
}

type AwooCompilerContextMemoryEntry struct {
	Name  string
	Start uint16
	Type  uint16
}

func PushCompilerScopeIdMemory(context *AwooCompilerContext, scopeId uint16, name string, t uint16) (uint16, error) {
	// TODO: error checking.
	scope := context.Scopes.Entries[scopeId]
	start := context.Scopes.Entries[scopeId].Memory.Position
	scope.Memory.Position += context.Parser.Lexer.Types.All[t].Size
	scope.Memory.Entries[name] = AwooCompilerContextMemoryEntry{
		Name:  name,
		Start: start,
		Type:  t,
	}
	context.Scopes.Entries[scopeId] = scope

	return start, nil
}

func PushCompilerScopeCurrentMemory(context *AwooCompilerContext, name string, t uint16) (uint16, error) {
	return PushCompilerScopeIdMemory(context, context.Scopes.Position, name, t)
}

func PopCompilerScopeIdMemory(context *AwooCompilerContext, scopeId uint16, name string) error {
	entry, err := GetCompilerScopeIdMemory(context, scopeId, name)
	if err != nil {
		return err
	}

	scope := context.Scopes.Entries[scopeId]
	scope.Memory.Position -= context.Parser.Lexer.Types.All[entry.Type].Size
	delete(scope.Memory.Entries, name)
	context.Scopes.Entries[scopeId] = scope

	return nil
}

func PopCompilerScopeCurrentMemory(context *AwooCompilerContext, name string) error {
	return PopCompilerScopeIdMemory(context, context.Scopes.Position, name)
}

func PopCompilerScopeMemory(context *AwooCompilerContext, name string) error {
	for i := context.Scopes.Position; i > 0; i-- {
		err := PopCompilerScopeIdMemory(context, i, name)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetCompilerScopeIdMemory(context *AwooCompilerContext, scopeId uint16, name string) (AwooCompilerContextMemoryEntry, error) {
	entry, ok := context.Scopes.Entries[scopeId].Memory.Entries[name]
	if !ok {
		return AwooCompilerContextMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}

	return entry, nil
}

func GetCompilerScopeCurrentMemory(context *AwooCompilerContext, name string) (AwooCompilerContextMemoryEntry, error) {
	return GetCompilerScopeIdMemory(context, context.Scopes.Position, name)
}

func GetCompilerScopeMemory(context *AwooCompilerContext, name string) (AwooCompilerContextMemoryEntry, error) {
	for i := context.Scopes.Position; i > 0; i-- {
		dest, err := GetCompilerScopeIdMemory(context, context.Scopes.Position, name)
		if err == nil {
			return dest, nil
		}
	}

	return AwooCompilerContextMemoryEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}
