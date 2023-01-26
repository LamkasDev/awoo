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

func CompileNodeExpressionAdd(context *compiler_context.AwooCompilerContext, d []byte) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADD,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		SourceTwo:   cpu.AwooRegisterTemporaryOne,
		Destination: cpu.AwooRegisterTemporaryZero,
	}, d)
}

func CompileNodeExpressionSubstract(context *compiler_context.AwooCompilerContext, d []byte) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSUB,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		SourceTwo:   cpu.AwooRegisterTemporaryOne,
		Destination: cpu.AwooRegisterTemporaryZero,
	}, d)
}

func CompileNodeExpressionMultiply(context *compiler_context.AwooCompilerContext, d []byte) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionMUL,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		SourceTwo:   cpu.AwooRegisterTemporaryOne,
		Destination: cpu.AwooRegisterTemporaryZero,
	}, d)
}

func CompileNodeExpressionDivide(context *compiler_context.AwooCompilerContext, d []byte) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionDIV,
		SourceOne:   cpu.AwooRegisterTemporaryZero,
		SourceTwo:   cpu.AwooRegisterTemporaryOne,
		Destination: cpu.AwooRegisterTemporaryZero,
	}, d)
}

func CompileNodeExpression(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	left := node.GetNodeExpressionLeft(&n)
	op := n.Token.Type
	right := node.GetNodeExpressionRight(&n)

	// TODO: make a chain of operation, work correctly
	switch op {
	case token.TokenOperatorAddition,
		token.TokenOperatorSubstraction,
		token.TokenOperatorMultiplication,
		token.TokenOperatorDivision:
		d, err := CompileNodeValue(context, left, d, CompileNodeValueDetails{Expression: false})
		if err != nil {
			return d, err
		}
		d, err = CompileNodeValue(context, right, d, CompileNodeValueDetails{Expression: true})
		if err != nil {
			return d, err
		}
		switch op {
		case token.TokenOperatorAddition:
			return CompileNodeExpressionAdd(context, d)
		case token.TokenOperatorSubstraction:
			return CompileNodeExpressionSubstract(context, d)
		case token.TokenOperatorMultiplication:
			return CompileNodeExpressionMultiply(context, d)
		case token.TokenOperatorDivision:
			return CompileNodeExpressionDivide(context, d)
		}
	}

	return d, fmt.Errorf("no idea how to compile expression with operator %s", gchalk.Red(context.Parser.Lexer.Tokens.All[op].Name))
}
