package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/node"
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/LamkasDev/awoo-emu/cmd/common/instructions"
	"github.com/jwalton/gchalk"
)

func HandleNodeExpressionLeftRight(ins instruction.AwooInstructionDefinition) compiler.AwooCompileNodeExpression {
	return func(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
		return encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: ins,
			SourceOne:   leftDetails.Register,
			SourceTwo:   rightDetails.Register,
			Destination: leftDetails.Register,
		}, d)
	}
}

func HandleNodeExpressionRightLeft(ins instruction.AwooInstructionDefinition) compiler.AwooCompileNodeExpression {
	return func(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
		return encoder.Encode(encoder.AwooEncodedInstruction{
			Instruction: ins,
			SourceOne:   rightDetails.Register,
			SourceTwo:   leftDetails.Register,
			Destination: leftDetails.Register,
		}, d)
	}
}

func CompileNodeExpressionEqEq(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := HandleNodeExpressionLeftRight(instructions.AwooInstructionSUB)(ccompiler, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionSLTIU,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpressionNotEq(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := HandleNodeExpressionLeftRight(instructions.AwooInstructionSUB)(ccompiler, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionSLTU,
		SourceOne:   cpu.AwooRegisterZero,
		SourceTwo:   leftDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionLTEQ(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := HandleNodeExpressionRightLeft(instructions.AwooInstructionSLT)(ccompiler, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionXORI,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpressionGTEQ(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := HandleNodeExpressionLeftRight(instructions.AwooInstructionSLT)(ccompiler, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instructions.AwooInstructionXORI,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpression(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	entry, ok := ccompiler.Settings.Mappings.NodeExpression[n.Token.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileOperator, gchalk.Red(ccompiler.Settings.Parser.Lexer.Tokens.All[n.Token.Type].Name))
	}

	var err error
	left := node.GetNodeExpressionLeft(&n)
	leftDetails := compiler_details.CompileNodeValueDetails{
		Type:     details.Type,
		Register: details.Register,
	}
	if d, err = CompileNodeValue(ccompiler, left, d, &leftDetails); err != nil {
		return d, err
	}
	right := node.GetNodeExpressionRight(&n)
	rightDetails := compiler_details.CompileNodeValueDetails{
		Type:     details.Type,
		Register: cpu.GetNextTemporaryRegister(details.Register),
	}
	if d, err = CompileNodeValue(ccompiler, right, d, &rightDetails); err != nil {
		return d, err
	}

	return entry(ccompiler, d, &leftDetails, &rightDetails)
}
