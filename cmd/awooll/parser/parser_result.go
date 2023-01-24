package parser

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/statement"
)

type AwooParserResult struct {
	Error      error
	Statements []statement.AwooParserStatement
}
