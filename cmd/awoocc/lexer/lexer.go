package lexer

import (
	"fmt"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	"github.com/LamkasDev/awoo-emu/cmd/common/cc"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/jwalton/gchalk"
)

type AwooLexer struct {
	Contents AwooLexerContents
	Current  AwooLexerCurrent
	Context  lexer_context.AwooLexerContext
	Settings AwooLexerSettings
}

type AwooLexerContents struct {
	Text   []rune
	Length uint32
}

type AwooLexerCurrent struct {
	Position  lexer_token.AwooLexerTokenPosition
	Character rune
}

type AwooLexerSettings struct {
	Path     string
	Tokens   token.AwooTokenMap
	Mappings AwooLexerMappings
}

func NewLexer(settings AwooLexerSettings) AwooLexer {
	lexer := AwooLexer{
		Context: lexer_context.AwooLexerContext{
			Types: types.SetupTypeMap(),
		},
		Settings: settings,
	}

	return lexer
}

func NewAwooLexerTokenPosition(lexer *AwooLexer, length uint32) lexer_token.AwooLexerTokenPosition {
	return lexer_token.AwooLexerTokenPosition{
		Line:   lexer.Current.Position.Line,
		Column: lexer.Current.Position.Column,
		Length: length,
	}
}

func LoadLexer(lexer *AwooLexer, contents []rune) {
	lexer.Contents = AwooLexerContents{
		Text:   contents,
		Length: (uint32)(len(contents)),
	}
	lexer.Current = AwooLexerCurrent{
		Position: lexer_token.AwooLexerTokenPosition{
			Line:   1,
			Column: 1,
		},
		Character: lexer.Contents.Text[lexer.Current.Position.Index],
	}
}

func AdvanceLexerFor(lexer *AwooLexer, n int32) (rune, error) {
	lexer.Current.Position.Index = (uint32)((int32)(lexer.Current.Position.Index) + n)
	if lexer.Current.Position.Index >= lexer.Contents.Length {
		return 0, awerrors.ErrorNoMoreTokens
	}
	lexer.Current.Position.Column = (uint32)((int32)(lexer.Current.Position.Column) + n)
	lexer.Current.Character = lexer.Contents.Text[lexer.Current.Position.Index]
	return lexer.Current.Character, nil
}

func AdvanceLexer(lexer *AwooLexer) (rune, error) {
	return AdvanceLexerFor(lexer, 1)
}

func PeekLexer(lexer *AwooLexer) (rune, error) {
	if lexer.Current.Position.Index+1 >= lexer.Contents.Length {
		return 0, awerrors.ErrorNoMoreTokens
	}
	return lexer.Contents.Text[lexer.Current.Position.Index+1], nil
}

func StepbackLexer(lexer *AwooLexer) (rune, error) {
	return AdvanceLexerFor(lexer, -1)
}

func RunLexer(lexer *AwooLexer) AwooLexerResult {
	result := AwooLexerResult{
		Contents: lexer.Contents,
		Context:  lexer.Context,
	}
	logger.LogExtra(gchalk.Yellow("> Lexer\n"))
	var err error
	for ; err == nil; _, err = AdvanceLexer(lexer) {
		if unicode.IsSpace(lexer.Current.Character) {
			switch lexer.Current.Character {
			case '\n':
				lexer.Current.Position.Column = 0
				lexer.Current.Position.Line++
			case '\t':
				lexer.Current.Position.Column += (cc.AwooTabIndent - 1)
			}
			continue
		}

		single, ok := lexer.Settings.Tokens.Single[lexer.Current.Character]
		if ok {
			token := lexer_token.NewAwooLexerToken(NewAwooLexerTokenPosition(lexer, 1), single)
			PrintNewToken(&lexer.Settings, string(lexer.Current.Character), &token)
			token.Position.Index = uint32(len(result.Tokens))
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if lexer.Current.Character == '\'' {
			token, matchedString, err := CreateTokenChar(lexer)
			if err != nil {
				result.Error = err
				break
			}
			PrintNewToken(&lexer.Settings, matchedString, &token)
			token.Position.Index = uint32(len(result.Tokens))
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if unicode.IsLetter(lexer.Current.Character) {
			token, matchedString := CreateTokenLetter(lexer)
			PrintNewToken(&lexer.Settings, matchedString, &token)
			token.Position.Index = uint32(len(result.Tokens))
			result.Tokens = append(result.Tokens, token)
			continue
		}
		if unicode.IsNumber(lexer.Current.Character) {
			token, matchedString, err := CreateTokenNumber(lexer)
			if err != nil {
				result.Error = err
				break
			}
			PrintNewToken(&lexer.Settings, matchedString, &token)
			token.Position.Index = uint32(len(result.Tokens))
			result.Tokens = append(result.Tokens, token)
			continue
		}

		token, matchedString, ok := CreateTokenCouple(lexer)
		if !ok {
			result.Error = fmt.Errorf("%w: %s", awerrors.ErrorIllegalCharacter, gchalk.Red((string)(lexer.Current.Character)))
			break
		}
		PrintNewToken(&lexer.Settings, matchedString, &token)
		token.Position.Index = uint32(len(result.Tokens))
		result.Tokens = append(result.Tokens, token)
		break
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
