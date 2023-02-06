package parser

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
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

func PeekParser(parser *AwooParser) (lexer_token.AwooLexerToken, bool) {
	if parser.Position+1 >= parser.Length {
		return lexer_token.AwooLexerToken{}, false
	}
	return parser.Contents.Tokens[parser.Position+1], true
}

func StepbackParser(parser *AwooParser) bool {
	return AdvanceParserFor(parser, -1)
}

func FetchTokenParser(cparser *AwooParser) (lexer_token.AwooLexerToken, error) {
	ok := AdvanceParser(cparser)
	if !ok {
		return lexer_token.AwooLexerToken{}, awerrors.ErrorNoMoreTokens
	}
	logger.Log("┣━ %s\n", lexer_token.PrintToken(&cparser.Contents.Context, &cparser.Current))
	return cparser.Current, nil
}

func ExpectTokenParser(cparser *AwooParser, tokenTypes []uint16, tokenName string) (lexer_token.AwooLexerToken, error) {
	t, err := FetchTokenParser(cparser)
	if err != nil {
		return t, err
	}
	if !util.Contains(tokenTypes, t.Type) {
		return t, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(tokenName))
	}

	return t, nil
}
