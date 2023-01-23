package lexer

import (
	"fmt"
	"unicode"

	"github.com/LamkasDev/awoo-emu/cmd/awoo/token"
	"github.com/jwalton/gchalk"
)

type AwooLexer struct {
	Contents []rune
	Length   uint16
	Position uint16
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
}

func RunLexer(lexer *AwooLexer) AwooLexerResult {
	result := AwooLexerResult{}
	fmt.Printf("Handling file of size %d\n", lexer.Length)
	for lexer.Position < lexer.Length {
		c := lexer.Contents[lexer.Position]
		lexer.Position++

		if unicode.IsSpace(c) {
			continue
		}

		single, ok := lexer.Settings.Tokens.Single[c]
		if ok {
			fmt.Printf("%c (%s)\n", c, gchalk.Green((string)(single.Type)))
			continue
		}
		if lexer.Position >= lexer.Length {
			break
		}
		if unicode.IsLetter(c) {
			cs := ""
			for unicode.IsLetter(c) || unicode.IsNumber(c) {
				cs += (string)(c)
				c = lexer.Contents[lexer.Position]
				lexer.Position++
				if lexer.Position >= lexer.Length {
					break
				}
			}
			lexer.Position--

			word, ok := lexer.Settings.Tokens.Words[cs]
			if !ok {
				result.Error = fmt.Errorf("unknown word %s", gchalk.Red((string)(cs)))
				break
			}
			fmt.Printf("word %s (%s)\n", cs, gchalk.Green((string)(word.Type)))
		} else if unicode.IsNumber(c) {
			cs := ""
			for unicode.IsNumber(c) {
				cs += (string)(c)
				c = lexer.Contents[lexer.Position]
				lexer.Position++
				if lexer.Position >= lexer.Length {
					break
				}
			}
			lexer.Position--

			fmt.Printf("number %s\n", cs)
		} else {
			result.Error = fmt.Errorf("illegal character (%s/%d)", gchalk.Red((string)(c)), c)
			break
		}
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
