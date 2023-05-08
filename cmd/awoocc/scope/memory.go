package scope

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/common/elf"
	"github.com/jwalton/gchalk"
)

func PushFunctionBlockSymbol(container *AwooScopeContainer, funcId uint16, blockId uint16, symbol elf.AwooElfSymbolTableEntry) (AwooScopeSymbolTableEntry, error) {
	functionBlock := container.Functions[funcId].Blocks[blockId]
	symbol.Start = functionBlock.SymbolTable.Position
	scopeSymbol := AwooScopeSymbolTableEntry{
		Symbol: symbol,
		Global: IsFunctionGlobal(GetCurrentFunction(container)),
	}
	if _, ok := functionBlock.SymbolTable.Internal[symbol.Name]; ok {
		return scopeSymbol, fmt.Errorf("%w: %s", awerrors.ErrorAlreadyDefinedVariable, gchalk.Red(symbol.Name))
	}

	functionBlock.SymbolTable.Internal[symbol.Name] = scopeSymbol
	functionBlock.SymbolTable.Position += symbol.Size
	container.Functions[funcId].Blocks[blockId] = functionBlock

	return scopeSymbol, nil
}

func PushFunctionBlockSymbolExternal(container *AwooScopeContainer, symbol elf.AwooElfSymbolTableEntry) (AwooScopeSymbolTableEntry, error) {
	functionBlock := container.Functions[AwooScopeGlobalFunctionId].Blocks[AwooScopeGlobalBlockId]
	scopeSymbol := AwooScopeSymbolTableEntry{
		Symbol: symbol,
		Global: true,
	}
	if _, ok := functionBlock.SymbolTable.External[symbol.Name]; ok {
		return scopeSymbol, fmt.Errorf("%w: %s", awerrors.ErrorAlreadyDefinedVariable, gchalk.Red(symbol.Name))
	}

	functionBlock.SymbolTable.External[symbol.Name] = scopeSymbol
	container.Functions[AwooScopeGlobalFunctionId].Blocks[AwooScopeGlobalBlockId] = functionBlock

	return scopeSymbol, nil
}

func PushCurrentFunctionSymbol(container *AwooScopeContainer, blockEntry elf.AwooElfSymbolTableEntry) (AwooScopeSymbolTableEntry, error) {
	scopeFunction, ok := container.Functions[container.NextFunctionId-1]
	if !ok {
		panic("not possible")
	}

	return PushFunctionBlockSymbol(container, scopeFunction.Id, scopeFunction.NextBlockId-1, blockEntry)
}

func PopFunctionBlockSymbol(container *AwooScopeContainer, funcId uint16, blockId uint16, name string) error {
	functionBlock := container.Functions[funcId].Blocks[blockId]
	if _, ok := functionBlock.SymbolTable.Internal[name]; !ok {
		return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}
	delete(functionBlock.SymbolTable.Internal, name)
	container.Functions[funcId].Blocks[blockId] = functionBlock

	return nil
}

func PopCurrentFunctionSymbol(container *AwooScopeContainer, name string) error {
	for funcId := int16(container.NextFunctionId) - 1; funcId >= 0; funcId-- {
		for blockId := int16(container.Functions[uint16(funcId)].NextBlockId) - 1; blockId >= 0; blockId-- {
			if err := PopFunctionBlockSymbol(container, uint16(funcId), uint16(blockId), name); err == nil {
				return nil
			}
		}
	}

	return fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetFunctionBlockSymbol(container *AwooScopeContainer, funcId uint16, blockId uint16, name string) (AwooScopeSymbolTableEntry, error) {
	if blockEntry, ok := container.Functions[funcId].Blocks[blockId].SymbolTable.Internal[name]; ok {
		return blockEntry, nil
	}
	if blockEntry, ok := container.Functions[funcId].Blocks[blockId].SymbolTable.External[name]; ok {
		return blockEntry, nil
	}

	return AwooScopeSymbolTableEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}

func GetCurrentFunctionSymbol(container *AwooScopeContainer, name string) (AwooScopeSymbolTableEntry, error) {
	for funcId := int16(container.NextFunctionId) - 1; funcId >= 0; funcId-- {
		for blockId := int16(container.Functions[uint16(funcId)].NextBlockId) - 1; blockId >= 0; blockId-- {
			if blockEntry, err := GetFunctionBlockSymbol(container, uint16(funcId), uint16(blockId), name); err == nil {
				return blockEntry, nil
			}
		}
	}

	return AwooScopeSymbolTableEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
}
