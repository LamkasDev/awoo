package compiler_context

type AwooCompilerScopeContainer struct {
	Entries            map[uint16]AwooCompilerScope
	Current            uint16
	CurrentMemoryStart uint16
}

type AwooCompilerScope struct {
	Id     uint16
	Name   string
	Memory AwooCompilerMemory
}

func PushCompilerScope(container *AwooCompilerScopeContainer, name string) uint16 {
	container.Current++
	scope := AwooCompilerScope{
		Id:   container.Current,
		Name: name,
		Memory: AwooCompilerMemory{
			Entries: make(map[string]AwooCompilerContextMemoryEntry),
		},
	}
	container.Entries[scope.Id] = scope

	return scope.Id
}

func PopCompilerScope(container *AwooCompilerScopeContainer) {
	delete(container.Entries, container.Current)
	container.Current--
}

func SetupCompilerScopeContainer() AwooCompilerScopeContainer {
	container := AwooCompilerScopeContainer{
		Entries: make(map[uint16]AwooCompilerScope),
	}
	PushCompilerScope(&container, "global")

	return container
}
