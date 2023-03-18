package lexer

import (
	"fmt"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
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

type AwooLexerSettings struct {
	Tokens   token.AwooTokenMap
	Mappings AwooLexerMappings
}

func SetupLexer(settings AwooLexerSettings) AwooLexer {
	lexer := AwooLexer{
		Context: lexer_context.AwooLexerContext{
			Types: types.SetupTypeMap(),
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

func AdvanceLexerFor(lexer *AwooLexer, n int16) (rune, error) {
	lexer.Position = (uint16)((int16)(lexer.Position) + n)
	if lexer.Position >= lexer.Length {
		return 0, awerrors.ErrorNoMoreTokens
	}
	lexer.Current = lexer.Contents[lexer.Position]
	return lexer.Current, nil
}

func AdvanceLexer(lexer *AwooLexer) (rune, error) {
	return AdvanceLexerFor(lexer, 1)
}

func PeekLexer(lexer *AwooLexer) (rune, error) {
	if lexer.Position+1 >= lexer.Length {
		return 0, awerrors.ErrorNoMoreTokens
	}
	return lexer.Contents[lexer.Position+1], nil
}

func StepbackLexer(lexer *AwooLexer) (rune, error) {
	return AdvanceLexerFor(lexer, -1)
}

func RunLexer(lexer *AwooLexer) AwooLexerResult {
	result := AwooLexerResult{
		Context: lexer.Context,
	}
	logger.Log(gchalk.Yellow("> Lexer\n"))
	var err error
	for ; err == nil; _, err = AdvanceLexer(lexer) {
		if unicode.IsSpace(lexer.Current) {
			continue
		}

		single, ok := lexer.Settings.Tokens.Single[lexer.Current]
		if ok {
			token := lexer_token.CreateToken(lexer.Position, single)
			PrintNewToken(&lexer.Settings, string(lexer.Current), &token)
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if lexer.Current == '\'' {
			token, matchedString, err := CreateTokenChar(lexer)
			if err != nil {
				result.Error = err
				break
			}
			PrintNewToken(&lexer.Settings, matchedString, &token)
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if unicode.IsLetter(lexer.Current) {
			token, matchedString := CreateTokenLetter(lexer)
			PrintNewToken(&lexer.Settings, matchedString, &token)
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if unicode.IsNumber(lexer.Current) {
			token, matchedString, err := CreateTokenNumber(lexer)
			if err != nil {
				result.Error = err
				break
			}
			PrintNewToken(&lexer.Settings, matchedString, &token)
			result.Tokens = append(result.Tokens, token)
			continue
		}

		token, matchedString, ok := CreateTokenCouple(lexer)
		if !ok {
			result.Error = fmt.Errorf("%w: %s", awerrors.ErrorIllegalCharacter, gchalk.Red((string)(lexer.Current)))
			break
		}
		PrintNewToken(&lexer.Settings, matchedString, &token)
		result.Tokens = append(result.Tokens, token)
		break
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
