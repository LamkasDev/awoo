package compiler_context

type AwooCompilerFunctionContainer struct {
	Entries map[string]AwooCompilerFunction
}

type AwooCompilerFunction struct {
	Name  string
	Start uint16
	Size  uint16
}

func PushCompilerFunction(context *AwooCompilerContext, entry AwooCompilerFunction) {
	context.Functions.Entries[entry.Name] = entry
}

func GetCompilerFunction(context *AwooCompilerContext, name string) (AwooCompilerFunction, bool) {
	f, ok := context.Functions.Entries[name]
	return f, ok
}
