package compiler_context

type AwooCompilerContextMemory struct {
	Memory   map[string]AwooCompilerContextMemoryEntry
	Position uint16
}

type AwooCompilerContextMemoryEntry struct {
	Start uint16
	Type  uint16
}

func SetContextMemory(context *AwooCompilerContext, name string, t uint16) error {
	context.Memory.Memory[name] = AwooCompilerContextMemoryEntry{
		Start: context.Memory.Position,
		Type:  t,
	}
	context.Memory.Position += context.Parser.Lexer.Types.All[t].Size

	return nil
}

func GetContextMemory(context *AwooCompilerContext, name string) (uint16, bool) {
	entry, ok := context.Memory.Memory[name]
	if !ok {
		return 0, false
	}

	return entry.Start, true
}
