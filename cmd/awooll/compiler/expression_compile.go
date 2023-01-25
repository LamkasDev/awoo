package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

func CompileNodeExpression(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	left := node.GetNodeExpressionLeft(&n)
	op := n.Token.Type
	right := node.GetNodeExpressionRight(&n)

	// TODO: make a chain of operation, work correctly
	switch op {
	case token.TokenOperatorAddition:
		d, err := CompileNodeValue(context, left, d, CompileNodeValueDetails{Expression: false})
		if err != nil {
			return d, err
		}
		d, err = CompileNodeValue(context, right, d, CompileNodeValueDetails{Expression: true})
		if err != nil {
			return d, err
		}
		return CompileNodeValueAdd(context, d, details)
	case token.TokenOperatorSubstraction:
		d, err := CompileNodeValue(context, left, d, CompileNodeValueDetails{Expression: false})
		if err != nil {
			return d, err
		}
		d, err = CompileNodeValue(context, right, d, CompileNodeValueDetails{Expression: true})
		return CompileNodeValueSubstract(context, d, details)
	}

	return d, fmt.Errorf("no idea how to compile expression with operator %s", gchalk.Red(context.Parser.Lexer.Tokens.All[op].Name))
}
