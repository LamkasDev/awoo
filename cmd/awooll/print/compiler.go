package print

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

func PrintNewCompile(context *lexer_context.AwooLexerContext, s *statement.AwooParserStatement, data []byte) {
	text := ""
	for _, b := range data {
		text += fmt.Sprintf("%#x ", b)
	}

	fmt.Printf("%s %s  %s\n",
		PrintStatement(context, s),
		gchalk.Gray("➔"),
		gchalk.Cyan(text),
	)
}
