package parser

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

type AwooParser struct {
	Contents lexer.AwooLexerResult
	Length   uint32
	Position uint32
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
		},
		Settings: settings,
	}
	return parser
}

func LoadParser(parser *AwooParser, contents lexer.AwooLexerResult) {
	parser.Contents = contents
	parser.Length = (uint32)(len(contents.Tokens))
	parser.Position = 0
	parser.Current = TransformToken(&parser.Context, parser.Contents.Tokens[parser.Position])
}

func AdvanceParserFor(parser *AwooParser, n int32) *parser_error.AwooParserError {
	pos := (uint32)((int32)(parser.Position) + n)
	if pos >= parser.Length {
		// TODO: finish this
		return parser_error.CreateParserError(parser_error.AwooParserErrorNoMoreTokens, lexer_token.AwooLexerTokenPosition{})
	}
	parser.Position = pos
	parser.Current = TransformToken(&parser.Context, parser.Contents.Tokens[parser.Position])
	return nil
}

func AdvanceParser(parser *AwooParser) *parser_error.AwooParserError {
	return AdvanceParserFor(parser, 1)
}

func StepbackParser(parser *AwooParser) *parser_error.AwooParserError {
	return AdvanceParserFor(parser, -1)
}

func PeekTokenFor(parser *AwooParser, n int32) (lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	pos := (uint32)((int32)(parser.Position) + n)
	if pos >= parser.Length {
		// TODO: finish this
		return lexer_token.AwooLexerToken{}, parser_error.CreateParserError(parser_error.AwooParserErrorNoMoreTokens, lexer_token.AwooLexerTokenPosition{})
	}
	return parser.Contents.Tokens[pos], nil
}

func PeekToken(parser *AwooParser) (lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	return PeekTokenFor(parser, 1)
}

func FetchToken(cparser *AwooParser) (lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	if err := AdvanceParser(cparser); err != nil {
		return lexer_token.AwooLexerToken{}, err
	}
	logger.LogExtra("┣━ %s\n", lexer.PrintToken(&cparser.Settings.Lexer, &cparser.Current))
	return cparser.Current, nil
}

func ExpectToken(cparser *AwooParser, tokenType uint16) (lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	t, err := FetchToken(cparser)
	if err != nil {
		return t, err
	}
	if t.Type != tokenType {
		return t, parser_error.CreateParserErrorText(parser_error.AwooParserErrorExpectedToken,
			fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorExpectedToken], gchalk.Red(cparser.Settings.Lexer.Tokens.All[tokenType].Key)),
			cparser.Current.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorExpectedToken])
	}

	return t, nil
}

func ExpectTokenOptional(cparser *AwooParser, tokenType uint16) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
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

func ExpectTokens(cparser *AwooParser, tokenTypes []uint16) (lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	t, err := FetchToken(cparser)
	if err != nil {
		return t, err
	}
	if !util.Contains(tokenTypes, t.Type) {
		return t, parser_error.CreateParserErrorText(parser_error.AwooParserErrorExpectedToken,
			fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorExpectedToken], gchalk.Red(lexer.PrintTokenTypes(&cparser.Settings.Lexer, tokenTypes))),
			cparser.Current.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorExpectedToken])
	}

	return t, nil
}

func ExpectTokensOptional(cparser *AwooParser, tokenTypes []uint16) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
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
