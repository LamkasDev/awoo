package token

import "unicode"

type AwooToken struct {
	Key    string
	Name   string
	Type   uint16
	Length uint8
}

type AwooTokenMap struct {
	All      map[uint16]AwooToken
	Single   map[rune]*AwooToken
	Couple   map[string]*AwooToken
	Keywords map[string]*AwooToken
}

func AddToken(m AwooTokenMap, key string, name string, tokenType uint16) {
	awooToken := AwooToken{
		Key:    key,
		Name:   name,
		Type:   tokenType,
		Length: uint8(len(key)),
	}
	m.All[tokenType] = awooToken
	if awooToken.Length == 0 {
		return
	}

	c := rune(key[0])
	if unicode.IsPunct(c) || unicode.IsSymbol(c) {
		if awooToken.Length == 1 {
			m.Single[rune(key[0])] = &awooToken
			return
		}
		m.Couple[key] = &awooToken
		return
	}

	m.Keywords[key] = &awooToken
}

func SetupTokenMap() AwooTokenMap {
	m := AwooTokenMap{
		All:      make(map[uint16]AwooToken),
		Single:   make(map[rune]*AwooToken),
		Couple:   make(map[string]*AwooToken),
		Keywords: make(map[string]*AwooToken),
	}

	// General
	AddToken(m, "", "id", TokenTypeIdentifier)
	AddToken(m, "", "prim", TokenTypePrimitive)
	AddToken(m, "", "type", TokenTypeType)
	AddToken(m, ";", ";", TokenTypeEndStatement)
	AddToken(m, "(", "(", TokenTypeBracketLeft)
	AddToken(m, ")", ")", TokenTypeBracketRight)

	// Operators
	AddToken(m, "+", "+", TokenOperatorAddition)
	AddToken(m, "-", "-", TokenOperatorSubstraction)
	AddToken(m, "*", "*", TokenOperatorMultiplication)
	AddToken(m, "/", "/", TokenOperatorDivision)
	AddToken(m, "=", "=", TokenOperatorEq)
	AddToken(m, "==", "==", TokenOperatorEqEq)

	// Keywords
	AddToken(m, "var", "var", TokenTypeVar)

	return m
}
