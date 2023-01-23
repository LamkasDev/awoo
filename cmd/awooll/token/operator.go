package token

const TokenOperatorAddition = 0x1000
const TokenOperatorSubstraction = 0x1001
const TokenOperatorMultiplication = 0x1002
const TokenOperatorDivision = 0x1003

var TokenMapOperators = map[rune]AwooToken{
	'+': {Type: TokenOperatorAddition, Length: 1},
	'-': {Type: TokenOperatorSubstraction, Length: 1},
	'*': {Type: TokenOperatorMultiplication, Length: 1},
	'/': {Type: TokenOperatorDivision, Length: 1},
}
