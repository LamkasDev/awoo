package parser

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/awerrors"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
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

type AwooParserSettings struct {
	Lexer    lexer.AwooLexerSettings
	Mappings AwooParserMappings
}

func SetupParser(settings AwooParserSettings, context lexer_context.AwooLexerContext) AwooParser {
	parser := AwooParser{
		Context: parser_context.AwooParserContext{
			Lexer:  context,
			Scopes: parser_context.SetupParserScopeContainer(),
			Functions: parser_context.AwooParserFunctionContainer{
				Entries: map[string]parser_context.AwooParserFunction{},
			},
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

func AdvanceParserFor(parser *AwooParser, n int16) error {
	parser.Position = (uint16)((int16)(parser.Position) + n)
	if parser.Position >= parser.Length {
		return awerrors.ErrorNoMoreTokens
	}
	parser.Current = TransformToken(&parser.Context, parser.Contents.Tokens[parser.Position])
	return nil
}

func AdvanceParser(parser *AwooParser) error {
	return AdvanceParserFor(parser, 1)
}

func StepbackParser(parser *AwooParser) error {
	return AdvanceParserFor(parser, -1)
}

func PeekToken(parser *AwooParser) (lexer_token.AwooLexerToken, error) {
	if parser.Position+1 >= parser.Length {
		return lexer_token.AwooLexerToken{}, awerrors.ErrorNoMoreTokens
	}
	return parser.Contents.Tokens[parser.Position+1], nil
}

func FetchToken(cparser *AwooParser) (lexer_token.AwooLexerToken, error) {
	if err := AdvanceParser(cparser); err != nil {
		return lexer_token.AwooLexerToken{}, err
	}
	logger.Log("┣━ %s\n", lexer.PrintToken(&cparser.Settings.Lexer, &cparser.Current))
	return cparser.Current, nil
}

func ExpectToken(cparser *AwooParser, tokenType uint16) (lexer_token.AwooLexerToken, error) {
	t, err := FetchToken(cparser)
	if err != nil {
		return t, err
	}
	if t.Type != tokenType {
		return t, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(cparser.Settings.Lexer.Tokens.All[tokenType].Key))
	}

	return t, nil
}

func ExpectTokenOptional(cparser *AwooParser, tokenType uint16) (*lexer_token.AwooLexerToken, error) {
	t, err := PeekToken(cparser)
	if err != nil {
		return nil, err
	}
	if t.Type != tokenType {
		return nil, nil
	}
	AdvanceParser(cparser)

	return &t, nil
}

func ExpectTokens(cparser *AwooParser, tokenTypes []uint16) (lexer_token.AwooLexerToken, error) {
	t, err := FetchToken(cparser)
	if err != nil {
		return t, err
	}
	if !util.Contains(tokenTypes, t.Type) {
		return t, fmt.Errorf("%w: %s", awerrors.ErrorExpectedToken, gchalk.Red(lexer.PrintTokenTypes(&cparser.Settings.Lexer, tokenTypes)))
	}

	return t, nil
}

func ExpectTokensOptional(cparser *AwooParser, tokenTypes []uint16) (*lexer_token.AwooLexerToken, error) {
	t, err := PeekToken(cparser)
	if err != nil {
		return nil, err
	}
	if !util.Contains(tokenTypes, t.Type) {
		return nil, nil
	}
	AdvanceParser(cparser)

	return &t, nil
}
