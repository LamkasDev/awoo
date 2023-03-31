package compiler_context

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type AwooCompilerScopeContainer struct {
	Functions map[uint16]AwooCompilerScopeFunction
}

type AwooCompilerScopeFunction struct {
	Id     uint16
	Name   string
	Blocks map[uint16]AwooCompilerScopeBlock
	Global bool
}

type AwooCompilerScopeBlock struct {
	Id     uint16
	Name   string
	Memory AwooCompilerMemory
}

func PushCompilerScopeFunction(context *AwooCompilerContext, scopeFunction AwooCompilerScopeFunction) AwooCompilerScopeFunction {
	scopeFunction.Id = uint16(len(context.Scopes.Functions))
	scopeFunction.Blocks = map[uint16]AwooCompilerScopeBlock{}
	context.Scopes.Functions[scopeFunction.Id] = scopeFunction
	PushCompilerScopeBlock(context, scopeFunction.Id, AwooCompilerScopeBlock{
		Name: "body",
	})

	return scopeFunction
}

func PushCompilerScopeBlock(context *AwooCompilerContext, funcId uint16, functionBlock AwooCompilerScopeBlock) AwooCompilerScopeBlock {
	scopeFunction := context.Scopes.Functions[funcId]
	functionBlockPosition := arch.AwooRegister(0)
	if len(scopeFunction.Blocks) > 0 {
		functionBlockPosition = scopeFunction.Blocks[uint16(len(scopeFunction.Blocks)-1)].Memory.Position
	}
	functionBlock.Id = uint16(len(scopeFunction.Blocks))
	functionBlock.Memory = AwooCompilerMemory{
		Entries:  map[string]AwooCompilerMemoryEntry{},
		Position: functionBlockPosition,
	}
	scopeFunction.Blocks[functionBlock.Id] = functionBlock
	context.Scopes.Functions[funcId] = scopeFunction

	return functionBlock
}

func PushCompilerScopeCurrentBlock(context *AwooCompilerContext, functionBlock AwooCompilerScopeBlock) AwooCompilerScopeBlock {
	return PushCompilerScopeBlock(context, uint16(len(context.Scopes.Functions)-1), functionBlock)
}

func PopCompilerScopeCurrentFunction(context *AwooCompilerContext) {
	delete(context.Scopes.Functions, uint16(len(context.Scopes.Functions)-1))
}

func PopCompilerScopeCurrentBlock(context *AwooCompilerContext) {
	funcId := uint16(len(context.Scopes.Functions) - 1)
	scopeFunction := context.Scopes.Functions[funcId]
	delete(scopeFunction.Blocks, uint16(len(scopeFunction.Blocks)-1))
	context.Scopes.Functions[funcId] = scopeFunction
}

func GetCompilerScopeCurrentFunction(context *AwooCompilerContext) AwooCompilerScopeFunction {
	return context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
}

func GetCompilerScopeCurrentFunctionSize(context *AwooCompilerContext) arch.AwooRegister {
	scopeFunction := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	return scopeFunction.Blocks[uint16(len(scopeFunction.Blocks)-1)].Memory.Position
}

func SetupCompilerScopeContainer() AwooCompilerScopeContainer {
	container := AwooCompilerScopeContainer{
		Functions: map[uint16]AwooCompilerScopeFunction{},
	}

	return container
}
