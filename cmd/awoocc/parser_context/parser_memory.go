package parser_context

import "github.com/LamkasDev/awoo-emu/cmd/common/elf"

const AwooCompilerGlobalFunctionId = uint16(0)
const AwooCompilerGlobalBlockId = uint16(0)

func PushParserScopeBlockSymbol(context *AwooParserContext, funcId uint16, blockId uint16, blockEntry elf.AwooElfSymbolTableEntry) (elf.AwooElfSymbolTableEntry, bool) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	if _, ok := functionBlock.SymbolTable.Internal[blockEntry.Name]; ok {
		return elf.AwooElfSymbolTableEntry{}, false
	}
	functionBlock.SymbolTable.Internal[blockEntry.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, true
}

func PushParserScopeBlockSymbolExternal(context *AwooParserContext, funcId uint16, blockId uint16, blockEntry elf.AwooElfSymbolTableEntry) (elf.AwooElfSymbolTableEntry, bool) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	if _, ok := functionBlock.SymbolTable.External[blockEntry.Name]; ok {
		return elf.AwooElfSymbolTableEntry{}, false
	}
	functionBlock.SymbolTable.External[blockEntry.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, true
}

func PushParserScopeCurrentBlockSymbol(context *AwooParserContext, blockEntry elf.AwooElfSymbolTableEntry) (elf.AwooElfSymbolTableEntry, bool) {
	scopeFunction, ok := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	if !ok {
		return blockEntry, false
	}

	return PushParserScopeBlockSymbol(context, scopeFunction.Id, uint16(len(scopeFunction.Blocks)-1), blockEntry)
}

func PopParserScopeBlockSymbol(context *AwooParserContext, funcId uint16, blockId uint16, name string) bool {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	if _, ok := functionBlock.SymbolTable.Internal[name]; !ok {
		return false
	}
	delete(functionBlock.SymbolTable.Internal, name)
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return true
}

func PopParserScopeFunctionSymbol(context *AwooParserContext, name string) bool {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			ok := PopParserScopeBlockSymbol(context, uint16(funcId), uint16(blockId), name)
			if ok {
				return true
			}
		}
	}

	return false
}

func GetParserScopeBlockSymbol(context *AwooParserContext, funcId uint16, blockId uint16, name string) (elf.AwooElfSymbolTableEntry, bool) {
	blockEntry, ok := context.Scopes.Functions[funcId].Blocks[blockId].SymbolTable.Internal[name]
	if ok {
		return blockEntry, true
	}
	blockEntry, ok = context.Scopes.Functions[funcId].Blocks[blockId].SymbolTable.External[name]
	if ok {
		return blockEntry, true
	}

	return elf.AwooElfSymbolTableEntry{}, false
}

func GetParserScopeFunctionSymbol(context *AwooParserContext, name string) (elf.AwooElfSymbolTableEntry, bool) {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			blockEntry, ok := GetParserScopeBlockSymbol(context, uint16(funcId), uint16(blockId), name)
			if ok {
				return blockEntry, true
			}
		}
	}

	return elf.AwooElfSymbolTableEntry{}, false
}
