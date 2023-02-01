package lexer_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awooll/token"
	"github.com/LamkasDev/awoo-emu/cmd/awooll/types"
)

type AwooLexerContext struct {
	Tokens token.AwooTokenMap
	Types  types.AwooTypeMap
}

func GetContextType(context *AwooLexerContext, key string) (*types.AwooType, bool) {
	t, ok := context.Types.Lookup[key]
	return t, ok
}

func GetContextTypeId(context *AwooLexerContext, id uint16) (types.AwooType, bool) {
	t, ok := context.Types.All[id]
	return t, ok
}

func AddContextType(context *AwooLexerContext, t types.AwooType) uint16 {
	return types.AddTypeUserDefined(&context.Types, t)
}

func AddContextTypeAlias(context *AwooLexerContext, original types.AwooType, t types.AwooType) uint16 {
	return AddContextType(context, types.AwooType{
		Key: t.Key, Type: original.Type,
		Length: original.Length, Size: original.Size,
		Flags: original.Flags,
	})
}
