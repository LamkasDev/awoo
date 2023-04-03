package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/parser_context"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
)

type AwooParserResult struct {
	Error      error
	Context    parser_context.AwooParserContext
	Statements []statement.AwooParserStatement
}
