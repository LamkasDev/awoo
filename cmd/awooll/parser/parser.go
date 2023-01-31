package parser

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
	"github.com/jwalton/gchalk"
)

type AwooParser struct {
	Contents lexer.AwooLexerResult
	Length   uint16
	Position uint16
	Current  lexer_token.AwooLexerToken
	Context  parser_context.AwooParserContext
	Settings AwooParserSettings
}

type AwooParserSettings struct{}

func SetupParser(settings AwooParserSettings, context lexer_context.AwooLexerContext) AwooParser {
	parser := AwooParser{
		Context: parser_context.AwooParserContext{
			Lexer:     context,
			Variables: make(map[string]parser_context.AwooParserContextVariable),
		},
		Settings: settings,
	}
	return parser
}

func LoadParser(parser *AwooParser, contents lexer.AwooLexerResult) {
	parser.Contents = contents
	parser.Length = (uint16)(len(contents.Tokens))
	parser.Position = 0
	parser.Current = TransformToken(&parser.Context, parser.Contents.Tokens[parser.Position])
}

func AdvanceParserFor(parser *AwooParser, n int16) bool {
	parser.Position = (uint16)((int16)(parser.Position) + n)
	if parser.Position >= parser.Length {
		return false
	}
	parser.Current = TransformToken(&parser.Context, parser.Contents.Tokens[parser.Position])
	return true
}

func AdvanceParser(parser *AwooParser) bool {
	return AdvanceParserFor(parser, 1)
}

func StepbackParser(parser *AwooParser) bool {
	return AdvanceParserFor(parser, -1)
}

func RunParser(parser *AwooParser) AwooParserResult {
	result := AwooParserResult{}
	fmt.Println(gchalk.Yellow("\n> Parser"))
	fmt.Printf("Input: %s\n", gchalk.Magenta(fmt.Sprintf("%v", parser.Contents.Tokens)))
	for ok := true; ok; ok = AdvanceParser(parser) {
		fmt.Printf("┏━ %s\n", lexer_token.PrintToken(&parser.Contents.Context, &parser.Current))
		st, err := statement.ConstructStatement(&parser.Context, parser.Current, func() (lexer_token.AwooLexerToken, error) {
			ok := AdvanceParser(parser)
			if !ok {
				return lexer_token.AwooLexerToken{}, fmt.Errorf("no more tokens")
			}
			fmt.Printf("┣━ %s\n", lexer_token.PrintToken(&parser.Contents.Context, &parser.Current))
			return parser.Current, nil
		})
		if err != nil {
			result.Error = err
			break
		}
		statement.PrintNewStatement(&parser.Contents.Context, &st)
		result.Statements = append(result.Statements, st)
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result
}
