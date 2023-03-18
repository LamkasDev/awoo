package parser_context

type AwooParserScopeContainer struct {
	Global    AwooParserMemory
	Functions map[uint16]AwooParserScopeFunction
}

type AwooParserScopeFunction struct {
	Id     uint16
	Name   string
	Blocks map[uint16]AwooParserScopeBlock
}

type AwooParserScopeBlock struct {
	Id     uint16
	Name   string
	Memory AwooParserMemory
}

func PushParserScopeFunction(context *AwooParserContext, function AwooParserScopeFunction) AwooParserScopeFunction {
	function.Id = uint16(len(context.Scopes.Functions))
	function.Blocks = map[uint16]AwooParserScopeBlock{}
	context.Scopes.Functions[function.Id] = function
	PushParserScopeBlock(context, function.Id, AwooParserScopeBlock{
		Name: "body",
	})

	return function
}

func PushParserScopeBlock(context *AwooParserContext, funcId uint16, functionBlock AwooParserScopeBlock) AwooParserScopeBlock {
	scopeFunction := context.Scopes.Functions[funcId]
	functionBlock.Id = uint16(len(scopeFunction.Blocks))
	functionBlock.Memory = AwooParserMemory{
		Entries: map[string]AwooParserMemoryEntry{},
	}
	scopeFunction.Blocks[functionBlock.Id] = functionBlock
	context.Scopes.Functions[funcId] = scopeFunction

	return functionBlock
}

func PushParserScopeCurrentBlock(context *AwooParserContext, functionBlock AwooParserScopeBlock) AwooParserScopeBlock {
	return PushParserScopeBlock(context, uint16(len(context.Scopes.Functions)-1), functionBlock)
}

func PopParserScopeCurrentFunction(context *AwooParserContext) {
	delete(context.Scopes.Functions, uint16(len(context.Scopes.Functions)-1))
}

func PopParserScopeCurrentBlock(context *AwooParserContext) {
	funcId := uint16(len(context.Scopes.Functions) - 1)
	scopeFunction := context.Scopes.Functions[funcId]
	delete(scopeFunction.Blocks, uint16(len(scopeFunction.Blocks)-1))
	context.Scopes.Functions[funcId] = scopeFunction
}

func SetupParserScopeContainer() AwooParserScopeContainer {
	container := AwooParserScopeContainer{
		Global: AwooParserMemory{
			Entries: map[string]AwooParserMemoryEntry{},
		},
		Functions: map[uint16]AwooParserScopeFunction{},
	}

	return container
}
