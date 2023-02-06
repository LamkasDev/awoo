package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

type AwooCompileNodeExpression func(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error)

var AwooCompileNodeExpressions = map[uint16]AwooCompileNodeExpression{
	token.TokenOperatorAddition:       CompileNodeExpressionAdd,
	token.TokenOperatorSubstraction:   CompileNodeExpressionSubstract,
	token.TokenOperatorMultiplication: CompileNodeExpressionMultiply,
	token.TokenOperatorDivision:       CompileNodeExpressionDivide,
	token.TokenOperatorEqEq:           CompileNodeExpressionEqEq,
	token.TokenOperatorNotEq:          CompileNodeExpressionNotEq,
	token.TokenOperatorLT:             CompileNodeExpressionLT,
	token.TokenOperatorLTEQ:           CompileNodeExpressionLTEQ,
	token.TokenOperatorGT:             CompileNodeExpressionGT,
	token.TokenOperatorGTEQ:           CompileNodeExpressionGTEQ,
}

func CompileNodeExpressionOp(context *compiler_context.AwooCompilerContext, ins instruction.AwooInstruction, r uint8, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: ins,
		SourceOne:   details.Register,
		SourceTwo:   r,
		Destination: details.Register,
	}, d)
}

func CompileNodeExpressionAdd(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionADD, cpu.GetNextTemporaryRegister(details.Register), d, details)
}

func CompileNodeExpressionSubstract(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionSUB, cpu.GetNextTemporaryRegister(details.Register), d, details)
}

func CompileNodeExpressionMultiply(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionMUL, cpu.GetNextTemporaryRegister(details.Register), d, details)
}

func CompileNodeExpressionDivide(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(context, instruction.AwooInstructionDIV, cpu.GetNextTemporaryRegister(details.Register), d, details)
}

func CompileNodeExpressionEqEq(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	r := cpu.GetNextTemporaryRegister(details.Register)
	d, err := CompileNodeExpressionOp(context, instruction.AwooInstructionSUB, r, d, details)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLTIU,
		SourceOne:   details.Register,
		Destination: details.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpressionNotEq(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	r := cpu.GetNextTemporaryRegister(details.Register)
	d, err := CompileNodeExpressionOp(context, instruction.AwooInstructionSUB, r, d, details)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLTU,
		SourceOne:   cpu.AwooRegisterZero,
		SourceTwo:   details.Register,
		Destination: details.Register,
	}, d)
}

func CompileNodeExpressionLT(_ *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	r := cpu.GetNextTemporaryRegister(details.Register)
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLT,
		SourceOne:   details.Register,
		SourceTwo:   r,
		Destination: details.Register,
	}, d)
}

func CompileNodeExpressionLTEQ(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionGT(context, d, details)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionXORI,
		SourceOne:   details.Register,
		Destination: details.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpressionGT(_ *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	r := cpu.GetNextTemporaryRegister(details.Register)
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLT,
		SourceOne:   r,
		SourceTwo:   details.Register,
		Destination: details.Register,
	}, d)
}

func CompileNodeExpressionGTEQ(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionLT(context, d, details)
	if err != nil {
		return d, err
	}
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionXORI,
		SourceOne:   details.Register,
		Destination: details.Register,
		Immediate:   1,
	}, d)
}

func CompileNodeExpression(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	entry, ok := AwooCompileNodeExpressions[n.Token.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileOperator, gchalk.Red(context.Parser.Lexer.Tokens.All[n.Token.Type].Name))
	}
	left := node.GetNodeExpressionLeft(&n)
	d, err := CompileNodeValue(context, left, d, CompileNodeValueDetails{Register: details.Register})
	if err != nil {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorFailedToCompileOperator, err)
	}
	right := node.GetNodeExpressionRight(&n)
	d, err = CompileNodeValue(context, right, d, CompileNodeValueDetails{Register: cpu.GetNextTemporaryRegister(details.Register)})
	if err != nil {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorFailedToCompileOperator, err)
	}
	d, err = entry(context, d, details)
	if err != nil {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorFailedToCompileOperator, err)
	}

	return d, nil
}
