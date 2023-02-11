package parser_details

import "github.com/LamkasDev/awoo-emu/cmd/awooll/types"

type ConstructStatementDetails struct {
	CanReturn bool
}

type ConstructExpressionDetails struct {
	Type            types.AwooType
	PendingBrackets uint8
	EndToken        uint16
}
