package compiler_context

type AwooCompilerScopeContainer struct {
	Global    AwooCompilerMemory
	Functions map[uint16]AwooCompilerScopeFunction
}

type AwooCompilerScopeFunction struct {
	Id     uint16
	Name   string
	Blocks map[uint16]AwooCompilerScopeBlock
}

type AwooCompilerScopeBlock struct {
	Id     uint16
	Name   string
	Memory AwooCompilerMemory
}

func PushCompilerScopeFunction(context *AwooCompilerContext, name string) uint16 {
	scope := AwooCompilerScopeFunction{
		Id:     uint16(len(context.Scopes.Functions)),
		Name:   name,
		Blocks: map[uint16]AwooCompilerScopeBlock{},
	}
	context.Scopes.Functions[scope.Id] = scope
	PushCompilerScopeBlock(context, scope.Id, "body")

	return scope.Id
}

func PushCompilerScopeBlock(context *AwooCompilerContext, funcId uint16, name string) uint16 {
	f := context.Scopes.Functions[funcId]
	pos := uint16(0)
	if len(f.Blocks) > 0 {
		pos = f.Blocks[uint16(len(f.Blocks)-1)].Memory.Position
	}
	scope := AwooCompilerScopeBlock{
		Id:   uint16(len(f.Blocks)),
		Name: name,
		Memory: AwooCompilerMemory{
			Entries:  map[string]AwooCompilerContextMemoryEntry{},
			Position: pos,
		},
	}
	f.Blocks[scope.Id] = scope
	context.Scopes.Functions[funcId] = f

	return scope.Id
}

func PushCompilerScopeCurrentBlock(context *AwooCompilerContext, name string) uint16 {
	return PushCompilerScopeBlock(context, uint16(len(context.Scopes.Functions)-1), name)
}

func PopCompilerScopeCurrentFunction(context *AwooCompilerContext) {
	delete(context.Scopes.Functions, uint16(len(context.Scopes.Functions)-1))
}

func PopCompilerScopeCurrentBlock(context *AwooCompilerContext) {
	funcId := uint16(len(context.Scopes.Functions) - 1)
	f := context.Scopes.Functions[funcId]
	delete(f.Blocks, uint16(len(f.Blocks)-1))
	context.Scopes.Functions[funcId] = f
}

func GetCompilerScopeCurrentFunctionSize(context *AwooCompilerContext) uint16 {
	f := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	return f.Blocks[uint16(len(f.Blocks)-1)].Memory.Position
}

func SetupCompilerScopeContainer() AwooCompilerScopeContainer {
	container := AwooCompilerScopeContainer{
		Global: AwooCompilerMemory{
			Entries: map[string]AwooCompilerContextMemoryEntry{},
		},
		Functions: map[uint16]AwooCompilerScopeFunction{},
	}

	return container
}
