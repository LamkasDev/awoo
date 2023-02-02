package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

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

func CompileNodeExpression(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	left := node.GetNodeExpressionLeft(&n)
	op := n.Token.Type
	right := node.GetNodeExpressionRight(&n)

	switch op {
	case token.TokenOperatorAddition,
		token.TokenOperatorSubstraction,
		token.TokenOperatorMultiplication,
		token.TokenOperatorDivision,
		token.TokenOperatorEqEq,
		token.TokenOperatorNotEq:
		d, err := CompileNodeValue(context, left, d, CompileNodeValueDetails{Register: details.Register})
		if err != nil {
			return d, err
		}
		d, err = CompileNodeValue(context, right, d, CompileNodeValueDetails{Register: cpu.GetNextTemporaryRegister(details.Register)})
		if err != nil {
			return d, err
		}
		switch op {
		case token.TokenOperatorAddition:
			return CompileNodeExpressionAdd(context, d, CompileNodeValueDetails{Register: details.Register})
		case token.TokenOperatorSubstraction:
			return CompileNodeExpressionSubstract(context, d, CompileNodeValueDetails{Register: details.Register})
		case token.TokenOperatorMultiplication:
			return CompileNodeExpressionMultiply(context, d, CompileNodeValueDetails{Register: details.Register})
		case token.TokenOperatorDivision:
			return CompileNodeExpressionDivide(context, d, CompileNodeValueDetails{Register: details.Register})
		case token.TokenOperatorEqEq:
			return CompileNodeExpressionEqEq(context, d, CompileNodeValueDetails{Register: details.Register})
		case token.TokenOperatorNotEq:
			return CompileNodeExpressionNotEq(context, d, CompileNodeValueDetails{Register: details.Register})
		}
	}

	return d, fmt.Errorf("no idea how to compile expression with operator %s", gchalk.Red(context.Parser.Lexer.Tokens.All[op].Name))
}
