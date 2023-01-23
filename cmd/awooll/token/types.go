package token

const TokenTypeInt = 0x0001

var TokenMapTypes = map[string]AwooToken{
	"int": {Type: TokenTypeInt, Length: (uint8)(len("int"))},
}
