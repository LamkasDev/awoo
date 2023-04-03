package parser_error

type AwooParserErrorType uint16

type AwooParserError struct {
	Type       AwooParserErrorType
	Message    string
	Highlights []AwooParserErrorHighlight
}

type AwooParserErrorHighlight struct {
	Start   uint32
	Length  uint32
	Details string
}

const AwooParserErrorNoMoreTokens = AwooParserErrorType(0x000)
const AwooParserErrorExpectedToken = AwooParserErrorType(0x001)
const AwooParserErrorUnknownVariable = AwooParserErrorType(0x002)
const AwooParserErrorUnexpectedToken = AwooParserErrorType(0x003)
const AwooParserErrorUnknownFunction = AwooParserErrorType(0x004)
const AwooParserErrorPrimitiveOverflow = AwooParserErrorType(0x005)
const AwooParserErrorPrimitiveUnderflow = AwooParserErrorType(0x006)

var AwooParserErrorMessages map[AwooParserErrorType]string = map[AwooParserErrorType]string{
	AwooParserErrorNoMoreTokens:       "no more tokens",
	AwooParserErrorExpectedToken:      "expected one of",
	AwooParserErrorUnknownVariable:    "unknown variable",
	AwooParserErrorUnexpectedToken:    "unexpected",
	AwooParserErrorUnknownFunction:    "unknown function",
	AwooParserErrorPrimitiveOverflow:  "primitive overflow",
	AwooParserErrorPrimitiveUnderflow: "primitive underflow",
}

var AwooParserErrorDetails map[AwooParserErrorType]string = map[AwooParserErrorType]string{
	AwooParserErrorNoMoreTokens:       "weh",
	AwooParserErrorExpectedToken:      "weh",
	AwooParserErrorUnknownVariable:    "weh",
	AwooParserErrorUnexpectedToken:    "weh",
	AwooParserErrorUnknownFunction:    "weh",
	AwooParserErrorPrimitiveOverflow:  "weh",
	AwooParserErrorPrimitiveUnderflow: "weh",
}

func CreateParserErrorText(errorType AwooParserErrorType, text string, start uint32, length uint32, details string) *AwooParserError {
	return &AwooParserError{
		Type:    errorType,
		Message: text,
		Highlights: []AwooParserErrorHighlight{
			{
				Start:   start,
				Length:  length,
				Details: details,
			},
		},
	}
}

func CreateParserError(errorType AwooParserErrorType, start uint32, length uint32) *AwooParserError {
	return CreateParserErrorText(errorType, AwooParserErrorMessages[errorType], start, length, AwooParserErrorDetails[errorType])
}
