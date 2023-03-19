package parser_details

import "github.com/LamkasDev/awoo-emu/cmd/common/types"

type ConstructStatementDetails struct {
	CanReturn bool
	EndToken  uint16
}

type ConstructExpressionDetails struct {
	Type            types.AwooTypeId
	PendingBrackets uint8
	EndTokens       []uint16
}
