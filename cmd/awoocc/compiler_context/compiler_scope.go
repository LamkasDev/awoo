package compiler_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_symbol"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/common/arch"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
)

type AwooCompilerScopeContainer struct {
	Functions map[uint16]AwooCompilerScopeFunction
}

func NewCompilerScopeContainer() AwooCompilerScopeContainer {
	return AwooCompilerScopeContainer{
		Functions: map[uint16]AwooCompilerScopeFunction{},
	}
}

type AwooCompilerScopeFunction struct {
	Id     uint16
	Name   string
	Blocks map[uint16]AwooCompilerScopeBlock
}

func NewCompilerScopeFunction(name string) AwooCompilerScopeFunction {
	return AwooCompilerScopeFunction{
		Name:   name,
		Blocks: map[uint16]AwooCompilerScopeBlock{},
	}
}

type AwooCompilerScopeBlock struct {
	Id          uint16
	Name        string
	SymbolTable compiler_symbol.AwooCompilerSymbolTable
}

func IsCompilerScopeFunctionGlobal(scopeFunction AwooCompilerScopeFunction) bool {
	return scopeFunction.Name == cc.AwooCompilerGlobalFunctionName
}

func PushCompilerScopeFunction(context *AwooCompilerContext, scopeFunction AwooCompilerScopeFunction) AwooCompilerScopeFunction {
	scopeFunction.Id = uint16(len(context.Scopes.Functions))
	scopeFunction.Blocks = map[uint16]AwooCompilerScopeBlock{}
	context.Scopes.Functions[scopeFunction.Id] = scopeFunction
	PushCompilerScopeBlock(context, scopeFunction.Id, AwooCompilerScopeBlock{
		Name: "body",
	})
	if IsCompilerScopeFunctionGlobal(scopeFunction) {
		for _, symbol := range context.Parser.Scopes.Functions[parser_context.AwooCompilerGlobalFunctionId].Blocks[parser_context.AwooCompilerGlobalBlockId].SymbolTable.External {
			PushCompilerScopeBlockSymbolExternal(context, parser_context.AwooCompilerGlobalFunctionId, parser_context.AwooCompilerGlobalBlockId, compiler_symbol.AwooCompilerSymbolTableEntry{
				Symbol: symbol,
				Global: true,
			})
		}
	}

	return scopeFunction
}

func PushCompilerScopeBlock(context *AwooCompilerContext, funcId uint16, functionBlock AwooCompilerScopeBlock) AwooCompilerScopeBlock {
	scopeFunction := context.Scopes.Functions[funcId]
	functionBlockPosition := arch.AwooRegister(0)
	if len(scopeFunction.Blocks) > 0 {
		functionBlockPosition = scopeFunction.Blocks[uint16(len(scopeFunction.Blocks)-1)].SymbolTable.Position
	}
	functionBlock.Id = uint16(len(scopeFunction.Blocks))
	functionBlock.SymbolTable = compiler_symbol.AwooCompilerSymbolTable{
		Internal: map[string]compiler_symbol.AwooCompilerSymbolTableEntry{},
		External: map[string]compiler_symbol.AwooCompilerSymbolTableEntry{},
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
	return scopeFunction.Blocks[uint16(len(scopeFunction.Blocks)-1)].SymbolTable.Position
}
