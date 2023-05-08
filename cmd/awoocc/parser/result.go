package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
)

type AwooParserResult struct {
	Error      error
	Context    AwooParserContext
	Statements []statement.AwooParserStatement
}
