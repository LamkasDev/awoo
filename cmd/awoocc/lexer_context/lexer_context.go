package lexer_context

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoocc/types"
	commonTypes "github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type AwooLexerContext struct {
	Types types.AwooTypeMap
}

func GetContextType(context *AwooLexerContext, key string) (*types.AwooType, bool) {
	t, ok := context.Types.Lookup[key]
	return t, ok
}

func GetContextTypeId(context *AwooLexerContext, id commonTypes.AwooTypeId) (types.AwooType, bool) {
	t, ok := context.Types.All[id]
	return t, ok
}

func AddContextType(context *AwooLexerContext, t types.AwooType) commonTypes.AwooTypeId {
	return types.AddTypeUserDefined(&context.Types, t)
}

func AddContextTypeAlias(context *AwooLexerContext, original types.AwooType, t types.AwooType) commonTypes.AwooTypeId {
	return AddContextType(context, types.AwooType{
		Key: t.Key, PrimitiveType: original.PrimitiveType,
		Length: original.Length, Size: original.Size,
		Flags: original.Flags,
	})
}
