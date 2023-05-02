package types

import "github.com/LamkasDev/awoo-emu/cmd/common/types"

type AwooTypeMap struct {
	All           map[types.AwooTypeId]AwooType
	Lookup        map[string]*AwooType
	UserDefinedId types.AwooTypeId
}

func AddTypeAt(m *AwooTypeMap, t AwooType) types.AwooTypeId {
	t.Length = uint8(len(t.Key))
	m.All[t.Id] = t
	if t.Length > 0 {
		m.Lookup[t.Key] = &t
	}

	return t.Id
}

func AddTypeBuiltin(m *AwooTypeMap, t AwooType) types.AwooTypeId {
	return AddTypeAt(m, t)
}

func AddTypeUserDefined(m *AwooTypeMap, t AwooType) types.AwooTypeId {
	t.Id = m.UserDefinedId
	awooType := AddTypeAt(m, t)
	m.UserDefinedId++

	return awooType
}

func SetupTypeMap() AwooTypeMap {
	m := AwooTypeMap{
		All:           make(map[types.AwooTypeId]AwooType),
		Lookup:        make(map[string]*AwooType),
		UserDefinedId: types.AwooTypeUserDefinedStart,
	}

	AddTypeBuiltin(&m, AwooType{
		Key: "bool", Id: types.AwooTypeBoolean, PrimitiveType: types.AwooTypeBoolean,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "byte", Id: types.AwooTypeByte, PrimitiveType: types.AwooTypeByte,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "char", Id: types.AwooTypeChar, PrimitiveType: types.AwooTypeChar,
		Size: 4, Flags: types.AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "ptr", Id: types.AwooTypePointer, PrimitiveType: types.AwooTypePointer,
		Size: 4,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "", Id: types.AwooTypeFunction, PrimitiveType: types.AwooTypeFunction,
	})

	AddTypeBuiltin(&m, AwooType{
		Key: "int8", Id: types.AwooTypeInt8, PrimitiveType: types.AwooTypeInt8,
		Size: 1, Flags: types.AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int16", Id: types.AwooTypeInt16, PrimitiveType: types.AwooTypeInt16,
		Size: 2, Flags: types.AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int32", Id: types.AwooTypeInt32, PrimitiveType: types.AwooTypeInt32,
		Size: 4, Flags: types.AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int64", Id: types.AwooTypeInt64, PrimitiveType: types.AwooTypeInt64,
		Size: 8, Flags: types.AwooTypeFlagsSign,
	})

	AddTypeBuiltin(&m, AwooType{
		Key: "uint8", Id: types.AwooTypeUInt8, PrimitiveType: types.AwooTypeUInt8,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint16", Id: types.AwooTypeUInt16, PrimitiveType: types.AwooTypeUInt16,
		Size: 2,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint32", Id: types.AwooTypeUInt32, PrimitiveType: types.AwooTypeUInt32,
		Size: 4,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint64", Id: types.AwooTypeUInt64, PrimitiveType: types.AwooTypeUInt64,
		Size: 8,
	})

	/* AddType(&m, "float", 4)
	AddType(&m, "float32", 4)
	AddType(&m, "float64", 8) */

	return m
}
