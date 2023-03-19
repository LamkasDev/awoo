package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
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
			Name:        argument.Name,
			Size:        argument.Size,
			Type:        argument.Type,
			TypeDetails: argument.TypeDetails,
		})
		if err != nil {
			return d, err
		}
		functionArgumentsOffset += uint32(argument.Size)
	}

	_, err := compiler_context.PushCompilerScopeCurrentBlockMemory(&ccompiler.Context, compiler_context.AwooCompilerMemoryEntry{
		Name: "_returnAddress",
		Size: 4,
		Type: commonTypes.AwooTypeId(types.AwooTypePointer),
	})
	if err != nil {
		return d, err
	}

	functionReturnTypeNode := statement.GetStatementFuncReturnType(&s)
	var functionReturnType *commonTypes.AwooTypeId
	if functionReturnTypeNode != nil {
		returnType := node.GetNodeTypeType(functionReturnTypeNode)
		functionReturnType = &returnType
	}

	stackAdjustmentInstruction := encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionSW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		SourceTwo:   cpu.AwooRegisterReturnAddress,
		Immediate:   functionArgumentsOffset,
	}
	if d, err = encoder.Encode(stackAdjustmentInstruction, d); err != nil {
		return d, err
	}

	compiler_context.PushCompilerFunction(&ccompiler.Context, compiler_context.AwooCompilerFunction{
		Name:       functionName,
		ReturnType: functionReturnType,
		Arguments:  statement.GetStatementFuncArguments(&s),
		Start:      compiler_context.GetProgramHeaderSize() + ccompiler.Context.CurrentAddress,
		Size:       uint16(len(d)),
	})
	if d, err = CompileStatementGroup(ccompiler, statement.GetStatementFuncBody(&s), d); err != nil {
		return d, err
	}
	compiler_context.PopCompilerScopeCurrentFunction(&ccompiler.Context)

	if ccompiler.Context.Functions.Start == "" {
		ccompiler.Context.Functions.Start = functionName
		d = append(make([]byte, compiler_context.GetProgramHeaderSize()), d...)
	}

	return d, nil
}
