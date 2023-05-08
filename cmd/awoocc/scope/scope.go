package scope

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
)

const AwooScopeGlobalFunctionId = uint16(0)
const AwooScopeGlobalBlockId = uint16(0)

type AwooScopeContainer struct {
	Functions      map[uint16]AwooScopeFunction
	NextFunctionId uint16
}

func NewScopeContainer() AwooScopeContainer {
	return AwooScopeContainer{
		Functions: map[uint16]AwooScopeFunction{},
	}
}

type AwooScopeFunction struct {
	Id          uint16
	Name        string
	Blocks      map[uint16]AwooScopeBlock
	NextBlockId uint16
}

func NewScopeFunction(name string) AwooScopeFunction {
	return AwooScopeFunction{
		Name:   name,
		Blocks: map[uint16]AwooScopeBlock{},
	}
}

type AwooScopeBlock struct {
	Id          uint16
	Name        string
	SymbolTable AwooScopeSymbolTable
}

func NewScopeBlock(name string) AwooScopeBlock {
	return AwooScopeBlock{
		Name:        name,
		SymbolTable: NewScopeSymbolTable(0),
	}
}

func IsFunctionGlobal(scopeFunction AwooScopeFunction) bool {
	return scopeFunction.Name == cc.AwooCompilerGlobalFunctionName
}

func PushFunction(container *AwooScopeContainer, scopeFunction AwooScopeFunction) AwooScopeFunction {
	scopeFunction.Id = container.NextFunctionId
	scopeFunction.Blocks = map[uint16]AwooScopeBlock{}
	container.Functions[scopeFunction.Id] = scopeFunction
	container.NextFunctionId++
	PushBlock(container, scopeFunction.Id, NewScopeBlock("body"))

	return scopeFunction
}

func GetCurrentFunction(container *AwooScopeContainer) AwooScopeFunction {
	return container.Functions[container.NextFunctionId-1]
}

func GetCurrentFunctionSize(container *AwooScopeContainer) arch.AwooRegister {
	return GetCurrentFunctionBlock(container).SymbolTable.Position
}

func PopCurrentFunction(container *AwooScopeContainer) {
	delete(container.Functions, container.NextFunctionId-1)
}

func PushBlock(container *AwooScopeContainer, funcId uint16, functionBlock AwooScopeBlock) AwooScopeBlock {
	scopeFunction := container.Functions[funcId]
	functionBlock.Id = scopeFunction.NextBlockId
	functionBlock.SymbolTable = NewScopeSymbolTable(GetCurrentFunctionSize(container))
	scopeFunction.Blocks[functionBlock.Id] = functionBlock
	scopeFunction.NextBlockId++
	container.Functions[funcId] = scopeFunction

	return functionBlock
}

func PushCurrentFunctionBlock(container *AwooScopeContainer, functionBlock AwooScopeBlock) AwooScopeBlock {
	return PushBlock(container, container.NextFunctionId-1, functionBlock)
}

func GetCurrentFunctionBlock(container *AwooScopeContainer) AwooScopeBlock {
	scopeFunction := container.Functions[container.NextFunctionId-1]
	return scopeFunction.Blocks[scopeFunction.NextBlockId-1]
}

func PopCurrentFunctionBlock(container *AwooScopeContainer) {
	scopeFunction := container.Functions[container.NextFunctionId-1]
	delete(scopeFunction.Blocks, scopeFunction.NextBlockId-1)
	scopeFunction.NextBlockId--
	container.Functions[container.NextFunctionId-1] = scopeFunction
}
