package token

import "unicode"

type AwooTokenMap struct {
	All      map[uint16]AwooToken
	Single   map[rune]*AwooToken
	Couple   map[string]*AwooToken
	Keywords map[string]*AwooToken
}

func AddToken(m AwooTokenMap, key string, name string, id uint16) {
	awooToken := AwooToken{
		Key:    key,
		Name:   name,
		Id:     id,
		Length: uint8(len(key)),
	}
	m.All[id] = awooToken
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
	AddToken(m, "", "identifier", TokenTypeIdentifier)
	AddToken(m, "", "primitive", TokenTypePrimitive)
	AddToken(m, "", "type", TokenTypeType)
	AddToken(m, ";", ";", TokenTypeEndStatement)
	AddToken(m, "(", "(", TokenTypeBracketLeft)
	AddToken(m, ")", ")", TokenTypeBracketRight)
	AddToken(m, "{", "{", TokenTypeBracketCurlyLeft)
	AddToken(m, "}", "}", TokenTypeBracketCurlyRight)
	AddToken(m, "!", "!", TokenTypeNot)
	AddToken(m, ",", ",", TokenTypeComma)
	AddToken(m, "[", "[", TokenTypeBracketSquareLeft)
	AddToken(m, "]", "]", TokenTypeBracketSquareRight)

	// Operators
	AddToken(m, "+", "+", TokenOperatorAddition)
	AddToken(m, "-", "-", TokenOperatorSubstraction)
	AddToken(m, "*", "*", TokenOperatorMultiplication)
	AddToken(m, "/", "/", TokenOperatorDivision)
	AddToken(m, "=", "=", TokenOperatorEq)
	AddToken(m, "", "==", TokenOperatorEqEq)
	AddToken(m, "", "!=", TokenOperatorNotEq)
	AddToken(m, "<", "<", TokenOperatorLT)
	AddToken(m, "", "<=", TokenOperatorLTEQ)
	AddToken(m, ">", ">", TokenOperatorGT)
	AddToken(m, "", ">=", TokenOperatorGTEQ)
	AddToken(m, "", "<<", TokenOperatorLS)
	AddToken(m, "", ">>", TokenOperatorRS)
	AddToken(m, "&", "&", TokenOperatorAnd)
	AddToken(m, "|", "|", TokenOperatorOr)

	// Keywords
	AddToken(m, "var", "var", TokenKeywordVar)
	AddToken(m, "type", "type", TokenKeywordTypeDefinition)
	AddToken(m, "if", "if", TokenKeywordIf)
	AddToken(m, "else", "else", TokenKeywordElse)
	AddToken(m, "func", "func", TokenKeywordFunc)
	AddToken(m, "return", "return", TokenKeywordReturn)
	AddToken(m, "for", "for", TokenKeywordFor)
	AddToken(m, "import", "import", TokenKeywordImport)

	return m
}
