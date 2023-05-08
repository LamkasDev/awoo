package parser

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_error"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/scope"
	"github.com/LamkasDev/awoo-emu/cmd/common/util"
	"github.com/jwalton/gchalk"
)

type AwooParser struct {
	Contents AwooParserContents
	Context  AwooParserContext
	Settings AwooParserSettings
}

type AwooParserContents struct {
	Data     lexer.AwooLexerResult
	Length   uint32
	Position uint32
}

type AwooParserContext struct {
	Lexer  lexer_context.AwooLexerContext
	Scopes scope.AwooScopeContainer
}

type AwooParserSettings struct {
	Lexer    lexer.AwooLexerSettings
	Mappings AwooParserMappings
}

func NewParser(settings AwooParserSettings, context lexer_context.AwooLexerContext, data lexer.AwooLexerResult) AwooParser {
	return AwooParser{
		Contents: AwooParserContents{
			Data:   data,
			Length: (uint32)(len(data.Tokens)),
		},
		Context: AwooParserContext{
			Lexer:  context,
			Scopes: scope.NewScopeContainer(),
		},
		Settings: settings,
	}
}

func GetParserToken(cparser *AwooParser) lexer_token.AwooLexerToken {
	return TransformToken(&cparser.Context, cparser.Contents.Data.Tokens[cparser.Contents.Position])
}

func AdvanceParserFor(cparser *AwooParser, n int32) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	pos := (uint32)((int32)(cparser.Contents.Position) + n)
	if pos >= cparser.Contents.Length {
		return nil, parser_error.CreateParserError(parser_error.AwooParserErrorNoMoreTokens, lexer_token.AwooLexerTokenPosition{})
	}
	cparser.Contents.Position = pos
	token := GetParserToken(cparser)
	return &token, nil
}

func AdvanceParser(cparser *AwooParser) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	return AdvanceParserFor(cparser, 1)
}

func StepbackParser(cparser *AwooParser) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	return AdvanceParserFor(cparser, -1)
}

func PeekParserFor(cparser *AwooParser, n int32) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	pos := (uint32)((int32)(cparser.Contents.Position) + n)
	if pos >= cparser.Contents.Length {
		return nil, parser_error.CreateParserError(parser_error.AwooParserErrorNoMoreTokens, lexer_token.AwooLexerTokenPosition{})
	}
	token := cparser.Contents.Data.Tokens[pos]
	return &token, nil
}

func PeekParser(cparser *AwooParser) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	return PeekParserFor(cparser, 1)
}

func ExpectToken(cparser *AwooParser, tokenType uint16) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	token, err := AdvanceParser(cparser)
	if err == nil && token.Type == tokenType {
		return token, nil
	}

	return nil, parser_error.CreateParserErrorText(parser_error.AwooParserErrorExpectedToken,
		fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorExpectedToken], gchalk.Red(cparser.Settings.Lexer.Tokens.All[tokenType].Key)),
		token.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorExpectedToken])
}

func ExpectTokenOptional(cparser *AwooParser, tokenType uint16) *lexer_token.AwooLexerToken {
	if token, _ := PeekParser(cparser); token != nil && token.Type == tokenType {
		AdvanceParser(cparser)
		return token
	}

	return nil
}

func ExpectTokens(cparser *AwooParser, tokenTypes []uint16) (*lexer_token.AwooLexerToken, *parser_error.AwooParserError) {
	token, err := AdvanceParser(cparser)
	if err == nil && util.Contains(tokenTypes, token.Type) {
		return token, nil
	}

	return nil, parser_error.CreateParserErrorText(parser_error.AwooParserErrorExpectedToken,
		fmt.Sprintf("%s: %s", parser_error.AwooParserErrorMessages[parser_error.AwooParserErrorExpectedToken], gchalk.Red(lexer.PrintTokenTypes(&cparser.Settings.Lexer, tokenTypes))),
		token.Position, parser_error.AwooParserErrorDetails[parser_error.AwooParserErrorExpectedToken])
}

func ExpectTokensOptional(cparser *AwooParser, tokenTypes []uint16) *lexer_token.AwooLexerToken {
	if token, _ := PeekParser(cparser); token != nil && util.Contains(tokenTypes, token.Type) {
		AdvanceParser(cparser)
		return token
	}

	return nil
}
