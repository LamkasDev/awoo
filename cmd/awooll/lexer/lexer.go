package lexer

import (
	"fmt"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

type AwooLexer struct {
	Contents []rune
	Length   uint16
	Position uint16
	Current  rune
	Context  lexer_context.AwooLexerContext
	Settings AwooLexerSettings
}

type AwooLexerSettings struct{}

func SetupLexer(settings AwooLexerSettings) AwooLexer {
	lexer := AwooLexer{
		Context: lexer_context.AwooLexerContext{
			Tokens: token.SetupTokenMap(),
			Types:  types.SetupTypeMap(),
		},
		Settings: settings,
	}

	return lexer
}

func LoadLexer(lexer *AwooLexer, contents []rune) {
	lexer.Contents = contents
	lexer.Length = (uint16)(len(contents))
	lexer.Position = 0
	lexer.Current = lexer.Contents[lexer.Position]
}

func AdvanceLexerFor(lexer *AwooLexer, n int16) (rune, bool) {
	lexer.Position = (uint16)((int16)(lexer.Position) + n)
	if lexer.Position >= lexer.Length {
		return 0, false
	}
	lexer.Current = lexer.Contents[lexer.Position]
	return lexer.Current, true
}

func AdvanceLexer(lexer *AwooLexer) (rune, bool) {
	return AdvanceLexerFor(lexer, 1)
}

func PeekLexer(lexer *AwooLexer) (rune, bool) {
	if lexer.Position+1 >= lexer.Length {
		return 0, false
	}
	return lexer.Contents[lexer.Position+1], true
}

func StepbackLexer(lexer *AwooLexer) (rune, bool) {
	return AdvanceLexerFor(lexer, -1)
}

func RunLexer(lexer *AwooLexer) AwooLexerResult {
	result := AwooLexerResult{
		Context: lexer.Context,
	}
	logger.Log(gchalk.Yellow("> Lexer\n"))
	logger.Log("Input: %s\n", gchalk.Magenta(string(lexer.Contents)))
	for ok := true; ok; _, ok = AdvanceLexer(lexer) {
		if unicode.IsSpace(lexer.Current) {
			continue
		}

		single, ok := lexer.Context.Tokens.Single[lexer.Current]
		if ok {
			token := lexer_token.CreateToken(lexer.Position, single)
			lexer_token.PrintNewToken(&lexer.Context, string(lexer.Current), &token)
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if unicode.IsLetter(lexer.Current) {
			token, matchedString := CreateTokenLetter(lexer)
			lexer_token.PrintNewToken(&lexer.Context, matchedString, &token)
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if unicode.IsNumber(lexer.Current) {
			token, matchedString, err := CreateTokenNumber(lexer)
			if err != nil {
				result.Error = err
				break
			}
			lexer_token.PrintNewToken(&lexer.Context, matchedString, &token)
			result.Tokens = append(result.Tokens, token)
			continue
		}

		token, matchedString, ok := CreateTokenCouple(lexer)
		if !ok {
			result.Error = fmt.Errorf("%w: %s", awerrors.ErrorIllegalCharacter, gchalk.Red((string)(lexer.Current)))
			break
		}
		lexer_token.PrintNewToken(&lexer.Context, matchedString, &token)
		result.Tokens = append(result.Tokens, token)
		break
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
