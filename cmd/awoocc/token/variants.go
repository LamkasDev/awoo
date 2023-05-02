package token

// General.
const TokenTypeIdentifier = uint16(0x000)
const TokenTypePrimitive = uint16(0x001)
const TokenTypeType = uint16(0x002)
const TokenTypeEndStatement = uint16(0x003)
const TokenTypeBracketLeft = uint16(0x004)
const TokenTypeBracketRight = uint16(0x005)
const TokenTypeBracketCurlyLeft = uint16(0x006)
const TokenTypeBracketCurlyRight = uint16(0x007)
const TokenTypeNot = uint16(0x008)
const TokenTypeComma = uint16(0x009)
const TokenTypeBracketSquareLeft = uint16(0x00A)
const TokenTypeBracketSquareRight = uint16(0x00B)

func IsTokenGeneral(t uint16) bool {
	return t < 0x100
}

// Operators.
const TokenOperatorAddition = uint16(0x100)
const TokenOperatorSubstraction = uint16(0x101)
const TokenOperatorMultiplication = uint16(0x102)
const TokenOperatorDereference = uint16(0x102)
const TokenOperatorDivision = uint16(0x103)
const TokenOperatorEq = uint16(0x104)
const TokenOperatorEqEq = uint16(0x105)
const TokenOperatorNotEq = uint16(0x106)
const TokenOperatorLT = uint16(0x107)
const TokenOperatorLTEQ = uint16(0x108)
const TokenOperatorGT = uint16(0x109)
const TokenOperatorGTEQ = uint16(0x110)
const TokenOperatorLS = uint16(0x111)
const TokenOperatorRS = uint16(0x112)
const TokenOperatorAnd = uint16(0x113)
const TokenOperatorReference = uint16(0x113)
const TokenOperatorOr = uint16(0x114)

func IsTokenOperator(t uint16) bool {
	return t >= 0x100 && t < 0x200
}
func IsTokenOperatorAddSub(t uint16) bool {
	return t >= TokenOperatorAddition && t <= TokenOperatorSubstraction
}
func IsTokenOperatorMulDiv(t uint16) bool {
	return t >= TokenOperatorMultiplication && t <= TokenOperatorDivision
}
func IsTokenOperatorUnary(t uint16) bool {
	return t >= TokenOperatorAddition && t <= TokenOperatorDivision
}
func IsTokenOperatorEquality(t uint16) bool {
	return t >= TokenOperatorEqEq && t <= TokenOperatorGTEQ
}
func DoesTokenTakePrecendence(op uint16, left uint16) bool {
	switch op {
	case TokenOperatorAddition,
		TokenOperatorSubstraction:
		return IsTokenOperatorEquality(left)
	case TokenOperatorMultiplication,
		TokenOperatorDivision:
		return IsTokenOperatorEquality(left) || IsTokenOperatorAddSub(left)
	}

	return false
}

// Keywords.
const TokenKeywordVar = uint16(0x200)
const TokenKeywordTypeDefinition = uint16(0x201)
const TokenKeywordIf = uint16(0x202)
const TokenKeywordElse = uint16(0x203)
const TokenKeywordFunc = uint16(0x204)
const TokenKeywordReturn = uint16(0x205)
const TokenKeywordFor = uint16(0x206)
const TokenKeywordImport = uint16(0x207)

func IsTokenKeyword(t uint16) bool {
	return t >= 0x200 && t < 0x300
}

// Print stuffs.
func GetTokenTypeName(t uint16) string {
	if IsTokenGeneral(t) {
		return "token"
	}
	if IsTokenOperator(t) {
		return "op"
	}
	if IsTokenKeyword(t) {
		return "keyword"
	}

	return "??"
}
