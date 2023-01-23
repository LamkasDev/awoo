package lexer

import (
	"fmt"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/jwalton/gchalk"
)

type AwooLexer struct {
	Contents []rune
	Length   uint16
	Position uint16
	Current  rune
	Settings AwooLexerSettings
}

type AwooLexerSettings struct {
	Tokens token.AwooTokenMap
}

type AwooLexerResult struct {
	Error  error
	Tokens []AwooLexerToken
}

func SetupLexer() AwooLexer {
	lexer := AwooLexer{
		Settings: AwooLexerSettings{
			Tokens: token.SetupTokenMap(),
		},
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

func StepbackLexer(lexer *AwooLexer) bool {
	return AdvanceLexerFor(lexer, -1)
}

func RunLexer(lexer *AwooLexer) AwooLexerResult {
	result := AwooLexerResult{}
	fmt.Printf("Handling file of size %d\n", lexer.Length)
	for ok := true; ok; ok = AdvanceLexer(lexer) {
		if unicode.IsSpace(lexer.Current) {
			continue
		}

		single, ok := lexer.Settings.Tokens.Single[lexer.Current]
		if ok {
			fmt.Printf("%c (%s)\n", lexer.Current, gchalk.Green((string)(single.Type)))
			continue
		}
		if unicode.IsLetter(lexer.Current) {
			cs := (string)(lexer.Current)
			for AdvanceLexer(lexer) {
				if !unicode.IsLetter(lexer.Current) && !unicode.IsNumber(lexer.Current) {
					break
				}
				cs += (string)(lexer.Current)
			}
			StepbackLexer(lexer)

			word, ok := lexer.Settings.Tokens.Words[cs]
			if !ok {
				result.Error = fmt.Errorf("unknown word %s", gchalk.Red((string)(cs)))
				break
			}
			fmt.Printf("word %s (%s)\n", cs, gchalk.Green((string)(word.Type)))
		} else if unicode.IsNumber(lexer.Current) {
			cs := (string)(lexer.Current)
			for AdvanceLexer(lexer) {
				if !unicode.IsNumber(lexer.Current) {
					break
				}
				cs += (string)(lexer.Current)
			}
			StepbackLexer(lexer)

			fmt.Printf("number %s\n", cs)
		} else {
			result.Error = fmt.Errorf("illegal character (%s/%d)", gchalk.Red((string)(lexer.Current)), lexer.Current)
			break
		}
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
