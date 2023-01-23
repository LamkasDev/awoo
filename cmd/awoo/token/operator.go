package token

const TokenOperatorPlus = 0x1000

var TokenMapOperators = map[rune]AwooToken{
	'+': {Type: TokenOperatorPlus, Length: 1},
}
