package statement_compile

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
)

func CompileStatementDefinition(context *compiler_context.AwooCompilerContext, s statement.AwooParserStatement, d []byte) ([]byte, error) {
	tNode := statement.GetStatementDefinitionVariableType(&s)
	entry := compiler_context.AwooCompilerContextMemoryEntry{}
	switch tNode.Type {
	case node.ParserNodeTypeType:
		entry.Type = node.GetNodeTypeType(&tNode)
	case node.ParserNodeTypePointer:
		entry.Type = types.AwooTypePointer
		// TODO: chaining pointers
		tNode = node.GetNodeSingleValue(&tNode)
		entry.Data = node.GetNodeTypeType(&tNode)
	}
	entry.Size = context.Parser.Lexer.Types.All[entry.Type].Size

	nameNode := statement.GetStatementDefinitionVariableIdentifier(&s)
	entry.Name = node.GetNodeIdentifierValue(&nameNode)
	valueNode := statement.GetStatementDefinitionVariableValue(&s)
	dest, err := compiler_context.PushCompilerScopeCurrentBlockMemory(context, entry)
	if err != nil {
		return d, err
	}
	details := compiler_context.CompileNodeValueDetails{Register: cpu.AwooRegisterTemporaryZero}
	d, err = CompileNodeValueFast(context, valueNode, d, &details)
	if err != nil {
		return d, err
	}

	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSW,
		SourceOne:   cpu.AwooRegisterSavedZero,
		SourceTwo:   details.Register,
		Immediate:   uint32(dest),
	}, d)
}
