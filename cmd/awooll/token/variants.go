package token

// General.
const TokenTypeIdentifier = 0x000
const TokenTypePrimitive = 0x001
const TokenTypeType = 0x002
const TokenTypeEndStatement = 0x003
const TokenTypeBracketLeft = 0x004
const TokenTypeBracketRight = 0x005
const TokenTypeBracketCurlyLeft = 0x006
const TokenTypeBracketCurlyRight = 0x007
const TokenTypeNot = 0x008
const TokenTypeReference = 0x009
const TokenTypeComma = 0x010

func IsTokenTypeGeneral(t uint16) bool {
	return t < 0x100
}

// Operators.
const TokenOperatorAddition = 0x100
const TokenOperatorSubstraction = 0x101
const TokenOperatorMultiplication = 0x102
const TokenOperatorDereference = 0x102
const TokenOperatorDivision = 0x103
const TokenOperatorEq = 0x104
const TokenOperatorEqEq = 0x105
const TokenOperatorNotEq = 0x106
const TokenOperatorLT = 0x107
const TokenOperatorLTEQ = 0x108
const TokenOperatorGT = 0x109
const TokenOperatorGTEQ = 0x110

func IsTokenTypeOperator(t uint16) bool {
	return t >= 0x100 && t < 0x200
}
func IsTokenTypeAddSub(t uint16) bool {
	return t >= TokenOperatorAddition && t <= TokenOperatorSubstraction
}
func IsTokenTypeMulDiv(t uint16) bool {
	return t >= TokenOperatorMultiplication && t <= TokenOperatorDivision
}
func IsTokenTypeUnary(t uint16) bool {
	return t >= TokenOperatorAddition && t <= TokenOperatorDivision
}
func IsTokenTypeEquality(t uint16) bool {
	return t >= TokenOperatorEqEq && t <= TokenOperatorGTEQ
}
func DoesTokenTakePrecendence(op uint16, left uint16) bool {
	switch op {
	case TokenOperatorAddition,
		TokenOperatorSubstraction:
		return IsTokenTypeEquality(left)
	case TokenOperatorMultiplication,
		TokenOperatorDivision:
		return IsTokenTypeEquality(left) || IsTokenTypeAddSub(left)
	}

	return false
}

// Keywords.
const TokenTypeVar = 0x200
const TokenTypeTypeDefinition = 0x201
const TokenTypeIf = 0x202
const TokenTypeElse = 0x203
const TokenTypeFunc = 0x204
const TokenTypeReturn = 0x205

func IsTokenTypeKeyword(t uint16) bool {
	return t >= 0x200 && t < 0x300
}

// Print stuffs.
func GetTokenTypeName(t uint16) string {
	if IsTokenTypeGeneral(t) {
		return "token"
	}
	if IsTokenTypeOperator(t) {
		return "op"
	}
	if IsTokenTypeKeyword(t) {
		return "keyword"
	}

	return "??"
}
