package types

type AwooTypeMap struct {
	All           map[uint16]AwooType
	Lookup        map[string]*AwooType
	UserDefinedId uint16
}

func AddTypeAt(m *AwooTypeMap, t AwooType) uint16 {
	t.Length = uint8(len(t.Key))
	m.All[t.Id] = t
	m.Lookup[t.Key] = &t

	return t.Id
}

func AddTypeBuiltin(m *AwooTypeMap, t AwooType) uint16 {
	return AddTypeAt(m, t)
}

func AddTypeUserDefined(m *AwooTypeMap, t AwooType) uint16 {
	t.Id = m.UserDefinedId
	awooType := AddTypeAt(m, t)
	m.UserDefinedId++

	return awooType
}

func SetupTypeMap() AwooTypeMap {
	m := AwooTypeMap{
		All:           make(map[uint16]AwooType),
		Lookup:        make(map[string]*AwooType),
		UserDefinedId: AwooTypeUserDefinedStart,
	}

	AddTypeBuiltin(&m, AwooType{
		Key: "bool", Id: AwooTypeBoolean, Type: AwooTypeBoolean,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "byte", Id: AwooTypeByte, Type: AwooTypeByte,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "char", Id: AwooTypeChar, Type: AwooTypeChar,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "ptr", Id: AwooTypePointer, Type: AwooTypePointer,
		Size: 4,
	})

	AddTypeBuiltin(&m, AwooType{
		Key: "int8", Id: AwooTypeInt8, Type: AwooTypeInt8,
		Size: 1, Flags: AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int16", Id: AwooTypeInt16, Type: AwooTypeInt16,
		Size: 2, Flags: AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int32", Id: AwooTypeInt32, Type: AwooTypeInt32,
		Size: 4, Flags: AwooTypeFlagsSign,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "int64", Id: AwooTypeInt64, Type: AwooTypeInt64,
		Size: 8, Flags: AwooTypeFlagsSign,
	})

	AddTypeBuiltin(&m, AwooType{
		Key: "uint8", Id: AwooTypeUInt8, Type: AwooTypeUInt8,
		Size: 1,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint16", Id: AwooTypeUInt16, Type: AwooTypeUInt16,
		Size: 2,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint32", Id: AwooTypeUInt32, Type: AwooTypeUInt32,
		Size: 4,
	})
	AddTypeBuiltin(&m, AwooType{
		Key: "uint64", Id: AwooTypeUInt64, Type: AwooTypeUInt64,
		Size: 8,
	})

	/* AddType(&m, "float", 4)
	AddType(&m, "float32", 4)
	AddType(&m, "float64", 8) */

	return m
}
