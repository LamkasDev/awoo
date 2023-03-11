package lexer

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

func PrintNewToken(settings *AwooLexerSettings, cs string, t *lexer_token.AwooLexerToken) {
	logger.Log("%s %s  %s\n", cs, gchalk.Gray("➔"), PrintToken(settings, t))
}

func PrintToken(settings *AwooLexerSettings, t *lexer_token.AwooLexerToken) string {
	details := gchalk.Green(fmt.Sprintf("%#x", t.Type))
	if t.Data != nil {
		details += fmt.Sprintf(", %v", gchalk.Blue(fmt.Sprint(t.Data)))
	}

	return fmt.Sprintf("%s %s (%s)", token.GetTokenTypeName(t.Type), settings.Tokens.All[t.Type].Name, details)
}

func PrintTokenTypes(settings *AwooLexerSettings, tokenTypes []uint16) string {
	text := settings.Tokens.All[tokenTypes[0]].Key
	for i := 1; i < len(tokenTypes)-1; i++ {
		text = fmt.Sprintf("%s, %s", text, settings.Tokens.All[tokenTypes[i]].Key)
	}
	return fmt.Sprintf("%s or %s", text, settings.Tokens.All[tokenTypes[len(tokenTypes)-1]].Key)
}
