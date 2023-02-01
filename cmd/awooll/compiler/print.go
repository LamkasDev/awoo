package compiler

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintNewCompile(context *lexer_context.AwooLexerContext, s *statement.AwooParserStatement, data []byte) {
	text := ""
	for _, b := range data {
		text += fmt.Sprintf("%#x ", b)
	}

	logger.Log("%s %s  %s\n",
		statement.PrintStatement(context, s),
		gchalk.Gray("➔"),
		gchalk.Cyan(text),
	)
}
