package parser_error

import "github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"

type AwooParserErrorType uint16

type AwooParserError struct {
	Type       AwooParserErrorType
	Message    string
	Highlights []AwooParserErrorHighlight
}

type AwooParserErrorHighlight struct {
	Position lexer_token.AwooLexerTokenPosition
	Details  string
}

const AwooParserErrorNoMoreTokens = AwooParserErrorType(0x000)
const AwooParserErrorExpectedToken = AwooParserErrorType(0x001)
const AwooParserErrorUnexpectedToken = AwooParserErrorType(0x002)
const AwooParserErrorUnknownVariable = AwooParserErrorType(0x003)
const AwooParserErrorUnknownFunction = AwooParserErrorType(0x004)
const AwooParserErrorAlreadyDefinedVariable = AwooParserErrorType(0x005)
const AwooParserErrorAlreadyDefinedFunction = AwooParserErrorType(0x006)
const AwooParserErrorPrimitiveOverflow = AwooParserErrorType(0x007)
const AwooParserErrorPrimitiveUnderflow = AwooParserErrorType(0x008)

var AwooParserErrorMessages = map[AwooParserErrorType]string{
	AwooParserErrorNoMoreTokens:           "no more tokens",
	AwooParserErrorExpectedToken:          "expected one of",
	AwooParserErrorUnexpectedToken:        "unexpected",
	AwooParserErrorUnknownVariable:        "unknown variable",
	AwooParserErrorUnknownFunction:        "unknown function",
	AwooParserErrorAlreadyDefinedVariable: "already defined variable",
	AwooParserErrorAlreadyDefinedFunction: "already defined function",
	AwooParserErrorPrimitiveOverflow:      "primitive overflow",
	AwooParserErrorPrimitiveUnderflow:     "primitive underflow",
}

var AwooParserErrorDetails = map[AwooParserErrorType]string{
	AwooParserErrorNoMoreTokens:           "last here",
	AwooParserErrorExpectedToken:          "found instead",
	AwooParserErrorUnexpectedToken:        "not applicable",
	AwooParserErrorUnknownVariable:        "not found",
	AwooParserErrorUnknownFunction:        "not found",
	AwooParserErrorAlreadyDefinedVariable: "pick different name",
	AwooParserErrorAlreadyDefinedFunction: "pick different name",
	AwooParserErrorPrimitiveOverflow:      "occurs here",
	AwooParserErrorPrimitiveUnderflow:     "occurs here",
}

func CreateParserErrorText(errorType AwooParserErrorType, text string, position lexer_token.AwooLexerTokenPosition, details string) *AwooParserError {
	return &AwooParserError{
		Type:    errorType,
		Message: text,
		Highlights: []AwooParserErrorHighlight{
			{
				Position: position,
				Details:  details,
			},
		},
	}
}

func CreateParserError(errorType AwooParserErrorType, position lexer_token.AwooLexerTokenPosition) *AwooParserError {
	return CreateParserErrorText(errorType, AwooParserErrorMessages[errorType], position, AwooParserErrorDetails[errorType])
}
