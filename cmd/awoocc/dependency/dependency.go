package dependency

import (
	"errors"

	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/lexer_token"
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/token"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

func ResolveDependencies(result *lexer.AwooLexerResult) ([]string, error) {
	dependencies := []string{}
	for i, t := range result.Tokens {
		if t.Type == token.TokenKeywordImport {
			next := result.Tokens[i+1]
			if next.Type != token.TokenTypePrimitive {
				return dependencies, errors.New("nope")
			}
			if lexer_token.GetTokenPrimitiveType(&next) != types.AwooTypeString {
				return dependencies, errors.New("nope")
			}
			dependencies = append(dependencies, lexer_token.GetTokenPrimitiveValue(&next).(string))
		}
	}

	return dependencies, nil
}
