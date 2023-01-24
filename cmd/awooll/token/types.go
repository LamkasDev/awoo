package token

// General
const TokenTypeIdentifier = 0x0000
const TokenTypePrimitive = 0x0001
const TokenTypeType = 0x0002

func IsTokenTypeGeneral(t uint16) bool {
	return t < 0x1000
}

// Operators
const TokenOperatorAddition = 0x1000
const TokenOperatorSubstraction = 0x1001
const TokenOperatorMultiplication = 0x1002
const TokenOperatorDivision = 0x1003
const TokenOperatorEq = 0x1004
const TokenOperatorEqEq = 0x1005
const TokenOperatorEndStatement = 0x1006

func IsTokenTypeOperator(t uint16) bool {
	return t >= 0x1000 && t < 0x2000
}

// Keywords
const TokenTypeVar = 0x2000

func IsTokenTypeKeyword(t uint16) bool {
	return t >= 0x2000 && t < 0x3000
}
