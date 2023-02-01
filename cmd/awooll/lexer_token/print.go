package lexer_token

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintNewToken(context *lexer_context.AwooLexerContext, cs string, t *AwooLexerToken) {
	logger.Log("%s %s  %s\n", cs, gchalk.Gray("➔"), PrintToken(context, t))
}

func PrintToken(context *lexer_context.AwooLexerContext, t *AwooLexerToken) string {
	tokenType := "token"
	if token.IsTokenTypeGeneral(t.Type) {
		tokenType = "token"
	} else if token.IsTokenTypeOperator(t.Type) {
		tokenType = "op"
	} else if token.IsTokenTypeKeyword(t.Type) {
		tokenType = "type"
	} else if token.IsTokenTypeKeyword(t.Type) {
		tokenType = "keyword"
	}
	tokenName := context.Tokens.All[t.Type].Name
	details := gchalk.Green(fmt.Sprintf("%#x", t.Type))
	if t.Data != nil {
		details += fmt.Sprintf(", %v", gchalk.Blue(fmt.Sprint(t.Data)))
	}

	return fmt.Sprintf("%s %s (%s)", tokenType, tokenName, details)
}
