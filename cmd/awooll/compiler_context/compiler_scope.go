package compiler_context

type AwooCompilerScopeContainer struct {
	Entries  map[uint16]AwooCompilerScope
	Position uint16
}

type AwooCompilerScope struct {
	Id     uint16
	Name   string
	Memory AwooCompilerMemory
}

func PushCompilerScope(container *AwooCompilerScopeContainer, name string) uint16 {
	prev, ok := container.Entries[container.Position]
	pos := uint16(0)
	if ok {
		pos = prev.Memory.Position
	}

	container.Position++
	scope := AwooCompilerScope{
		Id:   container.Position,
		Name: name,
		Memory: AwooCompilerMemory{
			Entries:  make(map[string]AwooCompilerContextMemoryEntry),
			Position: pos,
		},
	}
	container.Entries[scope.Id] = scope

	return scope.Id
}

func PopCompilerScope(container *AwooCompilerScopeContainer) {
	delete(container.Entries, container.Position)
	container.Position--
}

func SetupCompilerScopeContainer() AwooCompilerScopeContainer {
	container := AwooCompilerScopeContainer{
		Entries: make(map[uint16]AwooCompilerScope),
	}
	PushCompilerScope(&container, "global")

	return container
}
