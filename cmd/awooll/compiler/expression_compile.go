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

func CompileNodeExpressionAdd(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionADD,
		SourceOne:   details.Register,
		SourceTwo:   cpu.GetNextTemporaryRegister(details.Register),
		Destination: details.Register,
	}, d)
}

func CompileNodeExpressionSubstract(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSUB,
		SourceOne:   details.Register,
		SourceTwo:   cpu.GetNextTemporaryRegister(details.Register),
		Destination: details.Register,
	}, d)
}

func CompileNodeExpressionMultiply(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionMUL,
		SourceOne:   details.Register,
		SourceTwo:   cpu.GetNextTemporaryRegister(details.Register),
		Destination: details.Register,
	}, d)
}

func CompileNodeExpressionDivide(context *compiler_context.AwooCompilerContext, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionDIV,
		SourceOne:   details.Register,
		SourceTwo:   cpu.GetNextTemporaryRegister(details.Register),
		Destination: details.Register,
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
		// TODO: figure out which side
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
		}
	}

	return d, fmt.Errorf("no idea how to compile expression with operator %s", gchalk.Red(context.Parser.Lexer.Tokens.All[op].Name))
}
