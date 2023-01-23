package token

import "golang.org/x/exp/maps"

type AwooToken struct {
	Type   uint16
	Length uint8
}

type AwooTokenMap struct {
	Single map[rune]AwooToken
	Words  map[string]AwooToken
}

func SetupTokenMap() AwooTokenMap {
	tokenMap := AwooTokenMap{
		Single: make(map[rune]AwooToken),
		Words:  make(map[string]AwooToken),
	}
	maps.Copy(tokenMap.Words, TokenMapTypes)

	return tokenMap
}
