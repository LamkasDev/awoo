package token

// General
const TokenTypeIdentifier = 0x000
const TokenTypePrimitive = 0x001
const TokenTypeType = 0x002
const TokenTypeEndStatement = 0x003
const TokenTypeBracketLeft = 0x004
const TokenTypeBracketRight = 0x005
const TokenTypeBracketCurlyLeft = 0x006
const TokenTypeBracketCurlyRight = 0x007
const TokenTypeNot = 0x008

func IsTokenTypeGeneral(t uint16) bool {
	return t < 0x100
}

// Operators
const TokenOperatorAddition = 0x100
const TokenOperatorSubstraction = 0x101
const TokenOperatorMultiplication = 0x102
const TokenOperatorDivision = 0x103
const TokenOperatorEq = 0x104
const TokenOperatorEqEq = 0x105
const TokenOperatorNotEq = 0x106

func IsTokenTypeOperator(t uint16) bool {
	return t >= 0x100 && t < 0x200
}

// Keywords
const TokenTypeVar = 0x200
const TokenTypeTypeDefinition = 0x201
const TokenTypeIf = 0x202

func IsTokenTypeKeyword(t uint16) bool {
	return t >= 0x200 && t < 0x300
}
