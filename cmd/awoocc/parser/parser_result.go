package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/statement"
)

type AwooParserResult struct {
	Error      error
	Statements []statement.AwooParserStatement
}
