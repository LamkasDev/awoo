package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func CompileNodeExpressionOp(context *compiler_context.AwooCompilerContext, ins instruction.AwooInstruction, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: ins,
		SourceOne:   leftDetails.Register,
		SourceTwo:   rightDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionAdd(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionADD, d, leftDetails, rightDetails)
}

func CompileNodeExpressionSubstract(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionSUB, d, leftDetails, rightDetails)
}

func CompileNodeExpressionMultiply(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionMUL, d, leftDetails, rightDetails)
}

func CompileNodeExpressionDivide(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionDIV, d, leftDetails, rightDetails)
}

func CompileNodeExpressionEqEq(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionOp(context, instruction.AwooInstructionSUB, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLTIU,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpressionNotEq(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionOp(context, instruction.AwooInstructionSUB, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLTU,
		SourceOne:   cpu.AwooRegisterZero,
		SourceTwo:   leftDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionLT(_ *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLT,
		SourceOne:   leftDetails.Register,
		SourceTwo:   rightDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionLTEQ(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionGT(context, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionXORI,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpressionGT(_ *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLT,
		SourceOne:   rightDetails.Register,
		SourceTwo:   leftDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionGTEQ(context *compiler_context.AwooCompilerContext, d []byte, leftDetails *compiler_context.CompileNodeValueDetails, rightDetails *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionLT(context, d, leftDetails, rightDetails)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionXORI,
		SourceOne:   leftDetails.Register,
		Destination: leftDetails.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpression(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details *compiler_context.CompileNodeValueDetails) ([]byte, error) {
	entry, ok := context.MappingsNodeExpression[n.Token.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileOperator, gchalk.Red(context.Parser.Lexer.Tokens.All[n.Token.Type].Name))
	}
	left := node.GetNodeExpressionLeft(&n)
	leftDetails := compiler_context.CompileNodeValueDetails{Register: details.Register}
	rightDetails := compiler_context.CompileNodeValueDetails{Register: cpu.GetNextTemporaryRegister(details.Register)}
	d, err := CompileNodeValue(context, left, d, &leftDetails)
	if err != nil {
		return d, err
	}
	right := node.GetNodeExpressionRight(&n)
	d, err = CompileNodeValue(context, right, d, &rightDetails)
	if err != nil {
		return d, err
	}
	d, err = entry(context, d, &leftDetails, &rightDetails)
	if err != nil {
		return d, err
	}

	return d, nil
}
