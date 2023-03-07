package statement_compile

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_details"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/encoder"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/instruction"
	"github.com/jwalton/gchalk"
)

func CompileNodeExpressionOp(ccompiler *compiler.AwooCompiler, ins instruction.AwooInstruction, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: ins,
		SourceOne:   leftDetails.Register,
		SourceTwo:   rightDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionAdd(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(ccompiler, instruction.AwooInstructionADD, d, leftDetails, rightDetails)
}

func CompileNodeExpressionSubstract(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(ccompiler, instruction.AwooInstructionSUB, d, leftDetails, rightDetails)
}

func CompileNodeExpressionMultiply(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(ccompiler, instruction.AwooInstructionMUL, d, leftDetails, rightDetails)
}

func CompileNodeExpressionDivide(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	return CompileNodeExpressionOp(ccompiler, instruction.AwooInstructionDIV, d, leftDetails, rightDetails)
}

func CompileNodeExpressionEqEq(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionOp(ccompiler, instruction.AwooInstructionSUB, d, leftDetails, rightDetails)
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

func CompileNodeExpressionNotEq(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionOp(ccompiler, instruction.AwooInstructionSUB, d, leftDetails, rightDetails)
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

func CompileNodeExpressionLT(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLT,
		SourceOne:   leftDetails.Register,
		SourceTwo:   rightDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionLTEQ(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionGT(ccompiler, d, leftDetails, rightDetails)
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

func CompileNodeExpressionGT(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	return encoder.Encode(encoder.AwooEncodedInstruction{
		Instruction: instruction.AwooInstructionSLT,
		SourceOne:   rightDetails.Register,
		SourceTwo:   leftDetails.Register,
		Destination: leftDetails.Register,
	}, d)
}

func CompileNodeExpressionGTEQ(ccompiler *compiler.AwooCompiler, d []byte, leftDetails *compiler_details.CompileNodeValueDetails, rightDetails *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	d, err := CompileNodeExpressionLT(ccompiler, d, leftDetails, rightDetails)
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

func CompileNodeExpression(ccompiler *compiler.AwooCompiler, n node.AwooParserNode, d []byte, details *compiler_details.CompileNodeValueDetails) ([]byte, error) {
	entry, ok := ccompiler.Settings.Mappings.NodeExpression[n.Token.Type]
	if !ok {
		return d, fmt.Errorf("%w: %s", awerrors.ErrorCantCompileOperator, gchalk.Red(ccompiler.Settings.Parser.Lexer.Tokens.All[n.Token.Type].Name))
	}
	left := node.GetNodeExpressionLeft(&n)
	leftDetails := compiler_details.CompileNodeValueDetails{Register: details.Register}
	rightDetails := compiler_details.CompileNodeValueDetails{Register: cpu.GetNextTemporaryRegister(details.Register)}
	d, err := CompileNodeValue(ccompiler, left, d, &leftDetails)
	if err != nil {
		return d, err
	}
	right := node.GetNodeExpressionRight(&n)
	d, err = CompileNodeValue(ccompiler, right, d, &rightDetails)
	if err != nil {
		return d, err
	}

	return entry(ccompiler, d, &leftDetails, &rightDetails)
}
