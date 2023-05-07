package compiler_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_symbol"
	"github.com/jwalton/gchalk"
)

func PushCompilerScopeBlockSymbol(context *AwooCompilerContext, funcId uint16, blockId uint16, blockEntry compiler_symbol.AwooCompilerSymbolTableEntry) (compiler_symbol.AwooCompilerSymbolTableEntry, error) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	blockEntry.Symbol.Start = functionBlock.SymbolTable.Position
	functionBlock.SymbolTable.Position += blockEntry.Symbol.Size
	functionBlock.SymbolTable.Internal[blockEntry.Symbol.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, nil
}

func PushCompilerScopeBlockSymbolExternal(context *AwooCompilerContext, funcId uint16, blockId uint16, blockEntry compiler_symbol.AwooCompilerSymbolTableEntry) (compiler_symbol.AwooCompilerSymbolTableEntry, error) {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	functionBlock.SymbolTable.External[blockEntry.Symbol.Name] = blockEntry
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return blockEntry, nil
}

func PushCompilerScopeCurrentBlockSymbol(context *AwooCompilerContext, blockEntry compiler_symbol.AwooCompilerSymbolTableEntry) (compiler_symbol.AwooCompilerSymbolTableEntry, error) {
	scopeFunction, ok := context.Scopes.Functions[uint16(len(context.Scopes.Functions)-1)]
	if !ok {
		return blockEntry, nil
	}

	return PushCompilerScopeBlockSymbol(context, scopeFunction.Id, uint16(len(scopeFunction.Blocks)-1), blockEntry)
}

func PopCompilerScopeBlockSymbol(context *AwooCompilerContext, funcId uint16, blockId uint16, name string) error {
	functionBlock := context.Scopes.Functions[funcId].Blocks[blockId]
	if _, ok := functionBlock.SymbolTable.Internal[name]; !ok {
		return nil
	}
	delete(functionBlock.SymbolTable.Internal, name)
	context.Scopes.Functions[funcId].Blocks[blockId] = functionBlock

	return nil
}

func PopCompilerScopeFunctionSymbol(context *AwooCompilerContext, name string) error {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			err := PopCompilerScopeBlockSymbol(context, uint16(funcId), uint16(blockId), name)
			if err == nil {
				return nil
			}
		}
	}

	return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetCompilerScopeBlockSymbol(context *AwooCompilerContext, funcId uint16, blockId uint16, name string) (compiler_symbol.AwooCompilerSymbolTableEntry, error) {
	blockEntry, ok := context.Scopes.Functions[funcId].Blocks[blockId].SymbolTable.Internal[name]
	if ok {
		return blockEntry, nil
	}
	blockEntry, ok = context.Scopes.Functions[funcId].Blocks[blockId].SymbolTable.External[name]
	if ok {
		return blockEntry, nil
	}

	return compiler_symbol.AwooCompilerSymbolTableEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetCompilerScopeFunctionSymbol(context *AwooCompilerContext, name string) (compiler_symbol.AwooCompilerSymbolTableEntry, error) {
	for funcId := len(context.Scopes.Functions); funcId >= 0; funcId-- {
		for blockId := len(context.Scopes.Functions[uint16(funcId)].Blocks); blockId >= 0; blockId-- {
			blockEntry, err := GetCompilerScopeBlockSymbol(context, uint16(funcId), uint16(blockId), name)
			if err == nil {
				return blockEntry, nil
			}
		}
	}

	return compiler_symbol.AwooCompilerSymbolTableEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}
