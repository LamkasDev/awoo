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
		UserDefinedId: AwooTypeUserDefinedStart,
	}

	AddTypeBuiltin(&m, AwooType{
		Key: "bool", Id: types.AwooTypeId(AwooTypeBoolean), PrimitiveType: AwooTypeBoolean,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "byte", Id: types.AwooTypeId(AwooTypeByte), PrimitiveType: AwooTypeByte,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "char", Id: types.AwooTypeId(AwooTypeChar), PrimitiveType: AwooTypeChar,
		Size: 4, Flags: AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "ptr", Id: types.AwooTypeId(AwooTypePointer), PrimitiveType: AwooTypePointer,
		Size: 4,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "", Id: types.AwooTypeId(AwooTypeFunction), PrimitiveType: AwooTypeFunction,
	})

	AddTypeBuiltin(&m, AwooType{
		Key: "int8", Id: types.AwooTypeId(AwooTypeInt8), PrimitiveType: AwooTypeInt8,
		Size: 1, Flags: AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int16", Id: types.AwooTypeId(AwooTypeInt16), PrimitiveType: AwooTypeInt16,
		Size: 2, Flags: AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int32", Id: types.AwooTypeId(AwooTypeInt32), PrimitiveType: AwooTypeInt32,
		Size: 4, Flags: AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int64", Id: types.AwooTypeId(AwooTypeInt64), PrimitiveType: AwooTypeInt64,
		Size: 8, Flags: AwooTypeFlagsSign,
	})

	AddTypeBuiltin(&m, AwooType{
		Key: "uint8", Id: types.AwooTypeId(AwooTypeUInt8), PrimitiveType: AwooTypeUInt8,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint16", Id: types.AwooTypeId(AwooTypeUInt16), PrimitiveType: AwooTypeUInt16,
		Size: 2,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint32", Id: types.AwooTypeId(AwooTypeUInt32), PrimitiveType: AwooTypeUInt32,
		Size: 4,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint64", Id: types.AwooTypeId(AwooTypeUInt64), PrimitiveType: AwooTypeUInt64,
		Size: 8,
	})

	/* AddType(&m, "float", 4)
	AddType(&m, "float32", 4)
	AddType(&m, "float64", 8) */

	return m
}
