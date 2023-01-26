package lexer

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/print"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
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

func AdvanceLexerFor(lexer *AwooLexer, n int16) bool {
	lexer.Position = (uint16)((int16)(lexer.Position) + n)
	if lexer.Position >= lexer.Length {
		return false
	}
	lexer.Current = lexer.Contents[lexer.Position]
	return true
}

func AdvanceLexer(lexer *AwooLexer) bool {
	return AdvanceLexerFor(lexer, 1)
}

func PeekLexer(lexer *AwooLexer) rune {
	if lexer.Position+1 >= lexer.Length {
		return 0
	}
	return lexer.Contents[lexer.Position+1]
}

func StepbackLexer(lexer *AwooLexer) bool {
	return AdvanceLexerFor(lexer, -1)
}

type ConstructChunkValidator func(rune) bool

func ConstructChunk(lexer *AwooLexer, validate ConstructChunkValidator) string {
	cs := (string)(lexer.Current)
	for AdvanceLexer(lexer) {
		if !validate(lexer.Current) {
			break
		}
		cs += (string)(lexer.Current)
	}
	StepbackLexer(lexer)

	return cs
}

func RunLexer(lexer *AwooLexer) AwooLexerResult {
	result := AwooLexerResult{
		Context: lexer.Context,
	}
	fmt.Println(gchalk.Yellow("> Lexer"))
	fmt.Printf("Input: %s\n", gchalk.Magenta(string(lexer.Contents)))
	for ok := true; ok; ok = AdvanceLexer(lexer) {
		if unicode.IsSpace(lexer.Current) {
			continue
		}

		single, ok := lexer.Context.Tokens.Single[lexer.Current]
		if ok {
			t := lexer_token.CreateToken(lexer.Position, single)
			print.PrintNewToken(&lexer.Context, string(lexer.Current), &t)
			result.Tokens = append(result.Tokens, t)
			continue
		}
		if unicode.IsLetter(lexer.Current) {
			cs := ConstructChunk(lexer, func(c rune) bool {
				return unicode.IsLetter(c) || unicode.IsNumber(c)
			})
			keyword, ok := lexer.Context.Tokens.Keywords[cs]
			if ok {
				t, ok := lexer.Context.Types.Lookup[cs]
				if ok {
					t := lexer_token.CreateTokenType(lexer.Position, t.Type)
					print.PrintNewToken(&lexer.Context, cs, &t)
					result.Tokens = append(result.Tokens, t)
				} else {
					t := lexer_token.CreateToken(lexer.Position, keyword)
					print.PrintNewToken(&lexer.Context, cs, &t)
					result.Tokens = append(result.Tokens, t)
				}

				continue
			}

			t := lexer_token.CreateTokenIdentifier(lexer.Position, cs)
			print.PrintNewToken(&lexer.Context, cs, &t)
			result.Tokens = append(result.Tokens, t)
			continue
		}
		if unicode.IsNumber(lexer.Current) {
			cs := ConstructChunk(lexer, func(c rune) bool {
				return unicode.IsNumber(c)
			})

			n, err := strconv.ParseInt(cs, 10, 64)
			if err != nil {
				result.Error = err
				break
			}
			t := lexer_token.CreateTokenPrimitive(lexer.Position, types.AwooTypeInt32, n)
			print.PrintNewToken(&lexer.Context, cs, &t)
			result.Tokens = append(result.Tokens, t)
			continue
		}

		cs := ConstructChunk(lexer, func(c rune) bool {
			return unicode.IsPunct(c) || unicode.IsSymbol(c)
		})
		couple, ok := lexer.Context.Tokens.Couple[cs]
		if ok {
			t := lexer_token.CreateToken(lexer.Position, couple)
			print.PrintNewToken(&lexer.Context, cs, &t)
			result.Tokens = append(result.Tokens, t)
			continue
		}

		result.Error = fmt.Errorf("illegal character %s", gchalk.Red((string)(lexer.Current)))
		break
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
