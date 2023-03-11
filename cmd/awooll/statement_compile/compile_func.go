package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementFunc(ccompiler *compiler.AwooCompiler, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	functionNameNode := statement.GetStatementFuncIdentifier(&s)
	functionName := node.GetNodeIdentifierValue(&functionNameNode)
	compiler_context.PushCompilerScopeFunction(&ccompiler.Context, compiler_context.AwooCompilerScopeFunction{
		Name: functionName,
	})

	functionArguments := statement.GetStatementFuncArguments(&s)
	functionArgumentsOffset := uint32(0)
	for _, argument := range functionArguments {
		_, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerMemoryEntry{
			Name: argument.Name,
			Size: argument.Size,
			Type: argument.Type,
			Data: argument.Data,
		})
		if err != nil {
			return d, err
		}
		functionArgumentsOffset += uint32(argument.Size)
	}

	_, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerMemoryEntry{
		Name: "_returnAddress",
		Size: 4,
		Type: types.AwooTypePointer,
	})
	if err != nil {
		return d, err
	}

	functionReturnTypeNode := statement.GetStatementFuncReturnType(&s)
	var functionReturnType *uint16
	if functionReturnTypeNode != nil {
		returnType := node.GetNodeTypeType(functionReturnTypeNode)
		functionReturnType = &returnType
	}

	d, err = encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		SourceTwo:   cpu.AwooRegisterReturnAddress,
		Immediate:   functionArgumentsOffset,
	}, d)
	if err != nil {
		return d, err
	}

	compiler_context.PushCompilerFunction(&ccompiler.Context, compiler_context.AwooCompilerFunction{
		Name:       functionName,
		ReturnType: functionReturnType,
		Arguments:  statement.GetStatementFuncArguments(&s),
		Start:      compiler_context.GetProgramHeaderSize() + ccompiler.Context.CurrentAddress,
		Size:       uint16(len(d)),
	})
	d, err = CompileStatementGroup(ccompiler, statement.GetStatementFuncBody(&s), d)
	if err != nil {
		return d, err
	}

	compiler_context.PopCompilerScopeCurrentFunction(&ccompiler.Context)
	if ccompiler.Context.Functions.Start == "" {
		ccompiler.Context.Functions.Start = functionName
		d = append(make([]byte, compiler_context.GetProgramHeaderSize()), d...)
	}

	return d, nil
}
