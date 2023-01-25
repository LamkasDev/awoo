package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/compiler_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/node"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

type CompileNodeValueDetails struct {
	First bool
}

func CompileNodeValue(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	switch n.Type {
	case node.ParserNodeTypeIdentifier:
		return CompileNodeIdentifier(context, n, d, details)
	case node.ParserNodeTypePrimitive:
		return CompileNodePrimitive(context, n, d, details)
	case node.ParserNodeTypeExpression:
		return CompileNodeExpression(context, n, d, details)
	}

	return d, fmt.Errorf("no idea how to compile value node %s", gchalk.Red(fmt.Sprintf("%#x", n.Type)))
}

func CompileNodeValueFast(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte) ([]byte, error) {
	return CompileNodeValue(context, n, d, CompileNodeValueDetails{First: true})
}

func CompileNodeExpression(context *compiler_context.AwooCompilerContext, n node.AwooParserNode, d []byte, details CompileNodeValueDetails) ([]byte, error) {
	left := node.GetNodeExpressionLeft(&n)
	op := n.Token.Type
	right := node.GetNodeExpressionRight(&n)

	switch op {
	case token.TokenOperatorAddition:
		d, err := CompileNodeValue(context, left, d, details)
		if details.First {
			details.First = false
		}
		if err != nil {
			return d, err
		}
		return CompileNodeValue(context, right, d, details)
	}

	return d, fmt.Errorf("no idea how to compile expression with operator %s", gchalk.Red(context.Parser.Lexer.Tokens.All[op].Name))
}
