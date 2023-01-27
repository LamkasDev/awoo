package compiler_context

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

func SetCompilerScopeIdMemory(context *AwooCompilerContext, scopeId uint16, name string, t uint16) (uint16, error) {
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

func SetCompilerScopeCurrentMemory(context *AwooCompilerContext, name string, t uint16) (uint16, error) {
	return SetCompilerScopeIdMemory(context, context.Scopes.Current, name, t)
}

func GetCompilerScopeIdMemory(context *AwooCompilerContext, scopeId uint16, name string) (uint16, bool) {
	entry, ok := context.Scopes.Entries[scopeId].Memory.Entries[name]
	if !ok {
		return 0, false
	}

	return entry.Start, true
}

func GetCompilerScopeCurrentMemory(context *AwooCompilerContext, name string) (uint16, bool) {
	return GetCompilerScopeIdMemory(context, context.Scopes.Current, name)
}

func GetCompilerScopeMemory(context *AwooCompilerContext, name string) (uint16, bool) {
	for i := context.Scopes.Current; i > 0; i-- {
		dest, ok := GetCompilerScopeIdMemory(context, context.Scopes.Current, name)
		if ok {
			return dest, true
		}
	}

	return 0, false
}
