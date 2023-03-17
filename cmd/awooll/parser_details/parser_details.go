package parser_details

import "github.com/LamkasDev/awoo-emu/cmd/awooll/types"

type ConstructStatementDetails struct {
	CanReturn bool
	EndToken  uint16
}

type ConstructExpressionDetails struct {
	Type            types.AwooType
	PendingBrackets uint8
	EndTokens       []uint16
}
